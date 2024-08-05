package entity

import "time"

type Company struct {
	Name      string
	City      string
	Uf        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

func NewCompany(name, city, uf string) *Company {
	return &Company{
		Name:      name,
		City:      city,
		Uf:        uf,
		CreatedAt: time.Now().Format("2006-01-02"),
	}
}
