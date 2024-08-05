package services

import (
	"fmt"

	"github.com/sandronister/cassandra-go/internal/entity"
	"github.com/sandronister/cassandra-go/internal/infra/database/migrations/tables"
	"github.com/sandronister/cassandra-go/internal/infra/database/repositories"
)

type Migrations struct {
	repo       *repositories.MigrationsRepository
	infoTables []*tables.InfoTable
}

func NewMigrations(repo *repositories.MigrationsRepository) *Migrations {
	infoTables := tables.GetAllMigrations()
	return &Migrations{repo, infoTables}
}

func (s *Migrations) Run() error {

	if err := s.createKeyspace(); err != nil {
		return err
	}

	for _, infoTable := range s.infoTables {

		if err := s.createTable(infoTable); err != nil {
			return err
		}

		if err := s.createIndex(infoTable.Table, infoTable.Keyspace, infoTable.Indexes); err != nil {
			return err
		}

	}

	return nil
}

func (u *Migrations) createKeyspace() error {
	return u.repo.CreateKeyspace(entity.Keyspace)
}

func (u *Migrations) createTable(info *tables.InfoTable) error {
	return u.repo.CreateTable(info.Keyspace, info.Table, info.Fields, info.PrimaryKey, info.Orderby)
}

func (u *Migrations) createIndex(table, keyspace string, indexes map[string]string) error {
	for indexName, indexField := range indexes {
		if err := u.repo.CreateIndex(indexName, keyspace, table, indexField); err != nil {
			return fmt.Errorf("[TABLE] %s error creating index %s: %w", table, indexName, err)
		}
	}

	return nil
}
