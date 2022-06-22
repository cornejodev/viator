package app

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/cornejodev/viator/config"
	"github.com/cornejodev/viator/internal/http/rest"
	"github.com/cornejodev/viator/internal/service"
	"github.com/cornejodev/viator/internal/storage"
	"github.com/rs/zerolog"
)

func Run(cfg *config.Config) error {
	err := os.MkdirAll(filepath.Dir("logs.txt"), 0755)
	if err != nil && err != os.ErrExist {
		panic(err)
	}
	file, err := os.OpenFile("logs.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	multi := zerolog.MultiLevelWriter(consoleWriter, file)
	lgr := zerolog.New(multi).With().Timestamp().Logger()
	// setup logger
	// lgr := logger.NewLogger(os.Stdout, true)

	// set global logging time field format to Unix timestamp
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// set global to log errors with stack (or not) based on flag
	// logger.WriteErrorStackGlobal(true)
	// lgr.Info().Msgf("log error stack global set to %t", true)

	// Prepare storage
	stg := storage.New(cfg.Database, lgr)

	// Prepare services.
	svc, err := service.New(stg, lgr)
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
