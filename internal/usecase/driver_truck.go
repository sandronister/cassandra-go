package usecase

import (
	"github.com/sandronister/cassandra-go/internal/entity"
	"github.com/sandronister/cassandra-go/internal/infra/database/repositories"
)

type DriverUseCase struct {
	repo *repositories.DriverTruckRepository
}

func NewDriverUseCase(repo *repositories.DriverTruckRepository) *DriverUseCase {
	return &DriverUseCase{repo: repo}
}

func (u *DriverUseCase) GetDrivers() ([]*entity.DriversTruck, error) {
	return u.repo.FindAll()
}
