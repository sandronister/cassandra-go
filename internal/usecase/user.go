package usecase

import (
	"github.com/sandronister/cassandra-go/internal/entity"
	"github.com/sandronister/cassandra-go/internal/infra/database/repositories"
)

type UserUsecase struct {
	repo *repositories.UserRepository
}

func NewUserUsecase(repo *repositories.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) GetUsers() ([]*entity.User, error) {
	return u.repo.FindAll()
}
