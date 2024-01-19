package usecase

import (
	"clean-architecture/repository"
)

type Usecase interface {
	Ping() string
}

type usecase struct {
	Repo repository.Repository
}

func NewUsecase(r repository.Repository) Usecase {
	return &usecase{
		Repo: r,
	}
}

func (u *usecase) Ping() string {
	return u.Repo.Ping()
}
