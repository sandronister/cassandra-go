package tables

type InfoTable struct {
	Keyspace   string
	Table      string
	Fields     string
	PrimaryKey string
	Orderby    string
	Indexes    map[string]string
}

func GetDriverTruckMigration() *InfoTable {
	return &InfoTable{
		Keyspace:   "cassandra_go",
		Table:      "drivers_truck",
		Fields:     "company VARCHAR,license_id VARCHAR, name VARCHAR, vehicle_brand VARCHAR, vehicle_model VARCHAR, year INT, license_plate VARCHAR, created_at DATE, updated_at DATE, deleted_at DATE",
		PrimaryKey: "company,year",
		Orderby:    "year",
		Indexes: map[string]string{
			"inx_plate": "license_plate",
			"inx_brand": "vehicle_brand",
		},
	}
}

func GetDriversMigration() *InfoTable {
	return &InfoTable{
		Keyspace:   "cassandra_go",
		Table:      "drivers",
		Fields:     "name VARCHAR, license_id VARCHAR,birth_date DATE, city VARCHAR, phone VARCHAR, email VARCHAR, created_at DATE, updated_at DATE, deleted_at DATE",
		PrimaryKey: "company,city",
		Orderby:    "city",
		Indexes: map[string]string{
			"inx_phone": "phone",
		},
	}
}

func GetTruckMigrations() *InfoTable {
	return &InfoTable{
		Keyspace:   "cassandra_go",
		Table:      "trucks",
		Fields:     "brand VARCHAR, model VARCHAR, year INT, plate VARCHAR, created_at DATE, updated_at DATE, deleted_at DATE",
		PrimaryKey: "brand,plate",
		Orderby:    "plate",
		Indexes: map[string]string{
			"inx_brand": "brand",
			"inx_plate": "plate",
		},
	}
}

func GetCompanyMigrations() *InfoTable {
	return &InfoTable{
		Keyspace:   "cassandra_go",
		Table:      "companies",
		Fields:     "name VARCHAR, city VARCHAR, uf VARCHAR,created_at DATE, updated_at DATE, deleted_at DATE",
		PrimaryKey: "name,city",
		Orderby:    "city",
		Indexes: map[string]string{
			"inx_uf": "uf",
		},
	}
}

func GetAllMigrations() []*InfoTable {
	return []*InfoTable{
		GetDriverTruckMigration(),
		GetDriversMigration(),
		GetTruckMigrations(),
		GetCompanyMigrations(),
	}
}
