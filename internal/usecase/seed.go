package usecase

import (
	"github.com/sandronister/cassandra-go/internal/infra/database/repositories"
	"github.com/sandronister/cassandra-go/internal/infra/database/seed"
)

type Seed struct {
	repository *repositories.UserRepository
}

func NewSeed(repository *repositories.UserRepository) *Seed {
	return &Seed{repository}
}

func (s *Seed) CreateUserTable() error {
	return s.repository.CreateUserTable(seed.GetUsers())
}
