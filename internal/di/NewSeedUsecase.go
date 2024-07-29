package di

import (
	"github.com/gocql/gocql"
	"github.com/sandronister/cassandra-go/internal/infra/database/repositories"
	"github.com/sandronister/cassandra-go/internal/usecase"
)

func NewSeedUsecase(session *gocql.Session) *usecase.Seed {
	repository := repositories.NewUserRepository(session)
	return usecase.NewSeed(repository)
}
