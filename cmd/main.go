package main

import (
	"fmt"

	"github.com/sandronister/cassandra-go/internal/di"
)

func main() {

	logger, db, err := initialFactory()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	migrationUseCase := di.NewMigrationUseCase(db)
	err = migrationUseCase.Run()

	if err != nil {
		logger.Info(fmt.Sprintf("Error: %v", err))
	}

	seedService := di.NewSeedService(db)

	err = seedService.CreateDrivers()

	if err != nil {
		logger.Info(fmt.Sprintf("Error: %v", err))
	}

	driverUseCase := di.NewDriverUseCase(db)

}
