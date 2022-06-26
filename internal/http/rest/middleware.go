package rest

import (
	"context"
	"net"
	"net/http"

	// "github.com/gorilla/context"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
)

func RequestLogger(lgr zerolog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id := xid.New()

			ctx := r.Context()
			ctx = context.WithValue(ctx, "request_id", id.String())
			r = r.WithContext(ctx)

			lgr.Info().
				Str("request_id", id.String()).
				Str("remote_ip", getIP(r.RemoteAddr)).
				Str("user_agent", r.UserAgent()).
				Str("method", r.Method).
				Str("url", r.URL.String()).
				Int("status", 200).
				Msg("Request logged")

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
