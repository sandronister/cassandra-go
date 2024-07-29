package di

import (
	"github.com/gocql/gocql"
	"github.com/sandronister/cassandra-go/internal/infra/database/repositories"
	"github.com/sandronister/cassandra-go/internal/usecase"
)

func NewUserUsecase(session *gocql.Session) *usecase.UserUsecase {
	repository := repositories.NewUserRepository(session)
	return usecase.NewUserUsecase(repository)
}
