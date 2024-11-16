package log

import (
	"github.com/rs/zerolog"
)

// NewLogger creates a new instance of a zerolog.Logger configured with global settings.
//
// The logger includes timestamps by default. If `withCaller` is enabled,
// it also includes caller information (file and line number).
//
// Returns:
// - A new instance of zerolog.Logger.
func NewLogger() zerolog.Logger {
	if withCaller {
		return zerolog.New(globalWriter).With().Timestamp().Caller().Logger()
	}

	return zerolog.New(globalWriter).With().Timestamp().Logger()
}
