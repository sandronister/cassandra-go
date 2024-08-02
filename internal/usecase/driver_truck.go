package usecase

import (
	"database/sql"
	"time"

	"github.com/sandronister/cassandra-go/internal/entity"
	"github.com/sandronister/cassandra-go/internal/infra/web/dto"
	"github.com/sandronister/cassandra-go/internal/repository"
)

type DriverUseCase struct {
	repo repository.IDriverTruck
}

func NewDriverUseCase(repo repository.IDriverTruck) *DriverUseCase {
	return &DriverUseCase{repo: repo}
}

func (u *DriverUseCase) GetDrivers() ([]*dto.DriversTruck, error) {
	list, err := u.repo.FindAll()

	result := make([]*dto.DriversTruck, 0)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, nil
		}
		return nil, err
	}

	for _, item := range list {
		if item.DeletedAt.IsZero() {
			result = append(result, &dto.DriversTruck{
				CompanyId:    item.CompanyId,
				Name:         item.Name,
				LicenseId:    item.LicenseID,
				LicensePlate: item.LicensePlate,
				VehicleBrand: item.VehicleBrand,
				VehicleModel: item.VehicleModel,
				VehicleYear:  item.Year,
				CreatedAt:    item.CreatedAt.String(),
				UpdatedAt:    item.UpdatedAt.String(),
			})
		}
	}

	return result, nil
}

func (u *DriverUseCase) GetDriver(license_id string) (*dto.DriversTruck, error) {
	item, err := u.repo.FindByLicenseId(license_id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, nil
		}
		return nil, err
	}

	if item.DeletedAt.IsZero() {
		return &dto.DriversTruck{
			CompanyId:    item.CompanyId,
			Name:         item.Name,
			LicenseId:    item.LicenseID,
			LicensePlate: item.LicensePlate,
			VehicleBrand: item.VehicleBrand,
			VehicleModel: item.VehicleModel,
			VehicleYear:  item.Year,
			CreatedAt:    item.CreatedAt.String(),
			UpdatedAt:    item.UpdatedAt.String(),
		}, nil
	}

	return nil, nil
}

func (u *DriverUseCase) CreateDriver(driver *dto.DriversTruck) error {
	driverTruck, err := entity.NewDriverTruck(driver.CompanyId, driver.Name, driver.LicenseId, driver.LicensePlate, driver.VehicleBrand, driver.VehicleModel, driver.VehicleYear)

	if err != nil {
		return err
	}

	return u.repo.Save(driverTruck)
}

func (u *DriverUseCase) UpdateDriver(driver *dto.DriversTruck) error {
	driverTruck, err := u.repo.FindByLicenseId(driver.LicenseId)

	if err != nil {
		return err
	}

	if driver.CompanyId != "" {
		driverTruck.CompanyId = driver.CompanyId
	}

	if driver.Name != "" {
		driverTruck.Name = driver.Name
	}

	if driver.LicensePlate != "" {
		driverTruck.LicensePlate = driver.LicensePlate
	}

	if driver.VehicleBrand != "" {
		driverTruck.VehicleBrand = driver.VehicleBrand
	}

	if driver.VehicleModel != "" {
		driverTruck.VehicleModel = driver.VehicleModel
	}

	if driver.VehicleYear != 0 {
		driverTruck.Year = driver.VehicleYear
	}

	return u.repo.Save(driverTruck)

}

func (u *DriverUseCase) DeleteDriver(license_id string) error {
	driverTruck, err := u.repo.FindByLicenseId(license_id)
	if err != nil {
		return err
	}

	driverTruck.DeletedAt = time.Now()
	return u.repo.Save(driverTruck)
}
