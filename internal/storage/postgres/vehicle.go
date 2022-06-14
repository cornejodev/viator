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
		`INSERT INTO vehicle (type, license_plate, passenger_capacity, make, model, year, mileage, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
	RETURNING id`)
	if err != nil {
		log.Println("error while preparing statement: ", err)
		return err
	}
	defer stmt.Close()

	v.CreatedAt = time.Now()
	v.UpdatedAt = time.Now()
	// cant use Exec() and then LAstInsertId with lib/pq driver because postgres
	// doesn't automatically return the last insert id. Therefore we use QueryRow instead
	err = stmt.QueryRow(
		v.Type,
		v.LicensePlate,
		v.PassengerCapacity,
		v.Make,
		v.Model,
		v.Year,
		v.Mileage,
		v.CreatedAt,
		v.UpdatedAt,
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
		`SELECT id, type, license_plate, passenger_capacity, make, model, year, mileage, created_at, updated_at
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
		&v.CreatedAt,
		&v.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("error while trying to fetch vehicle: ", err)
			return nil, err
		} else {
			log.Println("error while trying to fetch vehicle: ", err)
			return nil, err

		}
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
			&v.CreatedAt,
			&v.UpdatedAt,
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

func (r *VehicleRepository) Update(v domain.Vehicle) error {
	stmt, err := r.db.Prepare(
		`UPDATE vehicle 
	SET 
		type = $1, 
		license_plate = $2, 
		passenger_capacity = $3, 
		make = $4, 
		model = $5, 
		year = $6, 
		mileage = $7,
		updated_at = $8 
	WHERE 
		id = $9`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	v.UpdatedAt = time.Now()

	result, err := stmt.Exec(
		v.Type,
		v.LicensePlate,
		v.PassengerCapacity,
		v.Make,
		v.Model,
		v.Year,
		v.Mileage,
		v.UpdatedAt,
		v.ID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return domain.ErrVehicleNotFound
	}

	return nil
}
