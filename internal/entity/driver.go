package entity

import (
	"errors"
	"time"
)

var (
	ErrInvalidEntity = errors.New("invalid entity")
	ErrCityInvalid   = errors.New("invalid city")
	ErrPhoneInvalid  = errors.New("invalid phone")
	ErrEmailInvalid  = errors.New("invalid email")
)

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

func NewDriver(name, licenseId, city, phone, email string, birthDate time.Time) (*Driver, error) {
	driver := &Driver{
		Name:      name,
		LicenseId: licenseId,
		BirthDate: birthDate,
		City:      city,
		Phone:     phone,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Time{},
	}

	err := driver.IsValid()

	if err != nil {
		return nil, err
	}

	return driver, nil
}

func (d *Driver) IsValid() error {
	if d.Name == "" {
		return ErrDriverNameInvalid
	}

	if d.LicenseId == "" {
		return ErrLicenseIDInvalid
	}

	if d.City == "" {
		return ErrCityInvalid
	}

	if d.Phone == "" {
		return ErrPhoneInvalid
	}

	if d.Email == "" {
		return ErrEmailInvalid
	}

	return nil

}
