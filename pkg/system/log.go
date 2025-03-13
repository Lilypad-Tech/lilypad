package system

import (
	"fmt"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	sdklog "go.opentelemetry.io/otel/sdk/log"
)

// Package global logger provider
var loggerProvider *sdklog.LoggerProvider

// Global logger

// Configure the default global logger
func SetupGlobalLogger(service Service, provider *sdklog.LoggerProvider) {
	loggerProvider = provider

	// Set global log level
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
	zerolog.SetGlobalLevel(logLevel)

	// Configure logger with OTel bridge + console or console only
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	if loggerProvider != nil {
		bridge := NewOTelBridge(loggerProvider, string(service), consoleWriter)
		log.Logger = zerolog.New(bridge).
			Level(logLevel).
			With().
			Timestamp().
			CallerWithSkipFrameCount(2).
			Logger()
	} else {
		log.Logger = log.
			Output(consoleWriter).
			Level(logLevel).
			With().
			Caller().
			CallerWithSkipFrameCount(2).
			Logger()
	}
}

// Logger

// Get a logger configured with contextual fields
func GetLogger(service Service) *zerolog.Logger {
	// Configure console writer with service badges
	consoleWriter := zerolog.ConsoleWriter{
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

	// Configure logger with OTel bridge + console or console only
	if loggerProvider != nil {
		bridge := NewOTelBridge(loggerProvider, string(service), consoleWriter)
		logger := zerolog.New(bridge).
			Level(log.Logger.GetLevel()).
			With().
			Timestamp().
			CallerWithSkipFrameCount(2).
			Logger()

		return &logger
	}

	logger := zerolog.New(consoleWriter).
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
	e := log.WithLevel(level).CallerSkipFrame(skipFrameCount).
		Str(GetServiceString(service, title), fmt.Sprintf("%+v", data))
	e.Caller().Msg("")
}

func Error(service Service, title string, err error) {
	logWithCaller(3, zerolog.ErrorLevel, service, title, err)
}

func Info(service Service, title string, data interface{}) {
	logWithCaller(3, zerolog.InfoLevel, service, title, data)
}

func Debug(service Service, title string, data interface{}) {
	logWithCaller(3, zerolog.DebugLevel, service, title, data)
}

func Trace(service Service, title string, data interface{}) {
	logWithCaller(3, zerolog.TraceLevel, service, title, data)
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
