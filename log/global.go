package log

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// globalWriter is the global output writer for logs. By default, it is set to os.Stdout.
var globalWriter io.Writer = os.Stdout

// withCaller determines whether the logger includes caller information in the logs.
var withCaller bool = false

// SetGlobalFormat sets the global log output format.
//
// Supported formats:
// - "json": Logs will be written in JSON format (default).
// - Any other value: Logs will be written in a human-readable console format.
//
// Parameters:
// - format: The desired log format ("json" or others).
func SetGlobalFormat(format string) {
	if format == "json" {
		globalWriter = os.Stdout
	} else {
		globalWriter = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
	}
}

// SetGlobalLevel sets the global log level.
//
// Supported levels:
// - "debug", "info", "warn", "error", "fatal", "panic", "trace".
//
// If an invalid level string is provided, the log level defaults to "info",
// and a warning message is logged.
//
// Parameters:
// - levelStr: The desired global log level as a string.
func SetGlobalLevel(levelStr string) {
	level, err := zerolog.ParseLevel(levelStr)
	if err != nil {
		log.Warn().Msgf("Invalid log level '%s', fallback to default \"INFO\"", levelStr)
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(level)
	}
}

// WithCaller enables or disables the inclusion of caller information (file and line number) in logs.
//
// Parameters:
// - enabled: If true, logs will include caller information.
func WithCaller(enabled bool) {
	withCaller = enabled
}
