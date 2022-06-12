package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cornejodev/viator/internal/domain"
	"github.com/cornejodev/viator/internal/service"
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
