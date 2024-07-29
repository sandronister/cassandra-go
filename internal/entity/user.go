package entity

var (
	Keyspace = "cassandra_go"
	Table    = "users"
	Fields   = "name TEXT, email TEXT, password TEXT"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}
