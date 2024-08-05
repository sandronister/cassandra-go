package main

import (
	"github.com/sandronister/cassandra-go/cmd/factory"
	"github.com/sandronister/cassandra-go/internal/di"
	"github.com/sandronister/cassandra-go/internal/infra/web/server"
)

func main() {

	config, logger, db, err := factory.GetObjects()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	migrationsService := di.NewMigrationService(db)
	err = migrationsService.Run()
	logger.Info(err)

	seedService := di.NewSeedService(db)
	err = seedService.CreateDrivers()
	logger.Info(err)

	driverHandler := di.NewDriverHandler(db)

	server := server.NewServer(config.Port)
	server.AddHDriverTruckHandler(driverHandler)
	server.Run()

}
