package repository

import "github.com/sandronister/cassandra-go/internal/entity"

type IDriverTruck interface {
	Save(driver *entity.DriversTruck) error
	FindAll() ([]*entity.DriversTruck, error)
	FindByLicenseId(license_id string) (*entity.DriversTruck, error)
	FindByLicensePlate(licensePlate string) (*entity.DriversTruck, error)
	FindByCompanyId(companyId string) ([]*entity.DriversTruck, error)
}
