package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/cornejodev/viator/internal/domain"
	"github.com/cornejodev/viator/internal/service"
	"github.com/gorilla/mux"
)

func addVehicle(s service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var form domain.AddVehicleForm

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&form); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err := s.Depot.Add(&domain.Vehicle{
			Type:              form.Type,
			LicensePlate:      form.LicensePlate,
			PassengerCapacity: form.PassengerCapacity,
			Make:              form.Make,
			Model:             form.Model,
			Year:              form.Year,
			Mileage:           form.Mileage,
		})
		if err != nil {
			log.Println(err)
			http.Error(w, "Incorrect or empty parameters in fields", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Vehicle added!")
	}
}

func getVehicle(s service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		vehicle, err := s.Depot.Find(id)
		if err != nil {
			log.Println(err)
			http.Error(w, "An error has ocurred. Unable to fetch vehicle.", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(vehicle)
	}
}
