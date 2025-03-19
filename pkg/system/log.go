package system

import (
	"fmt"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Global logger

// Configure the default global logger
func SetupGlobalLogger() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	logLevelString := os.Getenv("LOG_LEVEL")
	if logLevelString == "" {
		logLevelString = "info"
	}
	logLevel := zerolog.InfoLevel
	if logLevelString == "none" {
		logLevel = zerolog.NoLevel
	}
	parsedLogLevel, err := zerolog.ParseLevel(logLevelString)
	if err == nil {
		logLevel = parsedLogLevel
	}
	zerolog.CallerSkipFrameCount = 3 // Skip 3 frames (this function, log.Output, log.Logger)
	zerolog.SetGlobalLevel(logLevel)
	log.Logger = log.Output(output).With().Caller().Logger().Level(logLevel)
}

// Logger

// Get a logger configured with contextual fields
func GetLogger(service Service) *zerolog.Logger {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
		FormatMessage: func(m any) string {
			if m == nil {
				return ""
			}
			message := fmt.Sprintf("%+v", m)
			if message != "" {
				return fmt.Sprintf("%s %s", GetServiceBadge(service), message)
			}
			return message
		},
	}

	logger := zerolog.New(output).
		Level(log.Logger.GetLevel()).
		With().
		Timestamp().
		CallerWithSkipFrameCount(2).
		Logger()

	return &logger
}

// Service Logger

type ServiceLogger struct {
	service Service
}

// Get a service logger (deprecated, prefer the GetLogger in new implementations)
func NewServiceLogger(service Service) *ServiceLogger {
	return &ServiceLogger{
		service: service,
	}
}

func (s *ServiceLogger) Error(title string, err error) {
	Error(s.service, title, err)
}

func (s *ServiceLogger) Info(title string, data interface{}) {
	Info(s.service, title, data)
}

func (s *ServiceLogger) Debug(title string, data interface{}) {
	Debug(s.service, title, data)
}

func (s *ServiceLogger) Trace(title string, data interface{}) {
	Trace(s.service, title, data)
}

func logWithCaller(skipFrameCount int, level zerolog.Level, service Service, title string, data interface{}) {
	zerolog.CallerSkipFrameCount = skipFrameCount
	defer func() { zerolog.CallerSkipFrameCount = 3 }() // Reset to the default value

	e := log.WithLevel(level).
		Str(GetServiceString(service, title), fmt.Sprintf("%+v", data))
	e.Caller().Msg("")
}

func Error(service Service, title string, err error) {
	logWithCaller(5, zerolog.ErrorLevel, service, title, err)
}

func Info(service Service, title string, data interface{}) {
	logWithCaller(5, zerolog.InfoLevel, service, title, data)
}

func Debug(service Service, title string, data interface{}) {
	logWithCaller(5, zerolog.DebugLevel, service, title, data)
}

func Trace(service Service, title string, data interface{}) {
	logWithCaller(5, zerolog.TraceLevel, service, title, data)
}

// Dump

func DumpObject(d interface{}) {
	spew.Dump(d)
}

func DumpObjectDebug(d interface{}) {
	currentLogLevel := log.Logger.GetLevel()
	if currentLogLevel <= zerolog.DebugLevel {
		spew.Dump(d)
	}
}

func DumpObjectInfo(d interface{}) {
	currentLogLevel := log.Logger.GetLevel()
	if currentLogLevel <= zerolog.InfoLevel {
		spew.Dump(d)
	}
}
