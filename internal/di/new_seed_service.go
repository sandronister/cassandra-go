package di

import (
	"github.com/gocql/gocql"
	"github.com/sandronister/cassandra-go/internal/infra/database/repositories"
	"github.com/sandronister/cassandra-go/internal/services"
)

func NewSeedService(session *gocql.Session) *services.Seed {
	driverTruckRepo := repositories.NewDriverTruckRepository(session)
	companyRepo := repositories.NewCompanyRepository(session)
	return services.NewSeed(driverTruckRepo, companyRepo)
}
