package service

import (
	"github.com/cornejodev/viator/internal/storage"
	"github.com/rs/zerolog"
)

type Service struct {
	Storage storage.Storage
	Depot   DepotService
}

func New(s storage.Storage, lgr zerolog.Logger) (*Service, error) {
	r, err := s.ProvideRepository()
	if err != nil {
		lgr.Error().Err(err).Msgf("error from storage: %w", err)
		return nil, err
	}

	return &Service{
		Depot: NewDepotService(r.Vehicle),
	}, nil
}
