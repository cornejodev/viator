package middleware

import (
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		logger := zerolog.New(multi).With().Timestamp().Logger()

		// id := uuid.New()
		start := time.Now()
		// logger := zerolog.New(os.Stdout)

		logger.Info().
			Time("received_time", start).
			Str("remote_ip", getIP(r.RemoteAddr)).
			Str("user_agent", r.UserAgent()).
			// Str("request_id", id.String()). TODO: propagate request_id in order to connect both error and info logs
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Int("status", 200).
			Msg("")

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// get ip from host port
func getIP(hp string) string {
	h, _, err := net.SplitHostPort(hp)
	if err != nil {
		return ""
	}
	if len(h) > 0 && h[0] == '[' {
		return h[1 : len(h)-1]
	}
	return h
}

// func LoggingMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Println(r.RequestURI)

// 		defer func(startedAt time.Time) {
// 			log.Println(r.RequestURI, time.Since(startedAt))
// 		}(time.Now())

// 		next.ServeHTTP(w, r)
// 	})
// }

// func LoggingMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		log := zerolog.New(os.Stdout).With().
// 			Timestamp().
// 			Str("role", "my-service").
// 			Str("host", "local-hostname").
// 			Logger()
// 		hlog.NewHandler(log)

// 		// Call the next handler, which can be another middleware in the chain, or the final handler.
// 		next.ServeHTTP(w, r)
// 	})
// }
