package services

import (
	"github.com/sandronister/cassandra-go/internal/generates"
	"github.com/sandronister/cassandra-go/internal/repository"
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

func (s *Seed) CreateDrivers() error {

	list, _ := s.driverRepo.FindAll()

	if list != nil {
		return nil
	}

	listDrivers := generates.Drivers(50)
	for _, driver := range listDrivers {
		if err := s.driverRepo.Save(driver); err != nil {
			return err
		}
	}

	return nil
}

func (s *Seed) CreateTrucks() error {
	list, _ := s.driverTruckRepo.FindAll()

	if list != nil {
		return nil
	}

	listTrucks := generates.Trucks(50)
	for _, truck := range listTrucks {
		if err := s.driverTruckRepo.Save(truck); err != nil {
			return err
		}
	}

	return nil
}

func (s *Seed) CreateCompany() error {
	list, _ := s.companyRepo.GetAll()

	if list != nil {
		return nil
	}

	companies := generates.Company()

	for _, company := range companies {
		if err := s.companyRepo.Save(company); err != nil {
			return err
		}
	}

	return nil
}

func (s *Seed) Run() error {
	if err := s.CreateCompany(); err != nil {
		return err
	}

	if err := s.CreateDrivers(); err != nil {
		return err
	}

	if err := s.CreateTrucks(); err != nil {
		return err
	}

	return nil
}
