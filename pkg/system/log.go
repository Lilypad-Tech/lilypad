package system

import (
	"fmt"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/vincentfree/opentelemetry/otelzerolog"
	otelzlog "go.opentelemetry.io/contrib/bridges/otelzerolog"
	sdklog "go.opentelemetry.io/otel/sdk/log"
)

// Global logger

var provider *sdklog.LoggerProvider

// Configure the default global logger
func SetupGlobalLogger(service Service, loggerProvider *sdklog.LoggerProvider) {
	provider = loggerProvider

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
	// zerolog.CallerSkipFrameCount = 3 // Skip 3 frames (this function, log.Output, log.Logger)
	zerolog.SetGlobalLevel(logLevel)

	if provider != nil {
		writer := NewOTelWriter(provider, string(service))
		logOpts := []otelzerolog.LogOption{
			otelzerolog.WithOtelBridge(string(service), otelzlog.WithLoggerProvider(provider)),
		}
		log.Logger = otelzerolog.New(logOpts...).
			Output(writer).
			Level(logLevel).
			With().
			Timestamp().
			CallerWithSkipFrameCount(3).
			Logger()
	} else {
		log.Logger = log.Output(output).With().Caller().Logger().Level(logLevel)
	}

	////// Works but missing OTel attributes

	// if provider != nil {
	// 	logOpts := []otelzerolog.LogOption{
	// 		otelzerolog.WithOtelBridge(string(service), otelzlog.WithLoggerProvider(provider)),
	// 	}
	// 	log.Logger = otelzerolog.New(logOpts...).
	// 		Output(output).
	// 		Level(logLevel).
	// 		With().
	// 		Timestamp().
	// 		CallerWithSkipFrameCount(3).
	// 		Logger()
	// } else {
	// 	log.Logger = log.Output(output).With().Caller().Logger().Level(logLevel)
	// }

	/////// Multi writer

	// if provider != nil {
	// 	logOpts := []otelzerolog.LogOption{
	// 		otelzerolog.WithOtelBridge(string(service), otelzlog.WithLoggerProvider(provider)),
	// 	}
	// 	otelLogger := otelzerolog.New(logOpts...)

	// 	multiWriter := zerolog.MultiLevelWriter(output, otelLogger.Logger)
	// 	logger := zerolog.New(multiWriter).
	// 		Level(logLevel).
	// 		With().
	// 		Timestamp().
	// 		CallerWithSkipFrameCount(3).
	// 		Logger()

	// 	log.Logger = logger
	// } else {
	// 	log.Logger = log.Output(output).With().Caller().Logger().Level(logLevel)
	// }
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

	if provider != nil {
		writer := NewOTelWriter(provider, string(service))
		logOpts := []otelzerolog.LogOption{
			otelzerolog.WithOtelBridge(string(service), otelzlog.WithLoggerProvider(provider)),
		}
		logger := otelzerolog.New(logOpts...).
			Output(writer).
			Level(log.Logger.GetLevel()).
			With().
			Timestamp().
			CallerWithSkipFrameCount(3).
			Logger()

		return &logger
	}

	////// Works but missing OTel attributes

	// if provider != nil {
	// 	logOpts := []otelzerolog.LogOption{
	// 		otelzerolog.WithOtelBridge(string(service), otelzlog.WithLoggerProvider(provider)),
	// 	}
	// 	logger := otelzerolog.New(logOpts...).
	// 		Output(output).
	// 		Level(log.Logger.GetLevel()).
	// 		With().
	// 		Timestamp().
	// 		CallerWithSkipFrameCount(2).
	// 		Logger()

	// 	return &logger
	// }

	/////// Multi writer

	// if provider != nil {
	// 	logOpts := []otelzerolog.LogOption{
	// 		otelzerolog.WithOtelBridge(string(service), otelzlog.WithLoggerProvider(provider)),
	// 	}
	// 	otelLogger := otelzerolog.New(logOpts...)

	// 	multiWriter := zerolog.MultiLevelWriter(output, otelLogger.Logger)
	// 	logger := zerolog.New(multiWriter).
	// 		Level(log.Logger.GetLevel()).
	// 		With().
	// 		Timestamp().
	// 		CallerWithSkipFrameCount(2).
	// 		Logger()

	// 	return &logger
	// }

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
