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

func NewDriversTruck(companyId, licenseID, name, vehicleBrand, vehicleModel, licensePlate string, year int) (*DriversTruck, error) {
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

	if !driverTruck.IsValid() {
		return nil, fmt.Errorf("invalid driver truck")
	}

	return driverTruck, nil
}

func (dt *DriversTruck) IsValid() bool {
	if dt.CompanyId == "" || dt.CompanyId == "0" {
		return false
	}

	if dt.LicenseID == "" || dt.LicenseID == "0" {
		return false
	}

	if dt.Name == "" {
		return false
	}

	if dt.VehicleBrand == "" {
		return false
	}

	if dt.VehicleModel == "" {
		return false
	}

	if dt.LicensePlate == "" {
		return false
	}

	if dt.Year == 0 {
		return false
	}

	return true
}
