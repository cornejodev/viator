package rest

import (
	"net/http"

	"github.com/cornejodev/viator/internal/service"
	"github.com/gorilla/mux"
)

func Handler(s service.Service) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/vehicle", addVehicle(s)).Methods("POST")
	r.HandleFunc("/vehicle/{id}", getVehicle(s)).Methods("GET")

	return r
}
