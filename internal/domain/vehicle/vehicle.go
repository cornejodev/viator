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
	const op errs.Op = "vehicle.IsValid"

	switch {
	case v.Type == "":
		return errs.E(op, errs.Parameter("type"), errs.Code(errs.MissingField("type")), errs.Validation)
	case v.LicensePlate == "":
		return errs.E(op, errs.Parameter("licensePlate"), errs.Code(errs.MissingField("licensePlate")), errs.Validation)
	case v.PassengerCapacity <= 0:
		return errs.E(op, errs.Parameter("passengerCapacity"), errs.Code("passengerCapacity must be greater than zero"), errs.Validation)
	case v.Model == "":
		return errs.E(op, errs.Parameter("model"), errs.Code(errs.MissingField("model")), errs.Validation)
	case v.Make == "":
		return errs.E(op, errs.Parameter("make"), errs.Code(errs.MissingField("make")), errs.Validation)
	case v.Year <= 0:
		return errs.E(op, errs.Parameter("year"), errs.Code("year must be greater than zero"), errs.Validation)
	case v.Mileage == 0:
		return errs.E(op, errs.Parameter("mileage"), errs.Code("mileage must be greater than zero"), errs.Validation)
	}

	return nil
}
