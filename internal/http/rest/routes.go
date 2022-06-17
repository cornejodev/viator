package rest

import (
	"net/http"

	"github.com/cornejodev/viator/internal/service"
	"github.com/gorilla/mux"
)

func Handler(s service.Service) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/vehicles", addVehicle(s)).Methods("POST")
	r.HandleFunc("/vehicles", listVehicles(s)).Methods("GET")
	r.HandleFunc("/vehicles/{id}", getVehicle(s)).Methods("GET")
	r.HandleFunc("/vehicles/{id}", updateVehicle(s)).Methods("POST")
	r.HandleFunc("/vehicles/{id}", deleteVehicle(s)).Methods("DELETE")

	return r
}
