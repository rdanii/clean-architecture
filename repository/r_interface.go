package repository

import "gorm.io/gorm"

type Repository interface {
	Ping() string
}

type repository struct {
	postgresConn *gorm.DB
}

func NewRepository(postgresConn *gorm.DB) Repository {
	return &repository{
		postgresConn: postgresConn,
	}
}

func (r *repository) Ping() string {
	return "pong"
}
