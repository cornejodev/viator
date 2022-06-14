package errs

import "errors"

var ErrVehicleCantBeEmpty = errors.New("the vehicle fields can't be empty")
var ErrVehicleNotFound = errors.New("vehicle not found")
