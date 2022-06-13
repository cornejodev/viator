package service

import (
	"github.com/cornejodev/viator/internal/domain"
	"github.com/cornejodev/viator/internal/storage"
)

type DepotService interface {
	Add(f *domain.AddVehicleForm) error
	Find(id int) (*domain.VehicleCard, error)
	List() (domain.VehicleList, error)
}

type depotService struct {
	repo storage.VehicleRepository
}

func NewDepotService(repo storage.VehicleRepository) DepotService {
	return &depotService{repo}
}

func (ds *depotService) Add(f *domain.AddVehicleForm) error {
	if err := f.CheckEmptyFields(); err != nil {
		return err
	}
	v := &domain.Vehicle{
		Type:              f.Type,
		LicensePlate:      f.LicensePlate,
		PassengerCapacity: f.PassengerCapacity,
		Make:              f.Make,
		Model:             f.Model,
		Year:              f.Year,
		Mileage:           f.Mileage,
	}
	return ds.repo.Create(v)
}

func (ds *depotService) Find(id int) (*domain.VehicleCard, error) {
	if id == 0 {
		return &domain.VehicleCard{}, domain.ErrVehicleNotFound
	}

	v, err := ds.repo.ByID(id)
	if err != nil {
		return &domain.VehicleCard{}, err
	}

	vc := &domain.VehicleCard{
		ID:                v.ID,
		Type:              v.Type,
		LicensePlate:      v.LicensePlate,
		PassengerCapacity: v.PassengerCapacity,
		Model:             v.Model,
		Make:              v.Make,
		Year:              v.Year,
		Mileage:           v.Mileage,
	}

	return vc, nil
}

func (ds *depotService) List() (domain.VehicleList, error) {
	vehicles, _ := ds.repo.All()

	list := make(domain.VehicleList, 0, len(vehicles))

	assemble := func(v *domain.Vehicle) domain.VehicleCard {
		return domain.VehicleCard{
			ID:                v.ID,
			Type:              v.Type,
			LicensePlate:      v.LicensePlate,
			PassengerCapacity: v.PassengerCapacity,
			Model:             v.Model,
			Make:              v.Make,
			Year:              v.Year,
			Mileage:           v.Mileage,
		}
	}

	for _, v := range vehicles {
		list = append(list, assemble(v))
	}
	return list, nil

}
