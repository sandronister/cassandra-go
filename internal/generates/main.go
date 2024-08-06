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
	listBrand   = []string{"Ford", "Volvo", "Scania", "Mercedes", "DAF"}
)

func Company() []*entity.Company {
	var companies []*entity.Company

	for _, c := range listCompany {
		companies = append(companies, entity.NewCompany(c, "city", "uf"))
	}

	return companies
}

func Drivers(n int) []*entity.Driver {
	var drivers []*entity.Driver
	for i := 1; i <= n; i++ {
		drivers = append(drivers, &entity.Driver{
			Name:      fmt.Sprintf("Driver%d", i),
			LicenseId: uuid.New().String(),
			BirthDate: time.Now(),
			City:      "city",
			Phone:     "phone",
			Email:     "email",
			CreatedAt: time.Now(),
		})
	}

	return drivers
}

func Trucks(n int) []*entity.Truck {
	var trucks []*entity.Truck
	for i := 1; i <= n; i++ {
		trucks = append(trucks, &entity.Truck{
			Brand:     listBrand[rand.Intn(len(listBrand))],
			Model:     fmt.Sprintf("Model%d", i),
			Year:      2021,
			Plate:     fmt.Sprintf("Plate%d", i),
			CreatedAt: time.Now(),
		})
	}

	return trucks
}
