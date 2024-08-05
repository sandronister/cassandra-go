package main

import (
	"fmt"

	"github.com/gocql/gocql"
	configs "github.com/sandronister/cassandra-go/config"
	"github.com/sandronister/cassandra-go/internal/infra/database/connection"
	"github.com/sandronister/cassandra-go/pkg/logger"
)

func initialFactory() (*logger.Logger, *gocql.Session, error) {
	conf, err := configs.LoadConfig(".")

	if err != nil {
		return nil, nil, fmt.Errorf("Error loading config: %v", err)
	}

	logger, err := logger.NewLogger(conf.LoggerFolder)

	if err != nil {
		return nil, nil, fmt.Errorf("Error loading logger: %v", err)
	}

	db, err := connection.NewCassandraConnection(conf)

	if err != nil {
		return nil, nil, fmt.Errorf("Error loading database: %v", err)
	}

	return logger, db, nil

}
