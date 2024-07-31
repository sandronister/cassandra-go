package migrations

import "fmt"

const (
	CreateKeySpace = `CREATE KEYSPACE IF NOT EXISTS %s WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };`
	CreateTable    = `CREATE TABLE IF NOT EXISTS %s.%s (%s);`
	CreateIndex    = `CREATE INDEX IF NOT EXISTS %s ON %s.%s (%s);`
)

func CreateKeyspace(keyspace string) string {
	return fmt.Sprintf(CreateKeySpace, keyspace)
}

func CreateTableQuery(keyspace, table, fields string) string {
	return fmt.Sprintf(CreateTable, keyspace, table, fields)
}

func CreateIndexQuery(indexName, keyspace, table, fields string) string {
	return fmt.Sprintf(CreateIndex, indexName, keyspace, table, fields)
}
