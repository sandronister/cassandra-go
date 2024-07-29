package main

import (
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

	usecase := di.NewMigrationUseCase(db)

	err = usecase.Run()

	if err != nil {
		panic(err)
	}
}
