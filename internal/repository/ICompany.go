package repository

import "github.com/sandronister/cassandra-go/internal/entity"

type ICompany interface {
	Save(company *entity.Company) error
	GetAll() ([]*entity.Company, error)
	GetByID(companyID string) (*entity.Company, error)
	Update(company *entity.Company) error
	Delete(companyID string) error
}
