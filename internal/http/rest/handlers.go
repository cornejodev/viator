package rest

import (
	"encoding/json"
	"net/http"

	"github.com/cornejodev/viator/internal/domain/errs"
	"github.com/cornejodev/viator/internal/service"
	"github.com/rs/zerolog"
)

func addVehicle(s service.Service, lgr zerolog.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		const op errs.Op = "handlers.addVehicle"
		var rb service.AddVehicleRequest

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&rb); err != nil {
			err = errs.E(op, err)
			errs.HTTPErrorResponse(w, r, lgr, err)
			return
		}

		err := s.Depot.Add(r.Context(), rb)
		if err != nil {
			err = errs.E(op, err)
			errs.HTTPErrorResponse(w, r, lgr, err)
			return
		}

		JSON(w, http.StatusOK, nil)
	}
}

// func getVehicle(s service.Service, lgr zerolog.Logger) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		const op errs.Op = "handlers.getVehicle"

// 		params := mux.Vars(r)
// 		id, err := strconv.Atoi(params["id"])
// 		if err != nil {
// 			err = errs.E(op, err)
// 			errs.HTTPErrorResponse(w, r, lgr, err)
// 			return
// 		}

// 		v, err := s.Depot.Find(id)
// 		if err != nil {
// 			err = errs.E(op, err)
// 			errs.HTTPErrorResponse(w, r, lgr, err)
// 			return
// 		}

// 		JSON(w, http.StatusOK, v)
// 	}
// }

// func listVehicles(s service.Service, lgr zerolog.Logger) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		const op errs.Op = "handlers.listVehicles"

// 		vehicles, err := s.Depot.List()
// 		if err != nil {
// 			err = errs.E(op, err)
// 			errs.HTTPErrorResponse(w, r, lgr, err)
// 			return
// 		}

// 		JSON(w, http.StatusOK, vehicles)
// 	}
// }

// func updateVehicle(s service.Service, lgr zerolog.Logger) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		const op errs.Op = "handlers.updateVehicle"

// 		params := mux.Vars(r)
// 		id, err := strconv.Atoi(params["id"])
// 		if err != nil {
// 			err = errs.E(op, err)
// 			errs.HTTPErrorResponse(w, r, lgr, err)
// 			return
// 		}

// 		var rb service.UpdateVehicleRequest

// 		decoder := json.NewDecoder(r.Body)
// 		if err := decoder.Decode(&rb); err != nil {
// 			err = errs.E(op, err)
// 			errs.HTTPErrorResponse(w, r, lgr, err)
// 			return
// 		}

// 		if id != rb.ID {
// 			err := errs.E(
// 				op,
// 				errs.Parameter("id"),
// 				errs.Code("route variable and request body IDs do not match."),
// 				errs.Validation,
// 			)
// 			errs.HTTPErrorResponse(w, r, lgr, err)
// 			return
// 		}

// 		err = s.Depot.Update(rb)
// 		if err != nil {
// 			err = errs.E(op, err)
// 			errs.HTTPErrorResponse(w, r, lgr, err)
// 			return
// 		}

// 		JSON(w, http.StatusOK, nil)
// 	}
// }

// func deleteVehicle(s service.Service, lgr zerolog.Logger) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		const op errs.Op = "handlers.deleteVehicle"

// 		params := mux.Vars(r)
// 		id, err := strconv.Atoi(params["id"])
// 		if err != nil {
// 			err = errs.E(op, err)
// 			errs.HTTPErrorResponse(w, r, lgr, err)
// 			return
// 		}

// 		err = s.Depot.Remove(id)
// 		if err != nil {
// 			err = errs.E(op, err)
// 			errs.HTTPErrorResponse(w, r, lgr, err)
// 			return
// 		}

// 		JSON(w, http.StatusOK, nil)
// 	}
// }

// // decoderErr is a convenience function to handle errors returned by
// // json.NewDecoder(r.Body).Decode(&data) and return the appropriate
// // error response
// func decoderErr(err error) error {
// 	switch {
// 	// If the request body is empty (io.EOF)
// 	// return an error
// 	case err == io.EOF:
// 		return errs.E(errs.InvalidRequest, "Request Body cannot be empty")
// 	// If the request body has malformed JSON (io.ErrUnexpectedEOF)
// 	// return an error
// 	case err == io.ErrUnexpectedEOF:
// 		return errs.E(errs.InvalidRequest, "Malformed JSON")
// 	// return other errors
// 	case err != nil:
// 		return errs.E(err)
// 	}
// 	return nil
// }
