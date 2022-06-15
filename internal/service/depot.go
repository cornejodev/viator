package service

import (
	"log"

	"github.com/cornejodev/viator/internal/domain/errs"
	"github.com/cornejodev/viator/internal/domain/vehicle"
	"github.com/cornejodev/viator/internal/storage"
)

type DepotService interface {
	Add(f AddVehicleRequest) error
	Find(id int) (VehicleResponse, error)
	List() ([]VehicleResponse, error)
	Update(f UpdateVehicleRequest) error
	Remove(id int) error
}

type depotService struct {
	repo storage.VehicleRepository
}

func NewDepotService(repo storage.VehicleRepository) DepotService {
	return &depotService{repo}
}

// AddVehicleRequest is the request struct for adding a vehicle
type AddVehicleRequest struct {
	Type              string `json:"type"`
	LicensePlate      string `json:"licensePlate"`
	PassengerCapacity int    `json:"passengerCapacity"`
	Model             string `json:"model"`
	Make              string `json:"make"`
	Year              int    `json:"year"`
	Mileage           int    `json:"mileage"`
}

// UpdateVehicleRequest is the request struct for updating a vehicle
type UpdateVehicleRequest struct {
	ID                int    `json:"id"`
	Type              string `json:"type"`
	LicensePlate      string `json:"licensePlate"`
	PassengerCapacity int    `json:"passengerCapacity"`
	Model             string `json:"model"`
	Make              string `json:"make"`
	Year              int    `json:"year"`
	Mileage           int    `json:"mileage"`
}

// VehicleResponse returns a response struct containing fields that are relevant to client
type VehicleResponse struct {
	ID                int    `json:"id"`
	Type              string `json:"type"`
	LicensePlate      string `json:"licensePlate"`
	PassengerCapacity int    `json:"passengerCapacity"`
	Model             string `json:"model"`
	Make              string `json:"make"`
	Year              int    `json:"year"`
	Mileage           int    `json:"mileage"`
}

// Add is used to add a vehicle
func (ds *depotService) Add(rb AddVehicleRequest) error {
	v := vehicle.Vehicle{
		Type:              rb.Type,
		LicensePlate:      rb.LicensePlate,
		PassengerCapacity: rb.PassengerCapacity,
		Make:              rb.Make,
		Model:             rb.Model,
		Year:              rb.Year,
		Mileage:           rb.Mileage,
	}

	if err := v.IsValid(); err != nil {
		return err
	}

	if err := ds.repo.Create(v); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// Find is used to find a vehicle by ID
func (ds *depotService) Find(id int) (VehicleResponse, error) {
	if id == 0 {
		return VehicleResponse{}, errs.ErrVehicleNotFound
	}

	v, err := ds.repo.ByID(id)
	if err != nil {
		return VehicleResponse{}, err
	}

	vr := VehicleResponse{
		ID:                v.ID,
		Type:              v.Type,
		LicensePlate:      v.LicensePlate,
		PassengerCapacity: v.PassengerCapacity,
		Model:             v.Model,
		Make:              v.Make,
		Year:              v.Year,
		Mileage:           v.Mileage,
	}

	return vr, nil
}

// List is used to list all the vehicles in depot
func (ds *depotService) List() ([]VehicleResponse, error) {
	vehicles, err := ds.repo.All()
	if err != nil {
		return nil, err
	}

	list := make([]VehicleResponse, 0, len(vehicles))

	assemble := func(v vehicle.Vehicle) VehicleResponse {
		return VehicleResponse{
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

// Update is used to update a vehicle
func (ds *depotService) Update(rb UpdateVehicleRequest) error {
	v := vehicle.Vehicle{
		Type:              rb.Type,
		LicensePlate:      rb.LicensePlate,
		PassengerCapacity: rb.PassengerCapacity,
		Make:              rb.Make,
		Model:             rb.Model,
		Year:              rb.Year,
		Mileage:           rb.Mileage,
	}

	if err := v.IsValid(); err != nil {
		return err
	}

	v.ID = rb.ID

	if err := ds.repo.Update(v); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// Remove is used to remove a vehicle
func (ds *depotService) Remove(id int) error {
	return ds.repo.Delete(id)
}
