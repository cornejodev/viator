package rest

const (
	msgOK   = "success"
	msgFail = "fail"
	msgErr  = "error"
)

type response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// newResponse return standard JSON response based off JSend specification - https://github.com/omniti-labs/jsend
// Usage example: resp := newResponse(msgOK, "resource has been updated", data).
func newResponse(msgType, msg string, data interface{}) response {
	var r response

	switch msgType {
	case msgOK:
		r = response{
			Status:  msgOK,
			Message: msg,
			Data:    data,
		}
	case msgFail:
		r = response{
			Status:  msgFail,
			Message: msg,
			Data:    data,
		}
	case msgErr:
		r = response{
			Status:  msgErr,
			Message: msg,
			Data:    data,
		}
	}

	return r
}
