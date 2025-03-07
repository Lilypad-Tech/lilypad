package system

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	otellog "go.opentelemetry.io/otel/log"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/trace"
)

// OTelWriter bridges zerolog to OpenTelemetry
type OTelWriter struct {
	console  zerolog.ConsoleWriter
	provider *sdklog.LoggerProvider
	service  string
}

// NewOTelWriter creates a writer that forwards zerolog logs to OpenTelemetry
func NewOTelWriter(provider *sdklog.LoggerProvider, service string) *OTelWriter {
	return &OTelWriter{
		console: zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		},
		provider: provider,
		service:  service,
	}
}

// Write captures zerolog output and forwards to OpenTelemetry
func (w *OTelWriter) Write(p []byte) (n int, err error) {
    // Parse the JSON log entry to get all fields
    var entry map[string]interface{}
    if err := json.Unmarshal(p, &entry); err != nil {
        return 0, err
    }

    // Create a copy of the entry without trace/span IDs for console output
    consoleEntry := make(map[string]interface{})
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

    // Write filtered entry to console
    n, err = w.console.Write(consoleJSON)
    if err != nil {
        return n, err
    }

    // Create OTel logger
    logger := w.provider.Logger(w.service)

    // Extract basic fields
    level, _ := entry["level"].(string)
    msg, _ := entry["message"].(string)

    // Create a new record
    record := new(otellog.Record)

    // Set timestamps
    if timeStr, ok := entry["time"].(string); ok {
        timestamp, err := time.Parse(time.RFC3339, timeStr)
        if err == nil {
            record.SetTimestamp(timestamp)
        }
    }
    record.SetObservedTimestamp(time.Now())

    // Set body
    record.SetBody(otellog.StringValue(msg))

    // Map severity
    setSeverity(record, level)

    // Extract trace context and add to special attributes that will be recognized by the OTel protocol
    if traceIDStr, ok := entry["traceID"].(string); ok {
        // Use the official attribute keys defined by OpenTelemetry for trace/span IDs
        // These will be recognized and lifted to the correct protocol level
        record.AddAttributes(otellog.String("trace_id", traceIDStr))
        delete(entry, "traceID") // Remove so it's not added again
    }
    if spanIDStr, ok := entry["spanID"].(string); ok {
        record.AddAttributes(otellog.String("span_id", spanIDStr))
        delete(entry, "spanID") // Remove so it's not added again
    }

    // Add remaining fields as regular attributes
    for k, v := range entry {
        if k != "level" && k != "message" && k != "time" {
            record.AddAttributes(convertToAttribute(k, v))
        }
    }

    // Emit the log record
    logger.Emit(context.Background(), *record)

    return n, nil
}

// AddTraceContext adds trace context to a log event
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

// setSeverity maps zerolog level to OTel severity
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

// convertToAttribute converts a zerolog field to an OpenTelemetry attribute
func convertToAttribute(key string, value interface{}) otellog.KeyValue {
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
	case []interface{}:
		values := make([]otellog.Value, 0, len(v))
		for _, item := range v {
			values = append(values, convertToValue(item))
		}
		return otellog.Slice(key, values...)
	case map[string]interface{}:
		kvs := make([]otellog.KeyValue, 0, len(v))
		for k, val := range v {
			kvs = append(kvs, convertToAttribute(k, val))
		}
		return otellog.Map(key, kvs...)
	default:
		return otellog.String(key, fmt.Sprintf("%v", v))
	}
}

// convertToValue converts a generic value to an OpenTelemetry Value
func convertToValue(value interface{}) otellog.Value {
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
	case []interface{}:
		values := make([]otellog.Value, 0, len(v))
		for _, item := range v {
			values = append(values, convertToValue(item))
		}
		return otellog.SliceValue(values...)
	case map[string]interface{}:
		kvs := make([]otellog.KeyValue, 0, len(v))
		for k, val := range v {
			kvs = append(kvs, convertToAttribute(k, val))
		}
		return otellog.MapValue(kvs...)
	default:
		return otellog.StringValue(fmt.Sprintf("%v", v))
	}
}

