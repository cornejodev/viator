package rest

import (
	"net/http"

	"github.com/cornejodev/viator/internal/service"
	"github.com/gorilla/mux"
)

func Handler(s service.Service) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/demo", demoHandler(s)).Methods("GET")
	r.HandleFunc("/demo", demoHandlerPost(s)).Methods("POST")

	return r
}
