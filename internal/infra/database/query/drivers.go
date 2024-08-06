package query

import (
	"fmt"

	"github.com/sandronister/cassandra-go/internal/infra/database/migrations/tables"
)

const (
	insertDrivers           = "INSERT INTO %s.%s (city,birth_date,email,license_id,name,phone,created_at) VALUES (?,?,?,?,?,?,?)"
	selectDrivers           = "SELECT city,birth_date,email,license_id,name,phone,created_at FROM %s.%s"
	selectDriverByLicenseId = "SELECT city,birth_date,email,license_id,name,phone,created_at FROM %s.%s WHERE license_id = ?"
	updateDrivers           = "UPDATE %s.%s SET city = ?, birth_date = ?, email = ?, name = ?, phone = ? WHERE license_id = ?"
	deleteDrivers           = "DELETE FROM %s.%s WHERE license_id = ?"
)

var infoDriver tables.InfoTable

func GetInsertDrivers() string {
	infoDriver = *tables.GetDriversMigration()
	return fmt.Sprintf(insertDrivers, infoDriver.Keyspace, infoDriver.Table)
}

func GetSelectDrivers() string {
	infoDriver = *tables.GetDriversMigration()
	return fmt.Sprintf(selectDrivers, infoDriver.Keyspace, infoDriver.Table)
}

func GetDriverSelectByLicenseID() string {
	infoDriver = *tables.GetDriversMigration()
	return fmt.Sprintf(selectDriverByLicenseId, infoDriver.Keyspace, infoDriver.Table)
}

func GetUpdateDrivers() string {
	infoDriver = *tables.GetDriversMigration()
	return fmt.Sprintf(updateDrivers, infoDriver.Keyspace, infoDriver.Table)
}

func GetDeleteDrivers() string {
	infoDriver = *tables.GetDriversMigration()
	return fmt.Sprintf(deleteDrivers, infoDriver.Keyspace, infoDriver.Table)
}
