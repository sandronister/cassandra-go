package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sandronister/cassandra-go/internal/entity"
	"github.com/sandronister/cassandra-go/internal/repository"
	"golang.org/x/exp/rand"
)

var listDrivers = generateDrivers(50)

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

type Seed struct {
	repository repository.IDriverTruck
}

func NewSeed(repository repository.IDriverTruck) *Seed {
	return &Seed{repository}
}

func (s *Seed) CreateDrivers() error {

	list, _ := s.repository.FindAll()

	if list != nil {
		return nil
	}

	for _, driver := range listDrivers {
		if err := s.repository.Save(driver); err != nil {
			return err
		}
	}

	return nil
}
