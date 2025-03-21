package system

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	otellog "go.opentelemetry.io/otel/log"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/trace"
)

// OTelBridge bridges zerolog to OpenTelemetry
type OTelBridge struct {
	console  zerolog.ConsoleWriter
	provider *sdklog.LoggerProvider
	service  string
}

// Create a bridge that forwards zerolog logs to OpenTelemetry and writes to a console writer
func NewOTelBridge(provider *sdklog.LoggerProvider, service string, consoleWriter zerolog.ConsoleWriter) *OTelBridge {
	return &OTelBridge{
		console:  consoleWriter,
		provider: provider,
		service:  service,
	}
}

// Capture zerolog output, write to console, and forward to OpenTelemetry
func (w *OTelBridge) Write(p []byte) (n int, err error) {
	// Parse the JSON log entry to get all fields
	var entry map[string]any
	if err := json.Unmarshal(p, &entry); err != nil {
		return 0, err
	}

	// Create a copy of the entry without trace/span IDs for console output
	consoleEntry := make(map[string]any)
	for k, v := range entry {
		// Skip trace and span IDs in console output
		if k != "traceID" && k != "spanID" {
			consoleEntry[k] = v
		}
	}

	// Convert back to JSON for console output
	consoleJSON, err := json.Marshal(consoleEntry)
	if err != nil {
		return 0, err
	}

	// Write entry to console
	_, err = w.console.Write(consoleJSON)
	if err != nil {
		return n, err
	}

	// Create OTel logger
	logger := w.provider.Logger(w.service)

	// Prepare OTel record
	record := new(otellog.Record)

	// Set timestamp
	if timeStr, ok := entry["time"].(string); ok {
		timestamp, err := time.Parse(time.RFC3339, timeStr)
		if err == nil {
			record.SetTimestamp(timestamp)
		}
	}
	record.SetObservedTimestamp(time.Now())

	// Set body
	msg, _ := entry["message"].(string)
	record.SetBody(otellog.StringValue(msg))

	// Map severity
	level, _ := entry["level"].(string)
	setSeverity(record, level)

	// Extract trace context and add to special attributes that will be recognized by the OTel protocol
	if traceID, ok := entry["traceID"].(string); ok {
		record.AddAttributes(otellog.String("trace_id", traceID))
	}
	if spanID, ok := entry["spanID"].(string); ok {
		record.AddAttributes(otellog.String("span_id", spanID))
	}

	// Add remaining fields as attributes
	for k, v := range entry {
		if k != "level" && k != "message" && k != "time" && k != "traceID" && k != "spanID" {
			record.AddAttributes(convertToAttribute(k, v))
		}
	}

	// Emit the log record
	logger.Emit(context.Background(), *record)

	return len(p), nil
}

// Add trace context to a log event
func AddTraceContext(span trace.Span) func(e *zerolog.Event) {
	return func(e *zerolog.Event) {
		if span == nil {
			return
		}

		spanCtx := span.SpanContext()
		if spanCtx.IsValid() {
			e.Str("traceID", spanCtx.TraceID().String())
			e.Str("spanID", spanCtx.SpanID().String())
		}
	}
}

// Map zerolog level to OTel severity
func setSeverity(record *otellog.Record, level string) {
	switch level {
	case "trace":
		record.SetSeverity(otellog.SeverityTrace)
	case "debug":
		record.SetSeverity(otellog.SeverityDebug)
	case "info":
		record.SetSeverity(otellog.SeverityInfo)
	case "warn":
		record.SetSeverity(otellog.SeverityWarn)
	case "error":
		record.SetSeverity(otellog.SeverityError)
	case "fatal":
		record.SetSeverity(otellog.SeverityFatal)
	case "panic":
		record.SetSeverity(otellog.SeverityFatal)
	default:
		record.SetSeverity(otellog.SeverityInfo)
	}
	record.SetSeverityText(level)
}

// Convert a zerolog field to an OpenTelemetry attribute
func convertToAttribute(key string, value any) otellog.KeyValue {
	switch v := value.(type) {
	case string:
		return otellog.String(key, v)
	case float64:
		return otellog.Float64(key, v)
	case bool:
		return otellog.Bool(key, v)
	case int:
		return otellog.Int(key, v)
	case int64:
		return otellog.Int64(key, v)
	case []any:
		values := make([]otellog.Value, 0, len(v))
		for _, item := range v {
			values = append(values, convertToValue(item))
		}
		return otellog.Slice(key, values...)
	case map[string]any:
		kvs := make([]otellog.KeyValue, 0, len(v))
		for k, val := range v {
			kvs = append(kvs, convertToAttribute(k, val))
		}
		return otellog.Map(key, kvs...)
	default:
		return otellog.String(key, fmt.Sprintf("%+v", v))
	}
}

// Convert a generic value to an OpenTelemetry Value
func convertToValue(value any) otellog.Value {
	switch v := value.(type) {
	case string:
		return otellog.StringValue(v)
	case float64:
		return otellog.Float64Value(v)
	case bool:
		return otellog.BoolValue(v)
	case int:
		return otellog.IntValue(v)
	case int64:
		return otellog.Int64Value(v)
	case []any:
		values := make([]otellog.Value, 0, len(v))
		for _, item := range v {
			values = append(values, convertToValue(item))
		}
		return otellog.SliceValue(values...)
	case map[string]any:
		kvs := make([]otellog.KeyValue, 0, len(v))
		for k, val := range v {
			kvs = append(kvs, convertToAttribute(k, val))
		}
		return otellog.MapValue(kvs...)
	default:
		return otellog.StringValue(fmt.Sprintf("%v", v))
	}
}
