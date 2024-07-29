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

	userUcase := di.NewUserUsecase(db)

	list, err := userUcase.GetUsers()

	if err != nil {
		panic(err)
	}

	for _, user := range list {
		println(user.ID, user.Name, user.Email, user.Password)
	}

}
