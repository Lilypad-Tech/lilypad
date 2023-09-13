package system

import (
	"fmt"
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

func Error(service Service, title string, err error) {
	log.Error().
		Err(err).
		Msgf(GetServiceString(service, title))
}

func Info(service Service, title string, data interface{}) {
	log.Info().
		Str(GetServiceString(service, title), fmt.Sprintf("%+v", data)).
		Msgf("")
}

func Debug(service Service, title string, data interface{}) {
	log.Debug().
		Str(GetServiceString(service, title), fmt.Sprintf("%+v", data)).
		Msgf("")
}

func Trace(service Service, title string, data interface{}) {
	log.Trace().
		Str(GetServiceString(service, title), fmt.Sprintf("%+v", data)).
		Msgf("")
}
