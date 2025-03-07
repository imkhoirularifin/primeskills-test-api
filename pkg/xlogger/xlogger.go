package xlogger

import (
	"os"
	"primeskills-test-api/pkg/config"
	"time"

	"github.com/rs/zerolog"
)

var (
	Logger *zerolog.Logger
)

// Setup initializes the global logger based on the configuration.
// If the application is in development mode, it sets up a console logger with debug level.
// Otherwise, it sets up a standard logger that writes to stderr.
func Setup() {
	cfg := config.Cfg
	if cfg.IsDevelopment {
		l := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Timestamp().Logger()
		l.Level(zerolog.DebugLevel)
		Logger = &l
		return
	}
	l := zerolog.New(os.Stderr).With().Timestamp().Logger()
	Logger = &l
}
