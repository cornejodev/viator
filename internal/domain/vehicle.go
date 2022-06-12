package domain

import "errors"

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

func (v Vehicle) CheckEmptyFields() error {
	if v.Type == "" ||
		v.LicensePlate == "" ||
		v.PassengerCapacity == 0 ||
		v.Model == "" ||
		v.Make == "" ||
		v.Year == 0 ||
		v.Mileage == 0 {

		return ErrVehicleCantBeEmpty
	}
	return nil
}

/*
Types of vehicles
	- Car
	- Truck
	- SUV
	- Van
	- Motorcycle
*/
