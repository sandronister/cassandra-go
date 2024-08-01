package query

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sandronister/cassandra-go/internal/entity"
	"golang.org/x/exp/rand"
)

const (
	InsertDriverTruck = `INSERT INTO %s.%s (company_id,license_id, name, vehicle_brand, vehicle_model, year, license_plate, created_at) VALUES (?,?, ?, ?, ?, ?, ?, ?);`
	SelectDriverTruck = `SELECT license_id, name, vehicle_brand, vehicle_model, year, license_plate, created_at, updated_at, deleted_at FROM %s.%s`
)

var (
	Keyspace = "cassandra_go"
	Table    = "drivers_truck"
)

var ListDrivers = generateDrivers(50)

func generateDrivers(n int) []*entity.DriversTruck {
	listCompany := []string{"company1", "company2", "company3", "company4", "company5"}

	var drivers []*entity.DriversTruck
	for i := 1; i <= n; i++ {
		drivers = append(drivers, &entity.DriversTruck{
			CompanyId:    listCompany[rand.Intn(len(listCompany))],
			LicenseID:    uuid.New().String(),
			Name:         fmt.Sprintf("Driver%d", i),
			VehicleBrand: fmt.Sprintf("Brand%d", i),
			VehicleModel: fmt.Sprintf("Model%d", i),
			LicensePlate: fmt.Sprintf("ABC-%04d", i),
			Year:         rand.Intn(30) + 1990,
			CreatedAt:    time.Now(),
		})
	}

	return drivers
}

func GetInsert() string {
	return fmt.Sprintf(InsertDriverTruck, Keyspace, Table)
}

func GetSelect() string {
	return fmt.Sprintf(SelectDriverTruck, Keyspace, Table)
}

func GetDrivers() []*entity.DriversTruck {
	return ListDrivers
}
