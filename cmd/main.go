package main

import (
	"github.com/sandronister/cassandra-go/cmd/factory"
	"github.com/sandronister/cassandra-go/internal/di"
)

func main() {

	_, logger, db, err := factory.Initial()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	migrationsService := di.NewMigrationService(db)
	err = migrationsService.Run()
	logger.Error(err)

	seedService := di.NewSeedService(db)
	err = seedService.CreateDrivers()
	logger.Error(err)

	// driverHandler := di.NewDriverHandler(db)

	// server := server.NewServer(config.Port)
	// server.AddHDriverTruckHandler(driverHandler)
	// server.Run()

}
