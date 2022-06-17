package vehicle

import (
	"time"

	"github.com/cornejodev/viator/internal/domain/errs"
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
		return errs.E(errs.Parameter("type"), errs.Code("Missing field: type"), errs.Validation)
	case v.LicensePlate == "":
		return errs.E(errs.Parameter("licensePlate"), errs.Code("Missing field: licensePlate"), errs.Validation)
	case v.PassengerCapacity <= 0:
		return errs.E(errs.Parameter("passengerCapacity"), errs.Code("passengerCapacity must be greater than zero"), errs.Validation)
	case v.Model == "":
		return errs.E(errs.Parameter("model"), errs.Code("Missing field: model"), errs.Validation)
	case v.Make == "":
		return errs.E(errs.Parameter("make"), errs.Code("Missing field: make"), errs.Validation)
	case v.Year <= 0:
		return errs.E(errs.Parameter("year"), errs.Code("year must be greater than zero"), errs.Validation)
	case v.Mileage == 0:
		return errs.E(errs.Parameter("mileage"), errs.Code("mileage must be greater than zero"), errs.Validation)
	}

	return nil
}
