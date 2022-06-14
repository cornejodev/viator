package storage

import (
	"database/sql"
	"fmt"

	"github.com/cornejodev/viator/config"
	"github.com/cornejodev/viator/internal/domain/vehicle"
	"github.com/cornejodev/viator/internal/storage/postgres"
)

type Storage interface {
	ProvideRepository() (*Repository, error)
}

type storage struct {
	dbcfg config.Database
}

func New(dbcfg config.Database) *storage {
	return &storage{dbcfg}
}

func (s *storage) ProvideRepository() (*Repository, error) {
	var err error
	var db *sql.DB

	db, err = postgres.New(s.dbcfg)
	if err != nil {
		return nil, fmt.Errorf("postgres: %v", err)
	}

	return &Repository{
		Vehicle: postgres.NewVehicleRepository(db),
	}, nil
}

type Repository struct {
	Vehicle VehicleRepository
}

type VehicleRepository interface {
	Create(v vehicle.Vehicle) error
	ByID(id int) (vehicle.Vehicle, error)
	All() ([]vehicle.Vehicle, error)
	Update(v vehicle.Vehicle) error
	Delete(id int) error
}
