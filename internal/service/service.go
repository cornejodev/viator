package service

import (
	"github.com/cornejodev/viator/internal/domain/errs"
	"github.com/cornejodev/viator/internal/storage"
)

type Service struct {
	Storage storage.Storage
	Depot   DepotService
}

func New(s storage.Storage) (*Service, error) {
	var op errs.Op = "service.New"

	r, err := s.ProvideRepository()
	if err != nil {
		return nil, errs.E(op, err)
	}

	return &Service{
		Depot: NewDepotService(r.Vehicle),
	}, nil
}
