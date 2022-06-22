package logger

import (
	"io"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func NewLogger(w io.Writer, withTimestamp bool) zerolog.Logger {
	lgr := zerolog.New(w)
	if withTimestamp {
		lgr = lgr.With().Timestamp().Logger()
	}

	return lgr
}

// WriteErrorStackGlobal is a convenience wrapper to set the zerolog
// Global variable ErrorStackMarshaler to write Error stacks for logs
func WriteErrorStackGlobal(writeStack bool) {
	if !writeStack {
		zerolog.ErrorStackMarshaler = nil
		return
	}
	// set ErrorStackMarshaler to pkgerrors.MarshalStack
	// to enable error stack traces
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}
