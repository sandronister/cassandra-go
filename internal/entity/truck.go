package entity

import (
	"errors"
	"time"
)

var (
	ErrBrandIsRequired = errors.New("brand is required")
	ErrModelIsRequired = errors.New("model is required")
	ErrPlateIsRequired = errors.New("plate is required")
	ErrYearIsRequired  = errors.New("year is required")
)

type Truck struct {
	Brand     string
	Model     string
	Year      int
	Plate     string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

func NewTruck(brand, model, plate string, year int) (*Truck, error) {
	truck := &Truck{
		Brand:     brand,
		Model:     model,
		Year:      year,
		Plate:     plate,
		CreatedAt: time.Now().Format("2006-01-02"),
		UpdatedAt: "",
		DeletedAt: "",
	}

	err := truck.IsValid()

	if err != nil {
		return nil, err
	}

	return truck, nil

}

func (t *Truck) IsValid() error {
	if t.Brand == "" {
		return ErrBrandIsRequired
	}

	if t.Model == "" {
		return ErrModelIsRequired
	}

	if t.Plate == "" {
		return ErrPlateIsRequired
	}

	if t.Year == 0 {
		return ErrYearIsRequired
	}

	return nil
}
