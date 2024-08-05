package repositories

import (
	"github.com/gocql/gocql"
	"github.com/sandronister/cassandra-go/internal/entity"
)

type CompanyRepository struct {
	session *gocql.Session
}

func NewCompanyRepository(session *gocql.Session) *CompanyRepository {
	return &CompanyRepository{session}
}

func (r *CompanyRepository) Save(company *entity.Company) error {
	return r.session.Query(`INSERT INTO company (name, city, uf, created_at) VALUES (?, ?, ?, ?)`,
		company.Name, company.City, company.Uf, company.CreatedAt).Exec()
}

func (r *CompanyRepository) GetAll() ([]*entity.Company, error) {
	var companies []*entity.Company
	iter := r.session.Query(`SELECT name, city, uf, created_at FROM company`).Iter()
	for {
		var company entity.Company
		if !iter.Scan(&company.Name, &company.City, &company.Uf, &company.CreatedAt) {
			break
		}
		companies = append(companies, &company)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}
	return companies, nil
}

func (r *CompanyRepository) GetByID(companyID string) (*entity.Company, error) {
	var company entity.Company
	if err := r.session.Query(`SELECT name, city, uf, created_at FROM company WHERE name = ?`, companyID).Scan(&company.Name, &company.City, &company.Uf, &company.CreatedAt); err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *CompanyRepository) Update(company *entity.Company) error {
	return r.session.Query(`UPDATE company SET city = ?, uf = ? WHERE name = ?`, company.City, company.Uf, company.Name).Exec()
}

func (r *CompanyRepository) Delete(companyID string) error {
	return r.session.Query(`DELETE FROM company WHERE name = ?`, companyID).Exec()
}
