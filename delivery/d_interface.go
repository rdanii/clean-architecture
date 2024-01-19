package delivery

import (
	"clean-architecture/usecase"

	"github.com/gin-gonic/gin"
)

type Delivery interface {
	Ping(c *gin.Context)
	NoRoute(c *gin.Context)
}

type delivery struct {
	use usecase.Usecase
}

func NewDelivery(u usecase.Usecase) Delivery {
	return &delivery{
		use: u,
	}
}

func (d *delivery) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (dlvr *delivery) NoRoute(c *gin.Context) {
	c.JSON(404, gin.H{
		"message": "route not found",
	})
}
