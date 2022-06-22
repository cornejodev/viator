package app

import (
	"net/http"
	"os"
	"time"

	"github.com/cornejodev/viator/config"
	"github.com/cornejodev/viator/internal/domain/logger"
	"github.com/cornejodev/viator/internal/http/rest"
	"github.com/cornejodev/viator/internal/service"
	"github.com/cornejodev/viator/internal/storage"
	"github.com/rs/zerolog"
)

func Run(cfg *config.Config) error {
	// setup logger
	lgr := logger.NewLogger(os.Stdout, true)

	// set global logging time field format to Unix timestamp
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// set global to log errors with stack (or not) based on flag
	logger.WriteErrorStackGlobal(true)
	lgr.Info().Msgf("log error stack global set to %t", true)

	// Prepare storage
	stg := storage.New(cfg.Database, lgr)

	// Prepare services.
	svc, err := service.New(stg)
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
	lgr.Info().Msgf("Listening on port%s...", s.Addr)
	err = s.ListenAndServe()

	return err
}
