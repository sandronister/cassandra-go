package dto

type DriversTruck struct {
	CompanyId    string `json:"company_id"`
	Name         string `json:"name"`
	LicenseId    string `json:"license_id"`
	LicensePlate string `json:"license_plate"`
	VehicleBrand string `json:"vehicle_brand"`
	VehicleModel string `json:"vehicle_model"`
	VehicleYear  int    `json:"vehicle_year"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
