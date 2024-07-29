package di

import (
	"github.com/gocql/gocql"
	"github.com/sandronister/cassandra-go/internal/infra/database/repositories"
	"github.com/sandronister/cassandra-go/internal/usecase"
)

func NewMigrationUseCase(db *gocql.Session) *usecase.MigrationsUsecase {
	repository := repositories.NewMigrationsRepository(db)
	return usecase.NewMigrationsUsecase(repository)
}
