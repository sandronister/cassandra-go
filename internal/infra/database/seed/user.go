package seed

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sandronister/cassandra-go/internal/entity"
)

const (
	InsertUser = `INSERT INTO %s.%s (id, name, email, password) VALUES (?, ?, ?, ?);`
)

var (
	Keyspace = "cassandra_go"
	Table    = "users"
	Fields   = "name TEXT, email TEXT, password TEXT"
)

var ListUser = []struct {
	Name     string
	Email    string
	Password string
}{
	{Name: "Sandro", Email: "sandro@ig.com.br", Password: "123456"},
	{Name: "João", Email: "joao@uol.com.br", Password: "123456"},
	{Name: "Maria", Email: "maria@ig.com.br", Password: "123456"},
}

func GetInsert() string {
	return fmt.Sprintf(InsertUser, Keyspace, Table)
}

func GetUsers() []entity.User {
	var users []entity.User

	for _, user := range ListUser {
		id := uuid.New().String()
		users = append(users, entity.User{
			ID:       id,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		})
	}

	return users
}
