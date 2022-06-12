package postgres

import (
	"database/sql"

	"github.com/cornejodev/viator/internal/domain"
)

type VehicleRepository struct {
	db *sql.DB
}

func NewVehicleRepository(db *sql.DB) *VehicleRepository {
	return &VehicleRepository{
		db: db,
	}
}

func (r VehicleRepository) Create(v *domain.Vehicle) error {
	stmt, err := r.db.Prepare(`INSERT INTO vehicle 
	(type, license_plate, passenger_capacity, make, model, year, mileage) 
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(v.Type,
		v.LicensePlate,
		v.PassengerCapacity,
		v.Make,
		v.Model,
		v.Year,
		v.Mileage).Scan(&v.ID)
	if err != nil {
		return err
	}

	return nil
}
