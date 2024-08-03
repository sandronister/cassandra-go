package entity

import "time"

type Driver struct {
	Name      string
	LicenseId string
	BirthDate time.Time
	City      string
	Phone     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
