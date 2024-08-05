package repository

import "github.com/sandronister/cassandra-go/internal/entity"

type IDriver interface {
	Save(driver *entity.Driver) error
	FindAll() ([]*entity.Driver, error)
	FindByLicense(licenseID string) (*entity.Driver, error)
	Update(driver *entity.Driver) error
	Delete(licenseID string) error
}
