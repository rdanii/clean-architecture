package entity

const (
	TABLE_USERS            = "users"
	TABLE_USER_TRANSACTION = "user_transaction"
	TABLE_USER_HAS_COURSE  = "user_has_course"
	TABLE_USER_CART        = "user_cart"
	TABLE_ROLES            = "roles"
	TABLE_COURSES          = "courses"
	TABLE_CART_ITEM        = "cart_item"
)

type Tabler interface {
	TableName() string
}

func (Users) TableName() string {
	return TABLE_USERS
}
