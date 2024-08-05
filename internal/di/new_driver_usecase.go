package di

import (
	"github.com/gocql/gocql"
	"github.com/sandronister/cassandra-go/internal/infra/database/repositories"
	"github.com/sandronister/cassandra-go/internal/usecase"
)

func NewDriverUseCase(session *gocql.Session) *usecase.DriverUseCase {
	repository := repositories.NewDriverTruckRepository(session)
	return usecase.NewDriverUseCase(repository)
}
