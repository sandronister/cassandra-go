package usecase

import (
	"github.com/sandronister/cassandra-go/internal/entity"
	"github.com/sandronister/cassandra-go/internal/repository"
)

type DriverUseCase struct {
	repo repository.IDriverTruck
}

func NewDriverUseCase(repo repository.IDriverTruck) *DriverUseCase {
	return &DriverUseCase{repo: repo}
}

func (u *DriverUseCase) GetDrivers() ([]*entity.DriversTruck, error) {
	return u.repo.FindAll()
}
