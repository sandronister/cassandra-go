package entity

import "time"

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
