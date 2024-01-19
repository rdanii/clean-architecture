package auth

import (
	"clean-architecture/model/dto/response"
	"clean-architecture/model/dto/util"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Auth struct {
	User  string `json:"user" gorm:"column:user"`
	Pass  string `json:"password" gorm:"column:password"`
	Roles string `json:"roles" gorm:"column:roles"`
}

type Tabler interface {
	TableName() string
}

func (Auth) TableName() string {
	return "user_chub"
}

var authList []Auth

func Initialize(list []Auth) error {
	if len(list) == 0 {
		return errors.New("please provide list of authenticated users")
	}
	authList = list
	return nil
}

func InitializeUserFromDBGormV2(conn *gorm.DB) error {

	var serviceName = viper.GetString("SERVICE_NAME")
	var authUsers []Auth
	err := conn.
		Model(&Auth{}).
		Where("services = ANY(ARRAY[?, 'all'])", serviceName).
		Scan(&authUsers).
		Error

	if err != nil {
		return err
	}

	if len(authUsers) == 0 {
		return fmt.Errorf("no authenticated users found for service %s", serviceName)
	}

	authList = authUsers
	return nil
}

func getUser(user string) *Auth {
	for _, v := range authList {
		if v.User == user {
			return &v
		}
	}
	return nil
}

func ResponseUnauthorized(c *gin.Context) {
	resp := response.BuildUnauthorizedResponse(response.RESPONSE_MESSAGE_INVALID_AUTH_TOKEN, "Basic Auth Authorization")
	c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
	c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
}

func BasicAuth_v2() gin.HandlerFunc {

	return func(c *gin.Context) {

		user, password, hasAuth := c.Request.BasicAuth()

		// check if hasAuth
		if !hasAuth {
			ResponseUnauthorized(c)
			return
		}

		// get user
		fetchedUser := getUser(user)
		if fetchedUser == nil {
			ResponseUnauthorized(c)
			return
		}

		// check passS
		if fetchedUser.Pass != util.GetMD5Hash(password) {
			ResponseUnauthorized(c)
			return
		}

		c.Set("user", fetchedUser.User)
		c.Set("roles", fetchedUser.Roles)
	}
}
