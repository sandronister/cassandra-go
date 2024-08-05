package repositories

import (
	"github.com/gocql/gocql"
	"github.com/sandronister/cassandra-go/internal/entity"
)

type Truck struct {
	db *gocql.Session
}

func NewTruckRepository(db *gocql.Session) *Truck {
	return &Truck{
		db: db,
	}
}

func (t *Truck) Save(truck *entity.Truck) error {
	err := t.db.Query("INSERT INTO trucks (vehicle_brand,vehicle_model,plate,year,created_at) VALUES (?,?,?,?,?)", truck.Brand, truck.Model, truck.Plate, truck.Year, truck.CreatedAt).Exec()

	if err != nil {
		return err
	}

	return nil
}

func (t *Truck) FindAll() ([]*entity.Truck, error) {
	var trucks []*entity.Truck

	iter := t.db.Query("SELECT vehicle_brand,vehicle_model,license_plate,year,created_at FROM trucks").Iter()

	for {
		truck := &entity.Truck{}
		if !iter.Scan(&truck.Brand, &truck.Model, &truck.Plate, &truck.Year, &truck.CreatedAt) {
			break
		}
		trucks = append(trucks, truck)
	}

	return trucks, nil
}

func (t *Truck) FindByPlate(plate string) (*entity.Truck, error) {
	var truck entity.Truck

	err := t.db.Query("SELECT vehicle_brand,vehicle_model,license_plate,year,created_at FROM trucks WHERE license_plate = ?", plate).Scan(&truck.Brand, &truck.Model, &truck.Plate, &truck.Year, &truck.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &truck, nil
}

func (t *Truck) Update(truck *entity.Truck) error {
	err := t.db.Query("UPDATE trucks SET vehicle_brand = ?, vehicle_model = ?, year = ? WHERE license_plate = ?", truck.Brand, truck.Model, truck.Year, truck.Plate).Exec()

	if err != nil {
		return err
	}

	return nil
}

func (t *Truck) Delete(plate string) error {
	err := t.db.Query("DELETE FROM trucks WHERE license_plate = ?", plate).Exec()

	if err != nil {
		return err
	}

	return nil
}

func (t *Truck) FindByBrand(brand string) (*entity.Truck, error) {
	var truck entity.Truck

	err := t.db.Query("SELECT vehicle_brand,vehicle_model,license_plate,year,created_at FROM trucks WHERE vehicle_brand = ?", brand).Scan(&truck.Brand, &truck.Model, &truck.Plate, &truck.Year, &truck.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &truck, nil
}
