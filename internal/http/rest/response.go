package rest

const (
	msgOK  = "ok"
	msgErr = "error"
)

type response struct {
	Success       bool `json:"success"`
	*messageOK    `json:"message,omitempty"`
	*messageError `json:"error_message,omitempty"`
	Data          interface{} `json:"data,omitempty"`
}

type messageOK struct {
	Content string `json:"content"`
}

type messageError struct {
	Content string `json:"content"`
}

// newResponse returns standard JSON response
// Usage example: resp := newResponse(msgOK,"resource has been updated", data).
func newResponse(message, content string, data interface{}) response {
	var r response

	switch message {
	case msgOK:
		r = response{
			Success: true,
			messageOK: &messageOK{
				Content: content,
			},
			messageError: nil,
			Data:         data,
		}
	case msgErr:
		r = response{
			messageOK: nil,
			Success:   false,
			messageError: &messageError{
				Content: content,
			},
			Data: data,
		}
	}

	return r
}
