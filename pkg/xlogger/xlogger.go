package xlogger

import (
	"os"
	"primeskills-test-api/internal/config"
	"time"

	"github.com/rs/zerolog"
)

var (
	Logger *zerolog.Logger
)

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
