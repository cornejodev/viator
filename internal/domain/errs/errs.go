package errs

import (
	"errors"
	"fmt"
	"runtime"
)

var ErrVehicleCantBeEmpty = errors.New("the vehicle fields can't be empty")
var ErrVehicleNotFound = errors.New("vehicle not found")

// xerrors global var
var (
	_caller bool
)

/*======================================================================================================*/

type Kind uint8

// Kinds of errors.
//
// The values of the error kinds are common between both
// clients and servers. Do not reorder this list or remove
// any items since that will change their values.
// New items must be added only to the end.
const (
	Other          Kind = iota // Unclassified error. This value is not printed in the error message.
	Invalid                    // Invalid operation for this type of item.
	IO                         // External I/O error such as network failure.
	Exist                      // Item already exists.
	NotExist                   // Item does not exist.
	Private                    // Information withheld.
	Internal                   // Internal error or inconsistency.
	BrokenLink                 // Link target does not exist.
	Database                   // Error from database.
	Validation                 // Input validation error.
	Unanticipated              // Unanticipated error.
	InvalidRequest             // Invalid Request
	// Unauthenticated is used when a request lacks valid authentication credentials.
	//
	// For Unauthenticated errors, the response body will be empty.
	// The error is logged and http.StatusUnauthorized (401) is sent.
	Unauthenticated // Unauthenticated Request
	// Unauthorized is used when a user is authenticated, but is not authorized
	// to access the resource.
	//
	// For Unauthorized errors, the response body should be empty.
	// The error is logged and http.StatusForbidden (403) is sent.
	Unauthorized
)

func (k Kind) String() string {
	switch k {
	case Other:
		return "other_error"
	case Invalid:
		return "invalid_operation"
	case IO:
		return "I/O_error"
	case Exist:
		return "item_already_exists"
	case NotExist:
		return "item_does_not_exist"
	case BrokenLink:
		return "link_target_does_not_exist"
	case Private:
		return "information_withheld"
	case Internal:
		return "internal_error"
	case Database:
		return "database_error"
	case Validation:
		return "input_validation_error"
	case Unanticipated:
		return "unanticipated_error"
	case InvalidRequest:
		return "invalid_request_error"
	case Unauthenticated:
		return "unauthenticated_request"
	case Unauthorized:
		return "unauthorized_request"
	}
	return "unknown_error_kind"
}

/*======================================================================================================*/

// // Kind of errors
// type Kind int16

// // kind of errors
// const (
// 	KindOK Kind = iota
// 	KindNotFound
// 	KindBadRequest
// 	KindUnauthorized
// 	KindInternalError
// )

// Op is the operation when error happens
type Op string

// String  value of Op
func (op Op) String() string {
	return string(op)
}

// Fields of errors
type Fields map[string]interface{}

// Errors of xerrors
type Errors struct {
	Err      error
	InnerErr error
	kind     Kind
	op       Op
}

// New errors
func E(v ...interface{}) error {
	var (
		xerr = &Errors{}
		file string
		line int
	)

	// only cal _caller when xerrors _caller is true
	if _caller {
		_, file, line, _ = runtime.Caller(1)
	}

	for _, arg := range v {
		switch val := arg.(type) {
		case Op:
			xerr.op = val

		case string:
			if _caller {
				xerr.Err = fmt.Errorf("%s: %s: [file=%s, line=%d]", val, xerr.op, file, line)
				continue
			}
			// change 1
			// 			e.Err = errors.New(arg)
			// xerr.Err = errors.New(val)

			xerr.Err = fmt.Errorf("%s: %s", xerr.op, val)

		case Kind:
			xerr.kind = val

		case *Errors:
			val.op = xerr.op
			// copy the errors
			xerr = val

			if _caller {
				xerr.Err = fmt.Errorf("error executing %s: [file=%s, line=%d] \n%w", xerr.op, file, line, val.Err)
				continue
			}
			// xerr.Err = fmt.Errorf("error executing %s: %w:", xerr.op, val.Err) original
			xerr.Err = fmt.Errorf("error executing %s: %w", xerr.op, val.Err)

		case error:
			if _caller {
				xerr.Err = fmt.Errorf("%w: %s: [file=%s, line=%d]", val, xerr.op, file, line)
				continue
			}
			// xerr.Err = fmt.Errorf("%w %s", val, xerr.op) -- original
			// xerr.Err = fmt.Errorf("%s", val)
			xerr.Err = fmt.Errorf("%s %w", xerr.op, val)

		default:
			continue
		}
	}
	return xerr
}

// Error return string of error
func (e *Errors) Error() string {
	return e.Err.Error()
}

// Unwrap errors
func (e *Errors) Unwrap() error {
	return e.Err
}

// Kind of errors
func (e *Errors) Kind() Kind {
	return e.kind
}

// Is wrap the errors is
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As wrap the error as
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// // Unwrap error
// func Unwrap(err error) error {
// 	return errors.Unwrap(err)
// }

// XUnwrap return errors with xerror package type
func Unwrap(err error) *Errors {
	xerr, ok := err.(*Errors)
	if ok {
		return xerr
	}

	return nil
}

// SetCaller to print the stack-trace of the error
func SetCaller(c bool) {
	_caller = c
}
