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
		const op errs.Op = "handlers.addVehicle"
		var rb service.AddVehicleRequest

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&rb); err != nil {
			errs.HTTPErrorResponse(w, err)
			return
		}

		err := s.Depot.Add(rb)
		if err != nil {
			log.Printf("[ERROR] %v \n", errs.E(op, err))
			errs.HTTPErrorResponse(w, err)
			return
		}

		JSON(w, http.StatusOK, nil)
	}
}

func getVehicle(s service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		const op errs.Op = "handlers.getVehicle"

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			errs.HTTPErrorResponse(w, err)
			return
		}

		v, err := s.Depot.Find(id)
		if err != nil {
			log.Printf("[ERROR] %v \n", errs.E(op, err))
			errs.HTTPErrorResponse(w, err)
			return
		}

		JSON(w, http.StatusOK, v)
	}
}

func listVehicles(s service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		const op errs.Op = "handlers.listVehicles"

		vehicles, err := s.Depot.List()
		if err != nil {
			log.Printf("[ERROR] %v \n", errs.E(op, err))
			errs.HTTPErrorResponse(w, err)
			return
		}
		if len(vehicles) == 0 {
			log.Printf("[WARNING] Currently no vehicles in depot.")
		}

		JSON(w, http.StatusOK, vehicles)
	}
}

func updateVehicle(s service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		const op errs.Op = "handlers.updateVehicle"

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			errs.HTTPErrorResponse(w, err)
			return
		}

		var rb service.UpdateVehicleRequest

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&rb); err != nil {
			errs.HTTPErrorResponse(w, err)
			return
		}

		if id != rb.ID {
			err := errs.E(
				op,
				errs.Parameter("id"),
				errs.Code("Route variable and request body IDs do not match."),
				errs.Validation,
			)
			errs.HTTPErrorResponse(w, err)
			return
		}

		err = s.Depot.Update(rb)
		if err != nil {
			log.Printf("[ERROR] %v \n", errs.E(op, err))
			errs.HTTPErrorResponse(w, err)
			return
		}

		JSON(w, http.StatusOK, nil)
	}
}

func deleteVehicle(s service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		const op errs.Op = "handlers.deleteVehicle"

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			errs.HTTPErrorResponse(w, err)
			return
		}

		err = s.Depot.Remove(id)
		if err != nil {
			log.Printf("[ERROR] %v \n", errs.E(op, err))
			errs.HTTPErrorResponse(w, err)
			return
		}

		JSON(w, http.StatusOK, nil)
	}
}
