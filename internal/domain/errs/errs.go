package errs

import (
	"errors"
	"fmt"
	"runtime"
)

// errs global var
var (
	_caller bool // sets stacktrace
)

// Kinds of errors.
//
// The values of the error kinds are common between both
// clients and servers. Do not reorder this list or remove
// any items since that will change their values.
// New items must be added only to the end.
const (
	// Unclassified error. This value is not printed in the error message.
	Other Kind = iota
	// Invalid operation for this type of item.
	Invalid
	// External I/O error such as network failure.
	IO
	// Item already exists.
	Exist
	// Item does not exist.
	NotExist
	// Information withheld.
	Private
	// Internal error or inconsistency.
	Internal
	// Link target does not exist.
	BrokenLink
	// Error from database.
	Database
	// Input validation error.
	Validation
	// Unanticipated error.
	Unanticipated
	// Invalid Request
	InvalidRequest
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

// Kind defines the kind of error this is, mostly for use by systems
type Kind uint8

// Op is the operation when error happens
type Op string

// Parameter represents the parameter related to the error.
type Parameter string

// Code is a human-readable, short representation of the error
type Code string

// Errors of xerrors
type Error struct {
	// Op represents the caller when error happens
	Op Op
	// Kind is the class of error, such as permission failure,
	// or "Other" if its class is unknown or irrelevant.
	Kind Kind
	// Param represents the parameter related to the error.
	Param Parameter
	// Code is a human-readable, short representation of the error
	Code Code
	// The underlying error that triggered this one, if any.
	Err error
}

// New errors
func E(v ...interface{}) error {
	var (
		e    = &Error{}
		file string
		line int
	)

	// only cal _caller when errs _caller is true
	if _caller {
		_, file, line, _ = runtime.Caller(1)
	}

	for _, arg := range v {
		switch val := arg.(type) {
		case Op:
			e.Op = val
		case Code:
			e.Code = val
			e.Err = fmt.Errorf("error executing %s: %s", e.Op, val)
		case Parameter:
			e.Param = val
		case string:
			if _caller {
				e.Err = fmt.Errorf("%s: %s [file=%s, line=%d]", e.Op, val, file, line)
				continue
			}

			e.Err = fmt.Errorf("error executing %s: %s", e.Op, val)

		case Kind:
			e.Kind = val

		case *Error:
			val.Op = e.Op
			// copy the errors
			e = val

			if _caller {
				e.Err = fmt.Errorf("error executing %s: [file=%s, line=%d] \n%w", e.Op, file, line, val.Err)
				continue
			}

			e.Err = fmt.Errorf("error executing %s: %w", e.Op, val.Err)

		case error:
			if _caller {
				e.Err = fmt.Errorf("%s %w [file=%s, line=%d]", e.Op, val, file, line)
				continue
			}

			e.Err = fmt.Errorf("error executing %s: %w", e.Op, val)

		default:
			continue
		}
	}
	return e
}

// Error return string of error
func (e *Error) Error() string {
	return e.Err.Error()
}

// Unwrap errors
func (e *Error) Unwrap() error {
	return e.Err
}

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

// Kind of errors
func KindOf(err error) Kind {
	var e *Error
	errors.As(err, &e)
	return e.Kind
}

// String  value of Op
func (op Op) String() string {
	return string(op)
}

// Is wrap the errors is
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As wrap the error as
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// SetCaller to print the stack-trace of the error
func SetCaller(c bool) {
	_caller = c
}
