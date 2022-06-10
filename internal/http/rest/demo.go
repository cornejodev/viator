package rest

import (
	"encoding/json"
	"net/http"

	"github.com/cornejodev/viator/internal/service"

	"github.com/cornejodev/viator/internal/domain"
)

func demoHandler(s service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Testing...")
	}
}

func demoHandlerPost(s service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		form := domain.AddDemoForm{}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&form); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err := s.Demo.Add(&domain.Demo{
			Name: form.Name,
		})
		if err != nil {
			resp := newResponse(msgError, "ER004", err.Error(), nil)
			// fmt.Println(resp)
			http.Error(w, resp.messageError.Content, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New demo added")

	}
}
