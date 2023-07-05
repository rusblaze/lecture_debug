package tracing

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
	"io"
)

// newExporter returns a console exporter.
func newExporter(w io.Writer) (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		stdouttrace.WithPrettyPrint(),
		stdouttrace.WithoutTimestamps(),
	)
}

// newResource returns a resource describing this application.
func newResource(serviceName, serviceVersion string) *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
			semconv.ServiceVersion(serviceVersion),
		),
	)
	return r
}

func newTraceProvider(serviceName, serviceVersion string) (*trace.TracerProvider, error) {
	exp, err := newExporter(io.Discard)
	if err != nil {
		return nil, err
	}

	res := newResource(serviceName, serviceVersion)

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(res),
	)

	return tp, nil
}

func SetupGlobalTracer(serviceName, serviceVersion string) error {
	traceProvider, err := newTraceProvider(serviceName, serviceVersion)
	if err != nil {
		return err
	}
	otel.SetTracerProvider(traceProvider)
	return nil
}
