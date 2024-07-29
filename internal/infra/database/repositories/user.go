package repositories

import (
	"github.com/gocql/gocql"
	"github.com/sandronister/cassandra-go/internal/entity"
	"github.com/sandronister/cassandra-go/internal/infra/database/seed"
)

type UserRepository struct {
	session *gocql.Session
}

func NewUserRepository(session *gocql.Session) *UserRepository {
	return &UserRepository{session}
}

func (r *UserRepository) CreateUserTable(users []entity.User) error {
	for _, user := range users {
		if err := r.session.Query(seed.GetInsert(), user.ID, user.Name, user.Email, user.Password).Exec(); err != nil {
			return err
		}
	}
	return nil
}
