package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/cornejodev/viator/internal/domain/errs"
	"github.com/cornejodev/viator/internal/service"
	"github.com/gorilla/mux"
)

func addVehicle(s service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var rb service.AddVehicleRequest

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&rb); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err := s.Depot.Add(rb)
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

		v, err := s.Depot.Find(id)
		if err != nil {
			log.Println(err)
			http.Error(w, "An error has ocurred. Unable to fetch vehicle.", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(v)
	}
}

func listVehicles(s service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vehicles, err := s.Depot.List()
		if err != nil {
			log.Println(err)
			http.Error(w, "An error has ocurred. Unable to fetch vehicles.", http.StatusBadRequest)
			return
		}
		if len(vehicles) == 0 {
			http.Error(w, "Currently no vehicles in depot.", http.StatusOK)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(vehicles)
	}
}

func updateVehicle(s service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var rb service.UpdateVehicleRequest

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&rb); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// the id in the uri and request body must be the same
		if id != rb.ID {
			http.Error(w, "Route variable and request body IDs do not match.", http.StatusBadRequest)
			return
		}

		err = s.Depot.Update(rb)
		if err != nil {
			log.Println(err)
			http.Error(w, "An error has ocurred. Unable to update vehicle.", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Vehicle updated!")
	}
}

func deleteVehicle(s service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = s.Depot.Remove(id)
		if err != nil {
			if err == errs.ErrVehicleNotFound {
				http.Error(w, "Vehicle requested doesn't exist.", http.StatusBadRequest)
				return
			}
			log.Println(err)
			http.Error(w, "An error has ocurred. Unable to delete vehicle.", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Vehicle deleted!")
	}
}
