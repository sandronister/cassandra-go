package repositories

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/sandronister/cassandra-go/internal/entity"
	"github.com/sandronister/cassandra-go/internal/infra/database/query"
)

type Drivers struct {
	db *gocql.Session
}

func NewDriversRepository(db *gocql.Session) *Drivers {
	return &Drivers{
		db: db,
	}
}

func (d *Drivers) Save(driver *entity.Driver) error {
	err := d.db.Query(query.GetInsertDrivers(), driver.City, driver.BirthDate, driver.Email, driver.LicenseId, driver.Name, driver.Phone, driver.CreatedAt).Exec()

	if err != nil {
		return fmt.Errorf("[DRIVER] - error creating driver: %v", err)
	}

	return nil
}

func (d *Drivers) FindAll() ([]*entity.Driver, error) {
	var drivers []*entity.Driver

	iter := d.db.Query(query.GetSelectDrivers()).Iter()

	for {
		driver := &entity.Driver{}
		if !iter.Scan(&driver.City, &driver.BirthDate, &driver.Email, &driver.LicenseId, &driver.Name, &driver.Phone, &driver.CreatedAt) {
			break
		}
		drivers = append(drivers, driver)
	}

	return drivers, nil
}

func (d *Drivers) FindByLicense(licenseID string) (*entity.Driver, error) {
	var driver entity.Driver

	err := d.db.Query(query.GetDriverSelectByLicenseID(), licenseID).Scan(&driver.City, &driver.BirthDate, &driver.Email, &driver.LicenseId, &driver.Name, &driver.Phone, &driver.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &driver, nil
}

func (d *Drivers) Update(driver *entity.Driver) error {
	err := d.db.Query(query.GetUpdateDrivers(), driver.City, driver.BirthDate, driver.Email, driver.Name, driver.Phone, driver.LicenseId).Exec()

	if err != nil {
		return err
	}

	return nil
}

func (d *Drivers) Delete(licenseID string) error {
	err := d.db.Query(query.GetDeleteDrivers(), licenseID).Exec()

	if err != nil {
		return err
	}

	return nil
}
