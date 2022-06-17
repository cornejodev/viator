package app

import (
	"log"
	"net/http"
	"time"

	"github.com/cornejodev/viator/config"
	"github.com/cornejodev/viator/internal/http/rest"
	"github.com/cornejodev/viator/internal/service"
	"github.com/cornejodev/viator/internal/storage"
)

func Run(cfg *config.Config) error {
	// Prepare services.
	svc, err := service.New(storage.New(cfg.Database))
	if err != nil {
		return err
	}

	// Setup HTTP server
	mux := rest.Handler(*svc)

	// errs.SetCaller(true) // logging stacktrace

	s := &http.Server{
		Handler:      mux,
		Addr:         cfg.Port,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Printf("Listening on port%s...\n", s.Addr)
	err = s.ListenAndServe()

	return err
}
