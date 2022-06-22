package storage

import (
	"database/sql"

	"github.com/cornejodev/viator/config"
	"github.com/cornejodev/viator/internal/domain/vehicle"
	"github.com/cornejodev/viator/internal/storage/postgres"
	"github.com/rs/zerolog"
)

type Storage interface {
	ProvideRepository() (*Repository, error)
}

type storage struct {
	dbcfg config.Database
	lgr   zerolog.Logger
}

func New(dbcfg config.Database, lgr zerolog.Logger) *storage {
	return &storage{dbcfg, lgr}
}

func (s *storage) ProvideRepository() (*Repository, error) {
	var err error
	var db *sql.DB

	db, err = postgres.New(s.dbcfg)
	if err != nil {
		s.lgr.Error().Err(err).Msgf("postgres: %v", err)
		return nil, err
	}

	s.lgr.Info().Msg("Connected to Postgres!")

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
