package services

import (
	"github.com/sandronister/cassandra-go/internal/entity"
	"github.com/sandronister/cassandra-go/internal/generates"
	"github.com/sandronister/cassandra-go/internal/repository"
	"golang.org/x/exp/rand"
)

type Seed struct {
	driverTruckRepo repository.IDriverTruck
	companyRepo     repository.ICompany
	driverRepo      repository.IDriver
	truckRepo       repository.ITruck
}

func NewSeed(driverTruckRepo repository.IDriverTruck, companyRepo repository.ICompany, driverRepo repository.IDriver, truckRepo repository.ITruck) *Seed {
	return &Seed{
		driverTruckRepo: driverTruckRepo,
		companyRepo:     companyRepo,
		driverRepo:      driverRepo,
		truckRepo:       truckRepo,
	}
}

func (s *Seed) CreateDrivers() ([]*entity.Driver, error) {

	list, _ := s.driverRepo.FindAll()

	if list != nil {
		return list, nil
	}

	drivers := generates.Drivers(50)
	for _, driver := range drivers {
		if err := s.driverRepo.Save(driver); err != nil {
			return nil, err
		}
	}

	return drivers, nil
}

func (s *Seed) CreateTrucks() ([]*entity.Truck, error) {
	list, _ := s.truckRepo.FindAll()

	if list != nil {
		return list, nil
	}

	trucks := generates.Trucks(50)
	for _, truck := range trucks {
		if err := s.truckRepo.Save(truck); err != nil {
			return nil, err
		}
	}

	return trucks, nil
}

func (s *Seed) CreateCompany() ([]*entity.Company, error) {
	list, _ := s.companyRepo.GetAll()

	if list != nil {
		return list, nil
	}

	companies := generates.Company()

	for _, company := range companies {
		if err := s.companyRepo.Save(company); err != nil {
			return nil, err
		}
	}

	return companies, nil
}

func (s *Seed) CreateDriverTruck(company entity.Company, driver entity.Driver, truck entity.Truck) error {
	driverTruck, err := entity.NewDriverTruck(company.Name, driver.LicenseId, driver.Name, truck.Brand, truck.Model, truck.Plate, truck.Year)

	if err != nil {
		return err
	}

	return s.driverTruckRepo.Save(driverTruck)
}

func (s *Seed) Run() error {
	companies, err := s.CreateCompany()

	if err != nil {
		return err
	}

	drivers, err := s.CreateDrivers()

	if err != nil {
		return err
	}

	trucks, err := s.CreateTrucks()

	if err != nil {
		return err
	}

	for i, driver := range drivers {
		err = s.CreateDriverTruck(*companies[rand.Intn(len(companies))], *driver, *trucks[i])

		if err != nil {
			return err
		}
	}

	return nil
}
