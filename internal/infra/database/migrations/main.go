package migrations

import "fmt"

const (
	CreateKeySpace = `CREATE KEYSPACE IF NOT EXISTS %s WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };`
	CreateTable    = `CREATE TABLE IF NOT EXISTS %s.%s (id UUID PRIMARY KEY, %s);`
	DropTable      = `DROP TABLE IF EXISTS %s.%s;`
)

func CreateKeyspace(keyspace string) string {
	return fmt.Sprintf(CreateKeySpace, keyspace)
}

func CreateTableQuery(keyspace, table, fields string) string {
	return fmt.Sprintf(CreateTable, keyspace, table, fields)
}

func DropTableQuery(keyspace, table string) string {
	return fmt.Sprintf(DropTable, keyspace, table)
}