// // OTelWriter bridges zerolog to OpenTelemetry
// type OTelWriter struct {
// 	console  zerolog.ConsoleWriter
// 	provider *sdklog.LoggerProvider
// 	service  string
// }

// // NewOTelWriter creates a writer that sends zerolog output to OpenTelemetry
// func NewOTelWriter(provider *sdklog.LoggerProvider, service string) *OTelWriter {
// 	return &OTelWriter{
// 		console: zerolog.ConsoleWriter{
// 			Out:        os.Stdout,
// 			TimeFormat: time.RFC3339,
// 		},
// 		provider: provider,
// 		service:  service,
// 	}
// }

// // Write captures zerolog output and forwards to OpenTelemetry
// func (w *OTelWriter) Write(p []byte) (n int, err error) {
// 	// Parse the JSON log entry to get all fields
// 	var entry map[string]interface{}
// 	if err := json.Unmarshal(p, &entry); err != nil {
// 		return 0, err
// 	}

// 	// Create a copy of the entry without trace/span IDs for console output
// 	consoleEntry := make(map[string]interface{})
// 	for k, v := range entry {
// 		// Skip trace and span IDs in console output
// 		if k != "traceID" && k != "spanID" {
// 			consoleEntry[k] = v
// 		}
// 	}

// 	// Convert back to JSON for console output
// 	consoleJSON, err := json.Marshal(consoleEntry)
// 	if err != nil {
// 		return 0, err
// 	}

// 	// Write filtered entry to console
// 	n, err = w.console.Write(consoleJSON)
// 	if err != nil {
// 		return n, err
// 	}

// 	// Create OTel logger
// 	logger := w.provider.Logger(w.service)

// 	// Extract basic fields
// 	level, _ := entry["level"].(string)
// 	msg, _ := entry["message"].(string)

// 	// Create and emit the record
// 	record := new(otellog.Record)

// 	// Set timestamps
// 	if timeStr, ok := entry["time"].(string); ok {
// 		timestamp, err := time.Parse(time.RFC3339, timeStr)
// 		if err == nil {
// 			record.SetTimestamp(timestamp)
// 		}
// 	}
// 	record.SetObservedTimestamp(time.Now())

// 	// Set body
// 	record.SetBody(otellog.StringValue(msg))

// 	// Map severity
// 	setSeverity(record, level)

// 	// Add all attributes, including trace context
// 	for k, v := range entry {
// 		if k != "level" && k != "message" && k != "time" {
// 			record.AddAttributes(convertToAttribute(k, v))
// 		}
// 	}

// 	// Emit the log record
// 	logger.Emit(context.Background(), *record)

// 	return n, nil
// }

// // Write captures zerolog output and forwards to OpenTelemetry
// // func (w *OTelWriter) Write(p []byte) (n int, err error) {
// // 	// Write to console with nice formatting
// // 	n, err = w.console.Write(p)
// // 	if err != nil {
// // 		return n, err
// // 	}

// // 	// Parse the JSON log entry to get all fields
// // 	var entry map[string]interface{}
// // 	if err := json.Unmarshal(p, &entry); err != nil {
// // 		return n, err
// // 	}

// // 	// Create OTel logger
// // 	logger := w.provider.Logger(w.service)

// // 	// Extract basic fields
// // 	level, _ := entry["level"].(string)
// // 	msg, _ := entry["message"].(string)

// // 	// Create and emit the record
// // 	record := new(otellog.Record)

// // 	// Set timestamps
// // 	if timeStr, ok := entry["time"].(string); ok {
// // 		timestamp, err := time.Parse(time.RFC3339, timeStr)
// // 		if err == nil {
// // 			record.SetTimestamp(timestamp)
// // 		}
// // 	}
// // 	record.SetObservedTimestamp(time.Now())

