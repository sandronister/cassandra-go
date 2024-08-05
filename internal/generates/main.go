package generates

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sandronister/cassandra-go/internal/entity"
	"golang.org/x/exp/rand"
)

var (
	listCompany = []string{"company1", "company2", "company3", "company4", "company5"}
)

func Company() []*entity.Company {
	var companies []*entity.Company

	for _, c := range listCompany {
		companies = append(companies, entity.NewCompany(c, "city", "uf"))
	}

	return companies
}

func Drivers(n int) []*entity.DriversTruck {
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
