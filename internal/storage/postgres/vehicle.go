package postgres

import (
	"database/sql"
	"log"
	"time"

	"github.com/cornejodev/viator/internal/domain/errs"
	"github.com/cornejodev/viator/internal/domain/vehicle"
)

type VehicleRepository struct {
	db *sql.DB
}

func NewVehicleRepository(db *sql.DB) *VehicleRepository {
	return &VehicleRepository{
		db: db,
	}
}

// Create is used to create a vehicle in the database
func (r *VehicleRepository) Create(v vehicle.Vehicle) error {
	const op errs.Op = "VehicleRepository.Create"

	stmt, err := r.db.Prepare(`
	INSERT INTO 
		vehicle(
			type, 
			license_plate, 
			passenger_capacity, 
			make, 
			model, 
			year, 
			mileage, 
			created_at,
			updated_at
		) 
	VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9) 
	RETURNING id
	`)
	if err != nil {
		return errs.E(op, err)
	}
	defer stmt.Close()

	v.CreatedAt = time.Now()
	v.UpdatedAt = time.Now()

	// cant use Exec() and then LastInsertId with lib/pq driver because postgres
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
		return errs.E(op, err)
	}

	log.Printf("New vehicle added. Vehicle ID is: %d", v.ID)

	return nil
}

// ByID is used to find a vehicle in database via its ID. It returns a Vehicle struct to the caller
func (r *VehicleRepository) ByID(id int) (vehicle.Vehicle, error) {
	const op errs.Op = "VehicleRepository.ByID"
	var v vehicle.Vehicle

	stmt, err := r.db.Prepare(`
	SELECT 
		id, 
		type, 
		license_plate, 
		passenger_capacity, 
		make, 
		model, 
		year, 
		mileage, 
		created_at, 
		updated_at
	FROM 
		vehicle 
	WHERE 
		id = $1
	`)
	if err != nil {
		// log.Println("error while preparing statement: ", err)
		return vehicle.Vehicle{}, errs.E(op, err)
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
			return vehicle.Vehicle{}, errs.E(op, err, errs.NotExist)
		} else {
			return vehicle.Vehicle{}, errs.E(op, err)

		}
	}
	return v, nil
}

// All returns all the vehicles stored in database. It returns  a slice of Vehicle structs
func (r *VehicleRepository) All() ([]vehicle.Vehicle, error) {
	const op errs.Op = "VehicleRepository.All"

	stmt, err := r.db.Prepare("SELECT * FROM vehicle")
	if err != nil {
		return nil, errs.E(op, err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, errs.E(op, err)
	}
	defer rows.Close()

	vehicles := make([]vehicle.Vehicle, 0)
	for rows.Next() {
		var v vehicle.Vehicle

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
			return vehicles, errs.E(op, err)
		}

		vehicles = append(vehicles, v)
	}
	if err := rows.Err(); err != nil {
		return nil, errs.E(op, err)
	}

	return vehicles, nil
}

// Update is used to update a vehicle in the database
func (r *VehicleRepository) Update(v vehicle.Vehicle) error {
	const op errs.Op = "VehicleRepository.Update"

	stmt, err := r.db.Prepare(`
	UPDATE 
		vehicle 
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
		id = $9
	`)
	if err != nil {
		return errs.E(op, err)
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
		return errs.E(op, err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return errs.E(op, err)
	}

	if rows == 0 {
		return errs.E(op, errs.Code("item doesn't exist in database"), errs.NotExist)
	}

	log.Printf("Vehicle with ID: %d has been modified", v.ID)

	return nil
}

// Delete is used to delete a vehicle in the database
func (r *VehicleRepository) Delete(id int) error {
	const op errs.Op = "VehicleRepository.Delete"

	stmt, err := r.db.Prepare("DELETE FROM vehicle WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errs.E(op, errs.Code("item doesn't exist in database"), errs.NotExist)
	}

	log.Printf("Vehicle with ID: %d removed from DB", id)

	return nil
}
