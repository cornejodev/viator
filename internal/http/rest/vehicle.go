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

// func addVehicle(s service.Service) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var rb service.AddVehicleRequest

// 		decoder := json.NewDecoder(r.Body)
// 		if err := decoder.Decode(&rb); err != nil {
// 			resp, _ := json.Marshal(newResponse(msgErr, err.Error(), nil))
// 			http.Error(w, string(resp), http.StatusBadRequest)
// 			return
// 		}

// 		err := s.Depot.Add(rb)
// 		if err != nil {
// 			log.Println(err)
// 			resp, _ := json.Marshal(newResponse(msgErr, err.Error(), nil))
// 			http.Error(w, string(resp), http.StatusInternalServerError)
// 			return
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		resp := newResponse(msgOK, "Vehicle added!", nil)
// 		json.NewEncoder(w).Encode(resp)
// 	}
// }

func getVehicle(s service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		op := "getVehicle"
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			resp, _ := json.Marshal(newResponse(msgErr, "", err, nil))
			http.Error(w, string(resp), http.StatusBadRequest)
			return
		}

		v, err := s.Depot.Find(id)
		if err != nil {
			e := errs.E(errs.Op(op), err)
			log.Println(e)

			// info := errs.XUnwrap(e)
			// info.Kind()
			// resp, _ := json.Marshal(newResponse(msgErr, "", err, nil))
			resp := newResponse(msgErr, "", err, nil)
			http.Error(w, resp.String(), resp.Code)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		resp := newResponse(msgOK, "", nil, v)
		json.NewEncoder(w).Encode(resp)
	}
}

// func listVehicles(s service.Service) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		vehicles, err := s.Depot.List()
// 		if err != nil {
// 			log.Println(err)
// 			resp, _ := json.Marshal(newResponse(msgErr, "Unable to fetch vehicle.", nil))
// 			http.Error(w, string(resp), http.StatusInternalServerError)
// 			return
// 		}
// 		if len(vehicles) == 0 {
// 			log.Println("Currently no vehicles in depot.")
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		resp := newResponse(msgOK, "", vehicles)
// 		json.NewEncoder(w).Encode(resp)
// 	}
// }

// func updateVehicle(s service.Service) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		params := mux.Vars(r)
// 		id, err := strconv.Atoi(params["id"])
// 		if err != nil {
// 			resp, _ := json.Marshal(newResponse(msgErr, err.Error(), nil))
// 			http.Error(w, string(resp), http.StatusBadRequest)
// 			return
// 		}

// 		var rb service.UpdateVehicleRequest

// 		decoder := json.NewDecoder(r.Body)
// 		if err := decoder.Decode(&rb); err != nil {
// 			resp, _ := json.Marshal(newResponse(msgErr, err.Error(), nil))
// 			http.Error(w, string(resp), http.StatusBadRequest)
// 			return
// 		}

// 		// the id in the uri and request body must be the same
// 		if id != rb.ID {
// 			resp, _ := json.Marshal(newResponse(msgErr, "Route variable and request body IDs do not match.", nil))
// 			http.Error(w, string(resp), http.StatusBadRequest)
// 			return
// 		}

// 		err = s.Depot.Update(rb)
// 		if err != nil {
// 			log.Println(err)
// 			resp, _ := json.Marshal(newResponse(msgErr, "Unable to fetch vehicle.", nil))
// 			http.Error(w, string(resp), http.StatusInternalServerError)
// 			return
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		resp := newResponse(msgOK, "Vehicle updated", nil)
// 		json.NewEncoder(w).Encode(resp)
// 	}
// }

// func deleteVehicle(s service.Service) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		params := mux.Vars(r)
// 		id, err := strconv.Atoi(params["id"])
// 		if err != nil {
// 			resp, _ := json.Marshal(newResponse(msgErr, err.Error(), nil))
// 			http.Error(w, string(resp), http.StatusBadRequest)
// 		}

// 		err = s.Depot.Remove(id)
// 		if err != nil {
// 			log.Println(err)
// 			resp, _ := json.Marshal(newResponse(msgErr, "Unable to delete vehicle.", nil))
// 			http.Error(w, string(resp), http.StatusInternalServerError)
// 			return
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		resp := newResponse(msgOK, "Vehicle deleted!", nil)
// 		json.NewEncoder(w).Encode(resp)
// 	}
// }
