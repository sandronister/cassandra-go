package repositories

import (
	"github.com/gocql/gocql"
	"github.com/sandronister/cassandra-go/internal/entity"
	"github.com/sandronister/cassandra-go/internal/infra/database/query"
)

type DriverTruckRepository struct {
	session *gocql.Session
}

func NewUserRepository(session *gocql.Session) *DriverTruckRepository {
	return &DriverTruckRepository{session}
}

func (r *DriverTruckRepository) Save(driver *entity.DriversTruck) error {
	if err := r.session.Query(query.GetInsert(), driver.CompanyId, driver.LicenseID, driver.Name, driver.VehicleBrand, driver.VehicleModel, driver.Year, driver.LicensePlate, driver.CreatedAt).Exec(); err != nil {
		return err
	}
	return nil
}

func (r *DriverTruckRepository) FindAll() ([]*entity.DriversTruck, error) {
	var users []*entity.DriversTruck
	iter := r.session.Query(query.GetSelect()).Iter()
	for {
		var driver entity.DriversTruck
		if !iter.Scan(&driver.LicenseID, &driver.Name, &driver.VehicleBrand, &driver.VehicleModel, &driver.Year, &driver.LicensePlate, &driver.CreatedAt, &driver.UpdatedAt, &driver.DeletedAt) {
			break
		}
		if driver.DeletedAt.IsZero() {
			users = append(users, &driver)
		}
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *DriverTruckRepository) FindByLicenseId(licenseID string) (*entity.DriversTruck, error) {
	var driver entity.DriversTruck
	if err := r.session.Query(query.GetSelectByLicenseID(), licenseID).Scan(&driver.CompanyId, &driver.Name, &driver.VehicleBrand, &driver.VehicleModel, &driver.Year, &driver.LicensePlate, &driver.CreatedAt, &driver.UpdatedAt, &driver.DeletedAt); err != nil {
		return nil, err
	}
	return &driver, nil
}

func (r *DriverTruckRepository) FindByLicensePlate(licensePlate string) (*entity.DriversTruck, error) {
	var driver entity.DriversTruck
	if err := r.session.Query(query.GetSelectByLicensePlate(), licensePlate).Scan(&driver.CompanyId, &driver.Name, &driver.VehicleBrand, &driver.VehicleModel, &driver.Year, &driver.LicensePlate, &driver.CreatedAt, &driver.UpdatedAt, &driver.DeletedAt); err != nil {
		return nil, err
	}
	return &driver, nil
}

func (r *DriverTruckRepository) FindByCompanyId(companyId string) ([]*entity.DriversTruck, error) {
	var drivers []*entity.DriversTruck
	iter := r.session.Query(query.GetSelectByCompanyId(), companyId).Iter()
	for {
		var driver entity.DriversTruck
		if !iter.Scan(&driver.LicenseID, &driver.Name, &driver.VehicleBrand, &driver.VehicleModel, &driver.Year, &driver.LicensePlate, &driver.CreatedAt, &driver.UpdatedAt, &driver.DeletedAt) {
			break
		}
		if driver.DeletedAt.IsZero() {
			drivers = append(drivers, &driver)
		}
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}
	return drivers, nil
}
