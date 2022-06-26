package rest

import (
	"context"
	"net"
	"net/http"
	"time"

	// "github.com/gorilla/context"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
)

func RequestLogger(lgr zerolog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id := xid.New()
			start := time.Now()

			// context.Set(r, "request_id", id.String()) ->  gorilla/context solution

			ctx := r.Context()
			ctx = context.WithValue(ctx, "request_id", id.String())
			r = r.WithContext(ctx)

			lgr.Info().
				Time("received_time", start).
				Str("request_id", id.String()).
				Str("remote_ip", getIP(r.RemoteAddr)).
				Str("user_agent", r.UserAgent()).
				Str("method", r.Method).
				Str("url", r.URL.String()).
				Int("status", 200).
				Msg("Request received")

			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
		})
	}
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