// // 	// Set body
// // 	record.SetBody(otellog.StringValue(msg))

// // 	// Map severity
// // 	setSeverity(record, level)

// // 	// Add all attributes
// // 	for k, v := range entry {
// // 		if k != "level" && k != "message" && k != "time" {
// // 			record.AddAttributes(convertToAttribute(k, v))
// // 		}
// // 	}

// // 	// Emit the log record
// // 	logger.Emit(context.Background(), *record)

// // 	return n, nil
// // }

// // setSeverity maps zerolog level to OTel severity
// func setSeverity(record *otellog.Record, level string) {
// 	switch level {
// 	case "trace":
// 		record.SetSeverity(otellog.SeverityTrace)
// 	case "debug":
// 		record.SetSeverity(otellog.SeverityDebug)
// 	case "info":
// 		record.SetSeverity(otellog.SeverityInfo)
// 	case "warn":
// 		record.SetSeverity(otellog.SeverityWarn)
// 	case "error":
// 		record.SetSeverity(otellog.SeverityError)
// 	case "fatal":
// 		record.SetSeverity(otellog.SeverityFatal)
// 	case "panic":
// 		record.SetSeverity(otellog.SeverityFatal)
// 	default:
// 		record.SetSeverity(otellog.SeverityInfo)
// 	}
// 	record.SetSeverityText(level)
// }

// // convertToAttribute converts a zerolog field to an OpenTelemetry attribute
// func convertToAttribute(key string, value interface{}) otellog.KeyValue {
// 	switch v := value.(type) {
// 	case string:
// 		return otellog.String(key, v)
// 	case float64:
// 		return otellog.Float64(key, v)
// 	case bool:
// 		return otellog.Bool(key, v)
// 	case int:
// 		return otellog.Int(key, v)
// 	case int64:
// 		return otellog.Int64(key, v)
// 	case []interface{}:
// 		values := make([]otellog.Value, 0, len(v))
// 		for _, item := range v {
// 			values = append(values, convertToValue(item))
// 		}
// 		return otellog.Slice(key, values...)
// 	case map[string]interface{}:
// 		kvs := make([]otellog.KeyValue, 0, len(v))
// 		for k, val := range v {
// 			kvs = append(kvs, convertToAttribute(k, val))
// 		}
// 		return otellog.Map(key, kvs...)
// 	default:
// 		return otellog.String(key, fmt.Sprintf("%v", v))
// 	}
// }

// // convertToValue converts a generic value to an OpenTelemetry Value
// func convertToValue(value interface{}) otellog.Value {
// 	switch v := value.(type) {
// 	case string:
// 		return otellog.StringValue(v)
// 	case float64:
// 		return otellog.Float64Value(v)
// 	case bool:
// 		return otellog.BoolValue(v)
// 	case int:
// 		return otellog.IntValue(v)
// 	case int64:
// 		return otellog.Int64Value(v)
// 	case []interface{}:
// 		values := make([]otellog.Value, 0, len(v))
// 		for _, item := range v {
// 			values = append(values, convertToValue(item))
// 		}
// 		return otellog.SliceValue(values...)
// 	case map[string]interface{}:
// 		kvs := make([]otellog.KeyValue, 0, len(v))
// 		for k, val := range v {
// 			kvs = append(kvs, convertToAttribute(k, val))
// 		}
// 		return otellog.MapValue(kvs...)
// 	default:
// 		return otellog.StringValue(fmt.Sprintf("%v", v))
// 	}
// }

// // AddTraceContext adds trace context to a log event
// func AddTraceContext(span trace.Span) func(e *zerolog.Event) {
// 	return func(e *zerolog.Event) {
// 		if span == nil {
// 			return
// 		}

// 		spanCtx := span.SpanContext()
// 		if spanCtx.IsValid() {
// 			e.Str("traceID", spanCtx.TraceID().String())
// 			e.Str("spanID", spanCtx.SpanID().String())
// 		}
// 	}
// }
