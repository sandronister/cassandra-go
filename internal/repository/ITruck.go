package repository

import "github.com/sandronister/cassandra-go/internal/entity"

type ITruck interface {
	Save(truck *entity.Truck) error
	FindAll() ([]*entity.Truck, error)
	FindByPlate(plate string) (*entity.Truck, error)
	Update(truck *entity.Truck) error
	Delete(plate string) error
}
