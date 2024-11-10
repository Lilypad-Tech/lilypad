package system

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"runtime"
	"strings"

	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type Telemetry struct {
	TracerProvider *trace.TracerProvider
	Shutdown       func(context.Context) error
}

type TelemetryOptions struct {
	URL     string `json:"url" toml:"url"`
	Token   string `json:"token" toml:"token"`
	Disable bool
}

type TelemetryConfig struct {
	TelemetryURL   string
	TelemetryToken string
	Enabled        bool
	Service        Service
	Network        string
	Address        string
	GPU            []string
}

func SetupOTelSDK(ctx context.Context, config TelemetryConfig) (telemetry Telemetry, err error) {
	var shutdownFuncs []func(context.Context) error

	// Call registered cleanup handlers, calling each
	// cleanup handler once and joining error results.
	Shutdown := func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	// On error, call shutdown for cleanup and return all errors.
	handleErr := func(inErr error) {
		err = errors.Join(inErr, Shutdown(ctx))
	}

	// Set up propagator
	prop := newPropagator()
	otel.SetTextMapPropagator(prop)

	// Set resource with global attributes
	resource := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(GetOTelServiceName(config.Service)),
		semconv.ServiceVersionKey.String(Version),
		attribute.String("system.os", runtime.GOOS),
		attribute.String("system.arch", runtime.GOARCH),
		attribute.StringSlice("system.gpu", config.GPU),
		attribute.String("chain.network", config.Network),
		attribute.String("chain.address", config.Address),
	)

	// TODO(bgins) Investigate a better Noop provider
	TracerProvider := trace.NewTracerProvider()

	// Set up tracer provider.
	if config.Enabled {
		TracerProvider, err = newTracerProvider(ctx, resource, config)
		if err != nil {
			handleErr(err)
			return Telemetry{
				TracerProvider,
				Shutdown,
			}, err
		}
		shutdownFuncs = append(shutdownFuncs, TracerProvider.Shutdown)
	}

	// TODO(bgins) Add meter and logger providers

	return Telemetry{TracerProvider, Shutdown}, nil
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func newTracerProvider(ctx context.Context, resource *resource.Resource, config TelemetryConfig) (*trace.TracerProvider, error) {
	exporter, err := newTracerExporter(ctx, config)
	if err != nil {
		log.Error().Msgf("failed to configure trace exporter: %v", err)
		return nil, err
	}

	// Set tracer provider
	provider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource),
	)

	return provider, nil
}

func newTracerExporter(ctx context.Context, config TelemetryConfig) (*otlptrace.Exporter, error) {
	headers := map[string]string{"Authorization": fmt.Sprintf("Bearer %s", config.TelemetryToken)}
	url, err := url.ParseRequestURI(config.TelemetryURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse telemetry URL: %s", err)
	}

	var exporter *otlptrace.Exporter
	if url.Scheme == "https" {
		exporter, err = otlptracehttp.New(ctx,
			otlptracehttp.WithHeaders(headers),
			otlptracehttp.WithEndpointURL(config.TelemetryURL),
		)
		if err != nil {
			return nil, err
		}
	} else {
		exporter, err = otlptracehttp.New(ctx,
			otlptracehttp.WithHeaders(headers),
			otlptracehttp.WithEndpointURL(config.TelemetryURL),
			otlptracehttp.WithInsecure(),
		)
		if err != nil {
			return nil, err
		}
	}

	return exporter, nil
}

// Convert service names to use standardized OTel underscores
func GetOTelServiceName(service Service) string {
	return strings.Replace(string(service), "-", "_", -1)
}
