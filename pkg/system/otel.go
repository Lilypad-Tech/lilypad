package system

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"go.opentelemetry.io/otel"
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

type TelemetryConfig struct {
	Service Service
	URL     string
	Enabled bool
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

	// Set up tracer provider.
	var TracerProvider *trace.TracerProvider
	if config.Enabled {
		TracerProvider, err = newTracerProvider(ctx, config)
		if err != nil {
			handleErr(err)
			return Telemetry{
				TracerProvider,
				Shutdown,
			}, err
		}
		shutdownFuncs = append(shutdownFuncs, TracerProvider.Shutdown)
	} else {
		// TODO(bgins) Find a better Noop provider option
		TracerProvider = trace.NewTracerProvider()
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

func newTracerProvider(ctx context.Context, config TelemetryConfig) (*trace.TracerProvider, error) {
	// exporter, err := stdouttrace.New(
	// 	stdouttrace.WithPrettyPrint(),
	// )
	// if err != nil {
	// 	fmt.Println("*** Failed to configure exporter ***")
	// 	return nil, err
	// }

	// fmt.Printf("*** Configuring provider with %+v ***\n", config)

	exporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpointURL(config.URL),
		// TODO Add auth
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		fmt.Println("failed to configure trace exporter")
		return nil, err
	}

	// Set resource with global attributes
	resource := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(string(config.Service)),
	)

	provider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource),
	)

	return provider, nil
}

// Convert service names to use standardized OTel underscores
func GetOTelServiceName(service Service) string {
	return strings.Replace(string(service), "-", "_", -1)
}
