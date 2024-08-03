package repository

import "github.com/sandronister/cassandra-go/internal/entity"

type IDriver interface {
	Save(driver *entity.Driver) error
	GetAll() ([]*entity.Driver, error)
	GetByCompanyID(companyID string) ([]*entity.Driver, error)
	GetByLicenseID(licenseID string) (*entity.Driver, error)
	Update(driver *entity.Driver) error
	Delete(licenseID string) error
}
