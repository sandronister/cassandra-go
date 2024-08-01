package repositories

import (
	"github.com/gocql/gocql"
	"github.com/sandronister/cassandra-go/internal/infra/database/migrations"
)

type MigrationsRepository struct {
	session *gocql.Session
}

func NewMigrationsRepository(session *gocql.Session) *MigrationsRepository {
	return &MigrationsRepository{session}
}

func (r *MigrationsRepository) CreateKeyspace(keyspace string) error {
	return r.session.Query(migrations.CreateKeyspace(keyspace)).Exec()
}

func (r *MigrationsRepository) CreateTable(keyspace, table, fields, primaryKey, orderby string) error {
	return r.session.Query(migrations.CreateTableQuery(keyspace, table, fields, primaryKey, orderby)).Exec()
}

func (r *MigrationsRepository) CreateIndex(indexName, keyspace, table, fields string) error {
	return r.session.Query(migrations.CreateIndexQuery(indexName, keyspace, table, fields)).Exec()
}
