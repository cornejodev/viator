package storage

import (
	"context"
	"database/sql"

	"github.com/cornejodev/viator/config"
	"github.com/cornejodev/viator/internal/domain/errs"
	"github.com/cornejodev/viator/internal/domain/vehicle"
	"github.com/cornejodev/viator/internal/storage/postgres"
)

type Storage interface {
	ProvideRepository() (*Repository, error)
}

type storage struct {
	dbcfg config.Database
}

func New(dbcfg config.Database) (*storage, error) {
	return &storage{dbcfg}, nil
}

func (s *storage) ProvideRepository() (*Repository, error) {
	var op errs.Op = "storage.ProvideRepository"
	var err error
	var db *sql.DB

	db, err = postgres.New(s.dbcfg)
	if err != nil {
		return nil, errs.E(op, err)
	}

	return &Repository{
		Vehicle: postgres.NewVehicleRepository(db),
	}, nil
}

type Repository struct {
	Vehicle VehicleRepository
}

type VehicleRepository interface {
	Create(ctx context.Context, v vehicle.Vehicle) error
	// ByID(id int) (vehicle.Vehicle, error)
	// All() ([]vehicle.Vehicle, error)
	// Update(v vehicle.Vehicle) error
	// Delete(id int) error
}
