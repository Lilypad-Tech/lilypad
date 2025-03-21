//go:build unit

package system

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/log"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

func TestOTelBridge_Write(t *testing.T) {
	// Setup logger provider with mock processor
	mockProcessor := &mockProcessor{records: []*sdklog.Record{}}
	provider := sdklog.NewLoggerProvider(sdklog.WithProcessor(mockProcessor))

	// Setup console output capture and bridge
	consoleOutput := &bytes.Buffer{}
	consoleWriter := zerolog.ConsoleWriter{Out: consoleOutput, NoColor: true}
	bridge := NewOTelBridge(provider, "test-service", consoleWriter)

	// Create test log entry
	now := time.Now().UTC()
	entry := map[string]any{
		"level":   "error",
		"time":    now.Format(time.RFC3339),
		"message": "Test error message",
		"count":   42,
		"enabled": true,
		"traceID": "0123456789abcdef0123456789abcdef",
		"spanID":  "fedcba9876543210",
	}
	entryJSON, err := json.Marshal(entry)
	require.NoError(t, err)

	// Write to OTel and console writer
	n, err := bridge.Write(entryJSON)

	// Verify basics
	assert.NoError(t, err)
	assert.Equal(t, len(entryJSON), n)

	// Verify console output
	assert.Contains(t, consoleOutput.String(), "Test error message")
	assert.Contains(t, consoleOutput.String(), "42")
	assert.Contains(t, consoleOutput.String(), "true")
	// Trace context should not be in console output
	assert.NotContains(t, consoleOutput.String(), "0123456789abcdef0123456789abcdef")
	assert.NotContains(t, consoleOutput.String(), "fedcba9876543210")

	// Verify OTel record
	require.Len(t, mockProcessor.records, 1)
	record := mockProcessor.records[0]
	body := record.Body().AsString()
	assert.Equal(t, "Test error message", body)
	assert.Equal(t, log.SeverityError, record.Severity())
	assert.Equal(t, "error", record.SeverityText())
	// Check timestamp approximation
	assert.WithinDuration(t, now, record.Timestamp(), 2*time.Second)

	// Check attributes
	foundCount := false
	foundEnabled := false
	foundTraceID := false
	foundSpanID := false

	// Walk attribures checking existence and values
	record.WalkAttributes(func(attr log.KeyValue) bool {
		switch attr.Key {
		case "count":
			foundCount = true
			val := attr.Value.AsFloat64()
			assert.Equal(t, float64(42), val)
		case "enabled":
			foundEnabled = true
			val := attr.Value.AsBool()
			assert.Equal(t, true, val)
		case "trace_id":
			foundTraceID = true
			val := attr.Value.AsString()
			assert.Equal(t, "0123456789abcdef0123456789abcdef", val)
		case "span_id":
			foundSpanID = true
			val := attr.Value.AsString()
			assert.Equal(t, "fedcba9876543210", val)
		}
		return true
	})

	assert.True(t, foundCount, "Should have found 'count' attribute")
	assert.True(t, foundEnabled, "Should have found 'enabled' attribute")
	assert.True(t, foundTraceID, "Should have found 'trace_id' attribute")
	assert.True(t, foundSpanID, "Should have found 'span_id' attribute")
}

func TestAddTraceContext(t *testing.T) {
	// Create a mock span context
	mockTraceID, _ := trace.TraceIDFromHex("0123456789abcdef0123456789abcdef")
	mockSpanID, _ := trace.SpanIDFromHex("0123456789abcdef")
	mockSpanContext := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID: mockTraceID,
		SpanID:  mockSpanID,
		Remote:  false,
	})

	tests := []struct {
		name        string
		mockSpan    func() trace.Span
		expectTrace bool
		expectSpan  bool
	}{
		{
			name: "valid span",
			mockSpan: func() trace.Span {
				// Create a noop tracer and wrap our context in it
				tracer := noop.Tracer{}
				ctx := trace.ContextWithSpanContext(context.Background(), mockSpanContext)
				_, span := tracer.Start(ctx, "test-span")
				return span
			},
			expectTrace: true,
			expectSpan:  true,
		},
		{
			name: "nil span",
			mockSpan: func() trace.Span {
				return nil
			},
			expectTrace: false,
			expectSpan:  false,
		},
		{
			name: "invalid span context",
			mockSpan: func() trace.Span {
				// Create a noop tracer with a context that is missing
				// trace and span IDs
				tracer := noop.Tracer{}
				ctx := trace.ContextWithSpanContext(context.Background(), trace.SpanContext{})
				_, span := tracer.Start(ctx, "test-span")
				return span
			},
			expectTrace: false,
			expectSpan:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a zerolog event and capture its output
			output := &bytes.Buffer{}
			logger := zerolog.New(output)
			event := logger.Info()

			// Add trace context to event
			AddTraceContext(tt.mockSpan())(event)

			// Write message
			event.Msg("test message")

			// Parse output
			var entry map[string]any
			err := json.Unmarshal(output.Bytes(), &entry)
			require.NoError(t, err)

			// Check if trace/span IDs were added
			_, hasTraceID := entry["traceID"]
			_, hasSpanID := entry["spanID"]

			assert.Equal(t, tt.expectTrace, hasTraceID)
			assert.Equal(t, tt.expectSpan, hasSpanID)

			if tt.expectTrace {
				assert.Equal(t, mockTraceID.String(), entry["traceID"])
			}
			if tt.expectSpan {
				assert.Equal(t, mockSpanID.String(), entry["spanID"])
			}
		})
	}
}

