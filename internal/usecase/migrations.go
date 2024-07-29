package usecase

import "github.com/sandronister/cassandra-go/internal/infra/database/repositories"

var (
	keyspace = "cassandra_go"
	table    = "users"
	fields   = "name TEXT, email TEXT"
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

	return nil
}

func (u *MigrationsUsecase) createKeyspace() error {
	return u.repo.CreateKeyspace(keyspace)
}

func (u *MigrationsUsecase) createTable() error {
	return u.repo.CreateTable(keyspace, table, fields)
}

func (u *MigrationsUsecase) dropTable() error {
	return u.repo.DropTable(keyspace, table)
}
