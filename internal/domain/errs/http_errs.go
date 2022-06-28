package errs

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/rs/zerolog"
)

// ErrResponse is used as the Response Body
type ErrResponse struct {
	Status string       `json:"status"`
	Error  ServiceError `json:"error"`
}

// ServiceError has fields for Service errors. All fields with no data will
// be omitted
type ServiceError struct {
	Kind    string `json:"kind,omitempty"`
	Param   string `json:"param,omitempty"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// HTTPErrorResponse takes a writer and an error, performs a
// type switch to determine if the type is an Error (which meets
// the Error interface as defined in this package), then sends the
// Error as a response to the client. If the type does not meet the
// Error interface as defined in this package, then a proper error
// is still formed and sent to the client, however, the Kind and
// Code will be Unanticipated.
func HTTPErrorResponse(w http.ResponseWriter, r *http.Request, lgr zerolog.Logger, err error) {
	if err == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var e *Error
	errors.As(err, &e)
	typicalErrorResponse(w, r, lgr, e)
}

// typicalErrorResponse replies to the request with the specified error
// message and HTTP code. It does not otherwise end the request; the
// caller should ensure no further writes are done to w.
//
// Taken from standard library and modified.
// https://golang.org/pkg/net/http/#Error
func typicalErrorResponse(w http.ResponseWriter, r *http.Request, lgr zerolog.Logger, e *Error) {
	ctx := r.Context()
	val := ctx.Value("request_id")
	rID := val.(string) // maybe check is needed here maybe not...

	httpStatusCode := httpErrorStatusCode(e.Kind)

	lgr.Error().
		Str("request_id", rID).
		Str("kind", e.Kind.String()).
		Err(e).
		Str("remote_ip", r.RemoteAddr).
		Str("user_agent", r.UserAgent()).
		Str("method", r.Method).
		Str("url", r.URL.String()).
		Int("status", httpStatusCode).
		Msg("Error response sent")

	// get ErrResponse
	er := newErrResponse(e)

	// Marshal errResponse struct to JSON for the response body
	errJSON, _ := json.Marshal(er)
	ej := string(errJSON)

	// Write Content-Type headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	// Write HTTP Statuscode
	w.WriteHeader(httpStatusCode)

	// Write response body (json)
	fmt.Fprintln(w, ej)
}

func newErrResponse(err *Error) ErrResponse {
	const msg string = "internal server error - please contact support"

	switch err.Kind {
	case Internal, Database:
		return ErrResponse{
			Status: "error",
			Error: ServiceError{
				Kind:    Internal.String(),
				Message: msg,
			},
		}
	default:
		return ErrResponse{
			Status: "error",
			Error: ServiceError{
				Kind:  err.Kind.String(),
				Param: string(err.Param),
				Code:  string(err.Code),
				// Message: err.Error(),
			},
		}
	}
}

// httpErrorStatusCode maps an error Kind to an HTTP Status Code
func httpErrorStatusCode(k Kind) int {
	switch k {
	case Invalid, Exist, NotExist, Private, BrokenLink, Validation, InvalidRequest:
		return http.StatusBadRequest
	// the zero value of Kind is Other, so if no Kind is present
	// in the error, Other is used. Errors should always have a
	// Kind set, otherwise, a 500 will be returned and no
	// error message will be sent to the caller
	case Other, IO, Internal, Database, Unanticipated:
		return http.StatusInternalServerError
	case Timeout:
		return http.StatusServiceUnavailable
	default:
		return http.StatusInternalServerError
	}
}

func ipFromHostPort(hp string) string {
	h, _, err := net.SplitHostPort(hp)
	if err != nil {
		return ""
	}
	if len(h) > 0 && h[0] == '[' {
		return h[1 : len(h)-1]
	}
	return h
}
