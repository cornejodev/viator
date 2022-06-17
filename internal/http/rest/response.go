package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JSONResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	j := JSONResponse{
		Status: "success",
		Data:   data,
	}

	// Marshal JSONResponse struct to JSON for the response body
	r, _ := json.Marshal(j)
	resp := string(r)

	// Write Content-Type headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	// Write HTTP Statuscode
	w.WriteHeader(statusCode)

	// Write response body (json)
	fmt.Fprintln(w, resp)
}
