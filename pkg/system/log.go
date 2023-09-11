package system

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetupLogging() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	logLevelString := os.Getenv("LOG_LEVEL")
	if logLevelString == "" {
		logLevelString = "info"
	}
	logLevel := zerolog.InfoLevel
	parsedLogLevel, err := zerolog.ParseLevel(logLevelString)
	if err == nil {
		logLevel = parsedLogLevel
	}
	log.Logger = log.Output(output).With().Caller().Logger().Level(logLevel)
}
