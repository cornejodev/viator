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
		`INSERT INTO vehicle (type, license_plate, passenger_capacity, make, model, year, mileage, creation_date) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
	RETURNING id`)
	if err != nil {
		log.Println("error while preparing statement: ", err)
		return err
	}
	defer stmt.Close()

	v.CreationDate = time.Now()

	err = stmt.QueryRow(
		v.Type,
		v.LicensePlate,
		v.PassengerCapacity,
		v.Make,
		v.Model,
		v.Year,
		v.Mileage,
		v.CreationDate,
	).Scan(&v.ID)
	if err != nil {
		log.Println("error while trying to create vehicle: ", err)
		return err
	}

	log.Println("New record added. Record ID is:", v.ID)

	return nil
}

func (r *VehicleRepository) ByID(id int) (*domain.Vehicle, error) {
	v := &domain.Vehicle{}

	stmt, err := r.db.Prepare(
		`SELECT id, type, license_plate, passenger_capacity, make, model, year, mileage, creation_date
	FROM vehicle 
	WHERE id = $1`)
	if err != nil {
		log.Println("error while preparing statement: ", err)
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
		&v.CreationDate,
	)
	if err != nil {
		log.Println("error while trying to fetch vehicle: ", err)
		return nil, err
	}
	return v, nil
}

func (r *VehicleRepository) All() ([]*domain.Vehicle, error) {
	stmt, err := r.db.Prepare("SELECT * FROM vehicle")
	if err != nil {
		log.Println("error while preparing statement: ", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println("error while preparing rows query: ", err)
		return nil, err
	}
	defer rows.Close()

	vehicles := make([]*domain.Vehicle, 0)
	for rows.Next() {
		v := &domain.Vehicle{}

		err := rows.Scan(
			&v.ID,
			&v.Type,
			&v.LicensePlate,
			&v.PassengerCapacity,
			&v.Make,
			&v.Model,
			&v.Year,
			&v.Mileage,
			&v.CreationDate,
		)
		if err != nil {
			log.Println("error while trying to fetch vehicles: ", err)
			return vehicles, err
		}

		vehicles = append(vehicles, v)
	}
	if err := rows.Err(); err != nil {
		log.Println("error while iterating through rows: ", err)
		return nil, err
	}

	return vehicles, nil
}
