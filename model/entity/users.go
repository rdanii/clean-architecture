package entity

import (
	"time"
)

type Users struct {
	Id                 string    `json:"id" gorm:"column:id"`
	Username           string    `json:"username" gorm:"column:username"`
	Email              string    `json:"email" gorm:"column:email"`
	Password           string    `json:"password" gorm:"column:password"`
	Role               string    `json:"role" gorm:"column:role"`
	FullName           string    `json:"fullName" gorm:"column:full_name"`
	StatusVerification bool      `json:"statusVerification" gorm:"column:status_verification"`
	CreatedAt          time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt          time.Time `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt          time.Time `json:"deletedAt" gorm:"column:deleted_at"`
}
