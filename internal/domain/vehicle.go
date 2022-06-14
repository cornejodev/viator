package domain

import (
	"errors"
	"time"
)

var ErrVehicleCantBeEmpty = errors.New("the vehicle fields can't be empty")
var ErrVehicleNotFound = errors.New("vehicle not found")

type Vehicle struct {
	ID                int    // internal vehicle ID for Viator
	Type              string // car, truck, SUV, van, motorcycle
	LicensePlate      string
	PassengerCapacity int
	Make              string
	Model             string
	Year              int
	Mileage           int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type AddVehicleForm struct {
	Type              string `json:"type"`
	LicensePlate      string `json:"licensePlate"`
	PassengerCapacity int    `json:"passengerCapacity"`
	Model             string `json:"model"`
	Make              string `json:"make"`
	Year              int    `json:"year"`
	Mileage           int    `json:"mileage"`
}

type UpdateVehicleForm struct {
	ID                int    `json:"id"`
	Type              string `json:"type"`
	LicensePlate      string `json:"licensePlate"`
	PassengerCapacity int    `json:"passengerCapacity"`
	Model             string `json:"model"`
	Make              string `json:"make"`
	Year              int    `json:"year"`
	Mileage           int    `json:"mileage"`
}

// returns only fields that are relevant to client
type VehicleCard struct {
	ID                int    `json:"id"`
	Type              string `json:"type"`
	LicensePlate      string `json:"licensePlate"`
	PassengerCapacity int    `json:"passengerCapacity"`
	Model             string `json:"model"`
	Make              string `json:"make"`
	Year              int    `json:"year"`
	Mileage           int    `json:"mileage"`
}

func (f *AddVehicleForm) CheckEmptyFields() error {
	if f.Type == "" ||
		f.LicensePlate == "" ||
		f.PassengerCapacity == 0 ||
		f.Model == "" ||
		f.Make == "" ||
		f.Year == 0 ||
		f.Mileage == 0 {

		return ErrVehicleCantBeEmpty
	}
	return nil
}

func (f *UpdateVehicleForm) CheckEmptyFields() error {
	if f.ID == 0 ||
		f.Type == "" ||
		f.LicensePlate == "" ||
		f.PassengerCapacity == 0 ||
		f.Model == "" ||
		f.Make == "" ||
		f.Year == 0 ||
		f.Mileage == 0 {

		return ErrVehicleCantBeEmpty
	}
	return nil
}

type VehicleList []VehicleCard

/*
Types of vehicles
	- Car
	- Truck
	- SUV
	- Van
	- Motorcycle
*/
