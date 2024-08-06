package system

import (
	"context"
	"errors"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type TelemetryConfig struct {
	Service      Service
	CollectorURL string
	Enabled      bool
}

func setupOTelSDK(ctx context.Context, config TelemetryConfig) (shutdown func(context.Context) error, err error) {
	var shutdownFuncs []func(context.Context) error

	// Call registered cleanup handlers, calling each
	// cleanup handler once and joining error results.
	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	// On error, call shutdown for cleanup and return all errors.
	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	// Set up propagator
	prop := newPropagator()
	otel.SetTextMapPropagator(prop)

	// Set up tracer provider.
	if config.Enabled {
		tracerProvider, err := newTracerProvider(ctx, config)
		if err != nil {
			handleErr(err)
			return shutdown, err
		}
		shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
		otel.SetTracerProvider(tracerProvider)
	}

	provider := otel.GetTracerProvider()
	if provider != nil {
		fmt.Printf("*** tracer provider at setup: %+v ***\n", provider)
	}

	// TODO(bgins) Add meter and logger providers

	return shutdown, nil
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

	fmt.Printf("*** Configuring provider with %+v ***\n", config)

	exporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpointURL(config.CollectorURL),
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
