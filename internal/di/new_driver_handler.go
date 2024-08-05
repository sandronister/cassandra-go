package di

import (
	"github.com/gocql/gocql"
	"github.com/sandronister/cassandra-go/internal/infra/database/repositories"
	"github.com/sandronister/cassandra-go/internal/infra/web/handler"
	"github.com/sandronister/cassandra-go/internal/usecase"
)

func NewDriverHandler(db *gocql.Session) *handler.DriverTruckHandler {
	repository := repositories.NewDriverTruckRepository(db)
	usecase := usecase.NewDriverUseCase(repository)
	return handler.NewDriverTruckHandler(usecase)
}
