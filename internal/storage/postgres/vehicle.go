package postgres

import (
	"database/sql"
	"log"
	"time"

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

func (r *VehicleRepository) Create(v *domain.Vehicle) error {
	stmt, err := r.db.Prepare(
		`INSERT INTO vehicle (type, license_plate, passenger_capacity, make, model, year, mileage, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
	RETURNING id`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	v.CreatedAt = time.Now()

	err = stmt.QueryRow(
		v.Type,
		v.LicensePlate,
		v.PassengerCapacity,
		v.Make,
		v.Model,
		v.Year,
		v.Mileage,
		v.CreatedAt,
	).Scan(&v.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *VehicleRepository) ByID(id int) (*domain.Vehicle, error) {
	v := &domain.Vehicle{}

	stmt, err := r.db.Prepare(
		`SELECT id, type, license_plate, passenger_capacity, make, model, year, mileage, created_at
	FROM vehicle 
	WHERE id = $1`)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(
		&v.ID,
		&v.Type,
		&v.LicensePlate,
		&v.PassengerCapacity,
		&v.Make,
		&v.Model,
		&v.Year,
		&v.Mileage,
		&v.CreatedAt,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return v, nil
}
