package vehicle

import (
	"errors"
	"time"
)

type Vehicle struct {
	ID                int    // internal vehicle ID for Viator
	Type              string // Car, Truck, SUV, Van, Motorcycle
	LicensePlate      string
	PassengerCapacity int
	Make              string
	Model             string
	Year              int
	Mileage           int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

// IsValid performs validation of struct
func (v *Vehicle) IsValid() error {
	switch {
	case v.Type == "":
		return errors.New("Missing field: type")
	case v.LicensePlate == "":
		return errors.New("Missing field: licensePlate")
	case v.PassengerCapacity <= 0:
		return errors.New("passengerCapacity must be greater than zero")
	case v.Model == "":
		return errors.New("Missing field: model")
	case v.Make == "":
		return errors.New("Missing field: make")
	case v.Year <= 0:
		return errors.New("year must be greater than zero")
	case v.Mileage == 0:
		return errors.New("mileage must be greater than zero")
	}

	return nil
}
