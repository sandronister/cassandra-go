package di

import (
	"github.com/gocql/gocql"
	"github.com/sandronister/cassandra-go/internal/infra/database/repositories"
	"github.com/sandronister/cassandra-go/internal/services"
)

func NewMigrationService(db *gocql.Session) *services.Migrations {
	repository := repositories.NewMigrationsRepository(db)
	return services.NewMigrations(repository)
}
