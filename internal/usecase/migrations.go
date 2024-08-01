package usecase

import (
	"github.com/sandronister/cassandra-go/internal/entity"
	"github.com/sandronister/cassandra-go/internal/infra/database/repositories"
)

type MigrationsUsecase struct {
	repo *repositories.MigrationsRepository
}

func NewMigrationsUsecase(repo *repositories.MigrationsRepository) *MigrationsUsecase {
	return &MigrationsUsecase{repo}
}

func (u *MigrationsUsecase) Run() error {
	if err := u.createKeyspace(); err != nil {
		return err
	}

	if err := u.createTable(); err != nil {
		return err
	}

	if err := u.createIndex(); err != nil {
		return err
	}

	return nil
}

func (u *MigrationsUsecase) createKeyspace() error {
	return u.repo.CreateKeyspace(entity.Keyspace)
}

func (u *MigrationsUsecase) createTable() error {
	return u.repo.CreateTable(entity.Keyspace, entity.Table, entity.Fields, entity.PrimaryKey, entity.Orderby)
}

func (u *MigrationsUsecase) createIndex() error {
	for indexName, indexField := range entity.Indexes {
		if err := u.repo.CreateIndex(indexName, entity.Keyspace, entity.Table, indexField); err != nil {
			return err
		}
	}

	return nil
}