func TestSetSeverity(t *testing.T) {
	tests := []struct {
		level          string
		expectedSev    log.Severity
		expectedSevTxt string
	}{
		{"trace", log.SeverityTrace, "trace"},
		{"debug", log.SeverityDebug, "debug"},
		{"info", log.SeverityInfo, "info"},
		{"warn", log.SeverityWarn, "warn"},
		{"error", log.SeverityError, "error"},
		{"fatal", log.SeverityFatal, "fatal"},
		{"panic", log.SeverityFatal, "panic"},
		{"unknown", log.SeverityInfo, "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.level, func(t *testing.T) {
			record := new(log.Record)
			setSeverity(record, tt.level)

			assert.Equal(t, tt.expectedSev, record.Severity())
			assert.Equal(t, tt.expectedSevTxt, record.SeverityText())
		})
	}
}

func TestConvertToAttribute(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		value    any
		validate func(t *testing.T, kv log.KeyValue)
	}{
		{
			name:  "string",
			key:   "str",
			value: "hello",
			validate: func(t *testing.T, kv log.KeyValue) {
				assert.Equal(t, "str", kv.Key)
				v := kv.Value.AsString()
				assert.Equal(t, "hello", v)
			},
		},
		{
			name:  "float64",
			key:   "num",
			value: 3.14,
			validate: func(t *testing.T, kv log.KeyValue) {
				assert.Equal(t, "num", kv.Key)
				v := kv.Value.AsFloat64()
				assert.Equal(t, 3.14, v)
			},
		},
		{
			name:  "bool",
			key:   "flag",
			value: true,
			validate: func(t *testing.T, kv log.KeyValue) {
				assert.Equal(t, "flag", kv.Key)
				v := kv.Value.AsBool()
				assert.Equal(t, true, v)
			},
		},
		{
			name:  "int",
			key:   "count",
			value: 42,
			validate: func(t *testing.T, kv log.KeyValue) {
				assert.Equal(t, "count", kv.Key)
				v := kv.Value.AsInt64()
				assert.Equal(t, int64(42), v)
			},
		},
		{
			name:  "slice",
			key:   "items",
			value: []any{"a", "b", "c"},
			validate: func(t *testing.T, kv log.KeyValue) {
				assert.Equal(t, "items", kv.Key)
				assert.Equal(t, log.KindSlice, kv.Value.Kind())

				// Check slice values
				sliceVal := kv.Value.AsSlice()
				assert.Equal(t, 3, len(sliceVal), "Slice should have 3 elements")

				// Check each element in the slice
				values := []string{}
				for _, val := range sliceVal {
					values = append(values, val.AsString())
				}
				assert.Equal(t, []string{"a", "b", "c"}, values)
			},
		},
		{
			name:  "map",
			key:   "data",
			value: map[string]any{"a": 1, "b": "two"},
			validate: func(t *testing.T, kv log.KeyValue) {
				assert.Equal(t, "data", kv.Key)
				assert.Equal(t, log.KindMap, kv.Value.Kind())

				// Check map values
				kvPairs := kv.Value.AsMap()
				assert.Equal(t, 2, len(kvPairs), "Map should have two key-value pairs")

				// Create a map to easily look up values by key
				valueMap := make(map[string]log.Value)
				for _, pair := range kvPairs {
					valueMap[pair.Key] = pair.Value
				}

				// Check expected keys and values
				if val, ok := valueMap["a"]; ok {
					assert.Equal(t, int64(1), val.AsInt64())
				}

				if val, ok := valueMap["b"]; ok {
					assert.Equal(t, "two", val.AsString())
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv := convertToAttribute(tt.key, tt.value)
			tt.validate(t, kv)
		})
	}
}

// Mock log processor

type mockProcessor struct {
	records []*sdklog.Record
}

// Store records for test cases
func (p *mockProcessor) OnEmit(ctx context.Context, record *sdklog.Record) error {
	p.records = append(p.records, record)
	return nil
}

func (p *mockProcessor) ForceFlush(ctx context.Context) error {
	return nil
}

func (p *mockProcessor) Shutdown(ctx context.Context) error {
	return nil
}
