package system

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetupLogging() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	logLevel := zerolog.InfoLevel
	parsedLogLevel, err := zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err == nil {
		logLevel = parsedLogLevel
	}
	log.Logger = log.Output(output).With().Caller().Logger().Level(logLevel)
}
