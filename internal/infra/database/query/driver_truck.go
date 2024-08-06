package query

import (
	"fmt"

	"github.com/sandronister/cassandra-go/internal/infra/database/migrations/tables"
)

const (
	insertDriverTruck    = `INSERT INTO %s.%s (company_id,license_id, name, vehicle_brand, vehicle_model, year, license_plate, created_at) VALUES (?,?, ?, ?, ?, ?, ?, ?);`
	selectDriverTruck    = `SELECT license_id, name, vehicle_brand, vehicle_model, year, license_plate, created_at, updated_at, deleted_at FROM %s.%s`
	selectByLicenseId    = `SELECT company_id, name, vehicle_brand, vehicle_model, year, license_plate, created_at, updated_at, deleted_at FROM %s.%s WHERE license_id = ?`
	selectByLicensePlate = `SELECT company_id, name, vehicle_brand, vehicle_model, year, license_plate, created_at, updated_at, deleted_at FROM %s.%s WHERE license_plate = ?`
	selectByCompanyId    = `SELECT license_id, name, vehicle_brand, vehicle_model, year, license_plate, created_at, updated_at, deleted_at FROM %s.%s WHERE company_id = ?`
)

var infoDriverTruck tables.InfoTable

func GetInsert() string {
	infoDriverTruck = *tables.GetDriverTruckMigration()
	return fmt.Sprintf(insertDriverTruck, infoDriverTruck.Keyspace, infoDriverTruck.Table)
}

func GetSelect() string {
	infoDriverTruck = *tables.GetDriverTruckMigration()
	return fmt.Sprintf(selectDriverTruck, infoDriverTruck.Keyspace, infoDriverTruck.Table)
}

func GetSelectByLicenseID() string {
	infoDriverTruck = *tables.GetDriverTruckMigration()
	return fmt.Sprintf(selectByLicenseId, infoDriverTruck.Keyspace, infoDriverTruck.Table)
}

func GetSelectByLicensePlate() string {
	infoDriverTruck = *tables.GetDriverTruckMigration()
	return fmt.Sprintf(selectByLicensePlate, infoDriverTruck.Keyspace, infoDriverTruck.Table)
}

func GetSelectByCompanyId() string {
	infoDriverTruck = *tables.GetDriverTruckMigration()
	return fmt.Sprintf(selectByCompanyId, infoDriverTruck.Keyspace, infoDriverTruck.Table)
}
