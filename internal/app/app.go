package app

import (
	"net/http"
	"time"

	"github.com/cornejodev/viator/config"
	"github.com/cornejodev/viator/internal/domain/errs"
	"github.com/cornejodev/viator/internal/domain/logger"
	"github.com/cornejodev/viator/internal/http/rest"
	"github.com/cornejodev/viator/internal/service"
	"github.com/cornejodev/viator/internal/storage"
)

func Run(cfg *config.Config) error {
	var op errs.Op = "app.Run"

	// setup logger
	lgr, err := logger.NewLogger(true, "logs.txt")
	if err != nil {
		return errs.E(op, err)
	}

	// Prepare storage
	stg, err := storage.New(cfg.Database)
	if err != nil {
		lgr.Error().Err(err).Msg("Cannot start Postgres")
		return errs.E(op, err)
	}
	lgr.Info().Msg("Connected to Postgres")

	// Prepare services.
	svc, err := service.New(stg)
	if err != nil {
		lgr.Error().Err(err).Msg("Cannot start Service")
		return errs.E(op, err)
	}

	// Setup HTTP server
	mux := rest.Handler(*svc, lgr)

	// errs.SetCaller(true) // logging stacktrace

	s := &http.Server{
		Handler:      mux,
		Addr:         cfg.Port,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	lgr.Info().Msgf("Listening on port%s...", s.Addr)
	err = s.ListenAndServe()

	return errs.E(op, err)
}
