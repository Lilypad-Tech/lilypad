package system

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	// "go.opentelemetry.io/otel/exporters/otlp"
)

type ServiceLogger struct {
	service Service
}

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
func shouldLog(msg string) bool {
	// Define the string value you want to filter on
	// fmt.Println("msg", msg)
	filterString := "JobOfferAdded"
	// Check if the message contains the filter string
	return strings.Contains(msg, filterString)
}
func SetupLogging() {
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
	// log.Logger = log.Output(output).With().Caller().Logger().Level(logLevel)
	// Setup file logging
	///var/log/lilypad/
	file, err := os.OpenFile("lp.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("Error opening log file")
	}
	diodeWriter := diode.NewWriter(file, 1000, 10*time.Millisecond, func(missed int) {
		log.Printf("Logger dropped %d messages", missed)
	})

	// Combine console and file writers
	multi := zerolog.MultiLevelWriter(output, diodeWriter)

	// Setup ZeroLog with caller information
	zerolog.CallerSkipFrameCount = 3

	tracer := otel.Tracer("logger")
	_, span := tracer.Start(context.Background(), "logging-setup")

	// Add trace and span IDs for file output
	logger := log.Output(multi).With().
		Caller().
		Logger().Level(logLevel)

	filteredLogger := logger.Hook(zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, msg string) {
		if shouldLog(msg) {
			fmt.Println("should log")
			e.Str("trace_id", span.SpanContext().TraceID().String()). // Include trace ID
											Str("span_id", span.SpanContext().SpanID().String()). // Include span ID
											Msg(msg)
		}
	})).Level(logLevel)

	log.Logger = filteredLogger

	// logger := log.Output(multi).With().
	// 	Caller().
	// 	Str("trace_id", span.SpanContext().TraceID().String()). // Include trace ID
	// 	Str("span_id", span.SpanContext().SpanID().String()).   // Include span ID
	// 	Logger().Level(logLevel)
	// // log.Logger = logger
	// filteredLogger := logger.Hook(zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, msg string) {
	// 	if shouldLog(msg) {
	// 		e.Msg(msg)
	// 	}
	// })).Level(logLevel)

	// log.Logger = filteredLogger

	// Initialize OpenTelemetry tracer
	// otel.SetErrorHandler(pkgerrors.NewErrorHandler(pkgerrors.WithStack()))
	//tp := trace.NewNoopTracerProvider()
	//otel.SetTracerProvider(tp)
	//tp := trace.
	// tracerProvider := sdktrace.NewTracerProvider(
	// 	sdktrace.WithBatcher(otlpExporter),
	// 	sdktrace.WithResource(resource.NewWithAttributes(
	// 		semconv.ServiceNameKey.String("your-service-name"),
	// 	)),
	// )
	// tracerProvider := sdktrace.NewTracerProvider(
	// 	sdktrace.WithBatcher(exporter),
	// )
	// Create tracer provider with OTLP exporter
	// tracerProvider := sdktrace.NewTracerProvider(
	// 	sdktrace.WithBatcher(exporter,
	// )
	// otel.SetTracerProvider(tracerProvider)
	// Initialize a context with a span

	// otelzerolog.AddTracingContext()
	//log.Logger = log.Output(multi).With().Caller().Logger().Level(logLevel).Hook(otelzerolog.TraceContextHook{Tracer: tracer, Span: span, Ctx: ctx})
	//log.Logger = log.Output(multi).With().Caller().Logger().Level(logLevel)

	// Log tracing information
	// log.Trace().
	// 	Str("trace_id", span.SpanContext().TraceID().String()).
	// 	Str("span_id", span.SpanContext().SpanID().String()).
	// 	Msg("trace information")
	// log.Logger = log.Output(multi).With().Caller().Logger().Level(logLevel)
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
