package entity

import (
	"fmt"
	"time"
)

var (
	Keyspace   = "cassandra_go"
	Table      = "drivers_truck"
	Fields     = "company_id VARCHAR,license_id VARCHAR, name VARCHAR, vehicle_brand VARCHAR, vehicle_model VARCHAR, year INT, license_plate VARCHAR, created_at DATE, updated_at DATE, deleted_at DATE"
	PrimaryKey = "company_id,name"
	Orderby    = "name"
	Indexes    = map[string]string{
		"inx_plate": "license_plate",
		"inx_brand": "vehicle_brand",
	}
	ErrDriverTruckInvalid  = fmt.Errorf("invalid driver truck")
	ErrCompanyIDInvalid    = fmt.Errorf("invalid company id")
	ErrLicenseIDInvalid    = fmt.Errorf("invalid license id")
	ErrDriverNameInvalid   = fmt.Errorf("invalid driver name")
	ErrVehicleBrandInvalid = fmt.Errorf("invalid vehicle brand")
	ErrVehicleModelInvalid = fmt.Errorf("invalid vehicle model")
	ErrLicensePlateInvalid = fmt.Errorf("invalid license plate")
	ErrVehicleYearInvalid  = fmt.Errorf("invalid vehicle year")
)

type DriversTruck struct {
	CompanyId    string
	LicenseID    string
	Name         string
	VehicleBrand string
	VehicleModel string
	LicensePlate string
	Year         int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

func NewDriverTruck(companyId, licenseID, name, vehicleBrand, vehicleModel, licensePlate string, year int) (*DriversTruck, error) {
	driverTruck := &DriversTruck{
		CompanyId:    companyId,
		LicenseID:    licenseID,
		Name:         name,
		VehicleBrand: vehicleBrand,
		VehicleModel: vehicleModel,
		LicensePlate: licensePlate,
		Year:         year,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    time.Time{},
	}

	err := driverTruck.IsValid()

	if err != nil {
		return nil, err
	}

	return driverTruck, nil
}

func (dt *DriversTruck) IsValid() error {
	if dt.CompanyId == "" || dt.CompanyId == "0" {
		return ErrCompanyIDInvalid
	}

	if dt.LicenseID == "" || dt.LicenseID == "0" {
		return ErrLicenseIDInvalid
	}

	if dt.Name == "" {
		return ErrDriverNameInvalid
	}

	if dt.VehicleBrand == "" {
		return ErrVehicleBrandInvalid
	}

	if dt.VehicleModel == "" {
		return ErrVehicleModelInvalid
	}

	if dt.LicensePlate == "" {
		return ErrLicensePlateInvalid
	}

	if dt.Year == 0 {
		return ErrVehicleYearInvalid
	}

	return nil
}
