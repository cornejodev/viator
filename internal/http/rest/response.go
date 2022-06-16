package rest

import (
	"encoding/json"
	"net/http"

	"github.com/cornejodev/viator/internal/domain/errs"
)

const (
	msgOK   = "success"
	msgFail = "fail"
	msgErr  = "error"
)

type response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Code    int
}

// newResponse return standard JSON response based off JSend specification - https://github.com/omniti-labs/jsend
// Usage example: resp := newResponse(msgOK, "resource has been updated", data).
func newResponse(msgType, msg string, err error, data interface{}) response {
	var r response

	var code int
	var e *errs.Errors

	if err != nil {
		e = errs.Unwrap(err)
		code = httpErrorStatusCode(e.Kind())
	} else {
		code = 200

	}

	switch msgType {
	case msgOK:
		r = response{
			Status:  msgOK,
			Message: msg,
			Data:    data,
			Code:    code,
		}
	case msgFail:
		r = response{
			Status:  msgFail,
			Message: e.Kind().String(),
			Data:    data,
			Code:    code,
		}
	case msgErr:
		r = response{
			Status:  msgErr,
			Message: e.Kind().String(),
			Data:    data,
			Code:    code,
		}
	}

	return r
}

// httpErrorStatusCode maps an error Kind to an HTTP Status Code
func httpErrorStatusCode(k errs.Kind) int {
	switch k {
	case errs.Invalid, errs.Exist, errs.NotExist, errs.Private, errs.BrokenLink, errs.Validation, errs.InvalidRequest:
		return http.StatusBadRequest
	// the zero value of Kind is Other, so if no Kind is present
	// in the error, Other is used. Errors should always have a
	// Kind set, otherwise, a 500 will be returned and no
	// error message will be sent to the caller
	case errs.Other, errs.IO, errs.Internal, errs.Database, errs.Unanticipated:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func (r *response) String() string {
	res, _ := json.Marshal(r)

	return string(res)
}
