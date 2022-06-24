package rest

import (
	"net/http"

	"github.com/cornejodev/viator/internal/domain/middleware"
	"github.com/cornejodev/viator/internal/service"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

func Handler(s service.Service, lgr zerolog.Logger) http.Handler {
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)

	r.HandleFunc("/vehicles", addVehicle(s, lgr)).Methods("POST")
	r.HandleFunc("/vehicles", listVehicles(s, lgr)).Methods("GET")
	r.HandleFunc("/vehicles/{id}", getVehicle(s, lgr)).Methods("GET")
	r.HandleFunc("/vehicles/{id}", updateVehicle(s, lgr)).Methods("POST")
	r.HandleFunc("/vehicles/{id}", deleteVehicle(s, lgr)).Methods("DELETE")

	return r
}
