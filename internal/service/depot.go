package service

import (
	"log"

	"github.com/cornejodev/viator/internal/domain"
	"github.com/cornejodev/viator/internal/storage"
)

type DepotService interface {
	Add(f *domain.AddVehicleForm) error
	Find(id int) (*domain.VehicleCard, error)
	List() (domain.VehicleList, error)
	Update(f domain.UpdateVehicleForm) error
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
		return nil, domain.ErrVehicleNotFound
	}

	v, err := ds.repo.ByID(id)
	if err != nil {
		return nil, err
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
	vehicles, err := ds.repo.All()
	if err != nil {
		return nil, err
	}

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

func (ds *depotService) Update(f domain.UpdateVehicleForm) error {
	if err := f.CheckEmptyFields(); err != nil {
		log.Println("Error while trying to update vehicle:", err)
		return err
	}

	err := ds.repo.Update(domain.Vehicle{
		ID:                f.ID,
		Type:              f.Type,
		LicensePlate:      f.LicensePlate,
		PassengerCapacity: f.PassengerCapacity,
		Make:              f.Make,
		Model:             f.Model,
		Year:              f.Year,
		Mileage:           f.Mileage,
	})
	if err != nil {
		log.Println(err)
		return err
	}

	// vc := &domain.VehicleCard{
	// 	ID:                v.ID,
	// 	Type:              v.Type,
	// 	LicensePlate:      v.LicensePlate,
	// 	PassengerCapacity: v.PassengerCapacity,
	// 	Model:             v.Model,
	// 	Make:              v.Make,
	// 	Year:              v.Year,
	// 	Mileage:           v.Mileage,
	// }

	return nil
}
