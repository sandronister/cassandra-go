package repository

import "github.com/sandronister/cassandra-go/internal/entity"

type IDriver interface {
	Save(driver *entity.DriversTruck) error
	FindAll() ([]*entity.DriversTruck, error)
	FindById(id string) (*entity.DriversTruck, error)
	ExistsLicenseID(license string) bool
	Update(driver *entity.DriversTruck) error
	Delete(id string) error
}
