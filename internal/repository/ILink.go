package repository

import "github.com/sandronister/cassandra-go/internal/entity"

type ILink interface {
	Save(link *entity.DriversTruck) error
	ExistsLink(driverID, truckID string) bool
	Delete(driverID, truckID string) error
	GetTrucksByDriver(driverID string) (*entity.DriversTruck, error)
}
