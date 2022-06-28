package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/cornejodev/viator/internal/domain/errs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func NewLogger(withTimestamp bool) (zerolog.Logger, error) {
	var op errs.Op = "logger.NewLogger"
	const fname = "../../logs/logs.txt"

	file, err := CreateLogFile(fname)
	if err != nil {
		return zerolog.Logger{}, errs.E(op, err)
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	cw := SetupConsoleWriter()

	multi := zerolog.MultiLevelWriter(cw, file)

	lgr := zerolog.New(multi)
	if withTimestamp {
		lgr = lgr.With().Timestamp().Logger()
	}

	return lgr, nil
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

func CreateLogFile(fname string) (*os.File, error) {
	var op errs.Op = "logger.CreateLogFile"

	err := os.MkdirAll(filepath.Dir(fname), 0755)
	if err != nil && err != os.ErrExist {
		return nil, errs.E(op, err)
	}
	file, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, errs.E(op, err)
	}
	return file, nil
}

func SetupConsoleWriter() zerolog.ConsoleWriter {
	cw := zerolog.ConsoleWriter{Out: os.Stdout,
		// NoColor:       true,
		TimeFormat:    time.UnixDate,
		FieldsExclude: []string{"remote_ip", "user_agent", "message"},
	}

	cw.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	cw.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("(%s)", i)
	}
	cw.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	cw.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}

	return cw
}
