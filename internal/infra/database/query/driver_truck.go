package query

import (
	"fmt"
)

const (
	insertDriverTruck    = `INSERT INTO %s.%s (company_id,license_id, name, vehicle_brand, vehicle_model, year, license_plate, created_at) VALUES (?,?, ?, ?, ?, ?, ?, ?);`
	selectDriverTruck    = `SELECT license_id, name, vehicle_brand, vehicle_model, year, license_plate, created_at, updated_at, deleted_at FROM %s.%s`
	selectByLicenseId    = `SELECT company_id, name, vehicle_brand, vehicle_model, year, license_plate, created_at, updated_at, deleted_at FROM %s.%s WHERE license_id = ?`
	selectByLicensePlate = `SELECT company_id, name, vehicle_brand, vehicle_model, year, license_plate, created_at, updated_at, deleted_at FROM %s.%s WHERE license_plate = ?`
	selectByCompanyId    = `SELECT license_id, name, vehicle_brand, vehicle_model, year, license_plate, created_at, updated_at, deleted_at FROM %s.%s WHERE company_id = ?`
)

var (
	Keyspace = "cassandra_go"
	Table    = "drivers_truck"
)

func GetInsert() string {
	return fmt.Sprintf(insertDriverTruck, Keyspace, Table)
}

func GetSelect() string {
	return fmt.Sprintf(selectDriverTruck, Keyspace, Table)
}

func GetSelectByLicenseID() string {
	return fmt.Sprintf(selectByLicenseId, Keyspace, Table)
}

func GetSelectByLicensePlate() string {
	return fmt.Sprintf(selectByLicensePlate, Keyspace, Table)
}

func GetSelectByCompanyId() string {
	return fmt.Sprintf(selectByCompanyId, Keyspace, Table)
}
