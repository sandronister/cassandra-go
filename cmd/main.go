package main

import (
	"fmt"
	"log"

	configs "github.com/sandronister/cassandra-go/config"
	"github.com/sandronister/cassandra-go/internal/di"
	"github.com/sandronister/cassandra-go/internal/infra/database/connection"
)

func main() {
	conf, err := configs.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	db, err := connection.NewCassandraConnection(conf)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	migrationUseCase := di.NewMigrationUseCase(db)
	err = migrationUseCase.Run()

	if err != nil {
		panic(err)
	}

	// seedUseCase := di.NewSeedUsecase(db)

	// err = seedUseCase.CreateDrivers()

	// if err != nil {
	// 	panic(err)
	// }

	driverUseCase := di.NewDriverUseCase(db)

	list, err := driverUseCase.GetDrivers()

	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	for _, driver := range list {
		fmt.Printf("Driver License ID: %v\n", driver.LicenseID)
		fmt.Printf("Driver Name: %v\n", driver.Name)
		fmt.Printf("Vehicle Brand: %v\n", driver.VehicleBrand)
		fmt.Printf("Vehicle Model: %v\n", driver.VehicleModel)
		fmt.Printf("License Plate: %v\n", driver.LicensePlate)
		fmt.Printf("Year: %v\n", driver.Year)
		fmt.Printf("Created At: %v\n", driver.CreatedAt)
		fmt.Println("=====================================")
	}

}
