package repository

import "github.com/sandronister/cassandra-go/internal/entity"

type ITruck interface {
	Save(truck *entity.DriversTruck) error
	FindAll() (*[]entity.DriversTruck, error)
	FindById(id string) (*entity.DriversTruck, error)
	ExistsLicensePlate(licensePlate string) bool
	Update(truck *entity.DriversTruck) error
	Delete(id string) error
}
