package usecase

import (
	"github.com/sandronister/cassandra-go/internal/infra/database/query"
	"github.com/sandronister/cassandra-go/internal/infra/database/repositories"
)

type Seed struct {
	repository *repositories.DriverTruckRepository
}

func NewSeed(repository *repositories.DriverTruckRepository) *Seed {
	return &Seed{repository}
}

func (s *Seed) CreateDrivers() error {
	list := query.GetDrivers()

	for _, driver := range list {
		if err := s.repository.Save(driver); err != nil {
			return err
		}
	}

	return nil
}
