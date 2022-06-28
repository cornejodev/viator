package postgres

import (
	"context"
	"database/sql"
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

// var ctx = context.Background()

// Create is used to create a vehicle in the database
func (r *VehicleRepository) Create(ctx context.Context, v vehicle.Vehicle) error {
	const op errs.Op = "VehicleRepository.Create"

	// Get a Tx for making transaction requests.
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return errs.E(op, err)
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	stmt, err := r.db.PrepareContext(ctx, `
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
	err = stmt.QueryRowContext(ctx,
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
		if err.Error() == "pq: canceling statement due to user request" {
			return errs.E(
				op,
				errs.Code(ctx.Err().Error()+" "+err.Error()),
			)
		}
		return errs.E(op, err)
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return errs.E(op, err)
	}

	return nil
}

// ByID is used to find a vehicle in database via its ID. It returns a Vehicle struct to the caller
func (r *VehicleRepository) ByID(ctx context.Context, id int) (vehicle.Vehicle, error) {
	const op errs.Op = "VehicleRepository.ByID"
	var v vehicle.Vehicle

	// Get a Tx for making transaction requests.
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return vehicle.Vehicle{}, errs.E(op, err)
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	stmt, err := r.db.PrepareContext(ctx, `
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
		return vehicle.Vehicle{}, errs.E(op, err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, id).Scan(
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
		if err.Error() == "pq: canceling statement due to user request" {
			return vehicle.Vehicle{}, errs.E(
				op,
				errs.Code(ctx.Err().Error()+" "+err.Error()),
			)
		}
		if err == sql.ErrNoRows {
			return vehicle.Vehicle{}, errs.E(op, err, errs.NotExist)
		} else {
			return vehicle.Vehicle{}, errs.E(op, err)

		}
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return vehicle.Vehicle{}, errs.E(op, err)
	}

	return v, nil
}

// // All returns all the vehicles stored in database. It returns  a slice of Vehicle structs
func (r *VehicleRepository) All(ctx context.Context) ([]vehicle.Vehicle, error) {
	const op errs.Op = "VehicleRepository.All"

	// Get a Tx for making transaction requests.
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return []vehicle.Vehicle{}, errs.E(op, err)
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	stmt, err := r.db.PrepareContext(ctx, "SELECT * FROM vehicle")
	if err != nil {
		return nil, errs.E(op, err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		if err.Error() == "pq: canceling statement due to user request" {
			return nil, errs.E(
				op,
				errs.Code(ctx.Err().Error()+" "+err.Error()),
				errs.Timeout,
			)
		}
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

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return []vehicle.Vehicle{}, errs.E(op, err)
	}

	return vehicles, nil
}

// // Update is used to update a vehicle in the database
func (r *VehicleRepository) Update(ctx context.Context, v vehicle.Vehicle) error {
	const op errs.Op = "VehicleRepository.Update"

	// Get a Tx for making transaction requests.
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return errs.E(op, err)
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	stmt, err := r.db.PrepareContext(ctx, `
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

	result, err := stmt.ExecContext(ctx,
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
		if err.Error() == "pq: canceling statement due to user request" {
			return errs.E(
				op,
				errs.Code(ctx.Err().Error()+" "+err.Error()),
			)
		}
		return errs.E(op, err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return errs.E(op, err)
	}

	if rows == 0 {
		return errs.E(op, errs.Code("item doesn't exist in database"), errs.NotExist)
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return errs.E(op, err)
	}

	return nil
}

// Delete is used to delete a vehicle in the database
func (r *VehicleRepository) Delete(ctx context.Context, id int) error {
	const op errs.Op = "VehicleRepository.Delete"

	// Get a Tx for making transaction requests.
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return errs.E(op, err)
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	stmt, err := r.db.PrepareContext(ctx, "DELETE FROM vehicle WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, id)
	if err != nil {
		if err.Error() == "pq: canceling statement due to user request" {
			return errs.E(
				op,
				errs.Code(ctx.Err().Error()+" "+err.Error()),
			)
		}
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errs.E(op, errs.Code("item doesn't exist in database"), errs.NotExist)
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return errs.E(op, err)
	}

	return nil
}
