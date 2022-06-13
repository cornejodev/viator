package service

import (
	"github.com/cornejodev/viator/internal/domain"
	"github.com/cornejodev/viator/internal/storage"
)

type DepotService interface {
	Add(v *domain.Vehicle) error
	Find(id int) (*domain.Vehicle, error)
}

type depotService struct {
	repo storage.VehicleRepository
}

func NewDepotService(repo storage.VehicleRepository) DepotService {
	return &depotService{repo}
}

func (ds *depotService) Add(v *domain.Vehicle) error {
	if err := v.CheckEmptyFields(); err != nil {
		return err
	}
	return ds.repo.Create(v)
}

func (ds *depotService) Find(id int) (*domain.Vehicle, error) {
	if id == 0 {
		return &domain.Vehicle{}, domain.ErrVehicleNotFound
	}
	return ds.repo.ByID(id)
}
