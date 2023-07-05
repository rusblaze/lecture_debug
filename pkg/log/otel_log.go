package log

import (
	"context"
	"github.com/quay/zlog"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
)

const (
	TRACE_KEY = "trace_id"
	SPAN_KEY  = "span_id"
)

// AddCtx is the workhorse function that every facade function calls.
//
// If the passed Event is enabled, it will attach all the otel baggage to
// it and return it.
func addCtx(ctx context.Context, ev *zerolog.Event) *zerolog.Event {
	if !ev.Enabled() {
		return ev
	}

	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.HasTraceID() && spanCtx.TraceID().IsValid() {
		ev.Str(TRACE_KEY, spanCtx.TraceID().String())
	}
	if spanCtx.HasSpanID() && spanCtx.SpanID().IsValid() {
		ev.Str(SPAN_KEY, spanCtx.SpanID().String())
	}

	return ev
}

// Log starts a new message with no level.
func Log(ctx context.Context) *zerolog.Event {
	return addCtx(ctx, zlog.Log(ctx))
}

// WithLevel starts a new message with the specified level.
func WithLevel(ctx context.Context, l zerolog.Level) *zerolog.Event {
	return addCtx(ctx, zlog.WithLevel(ctx, l))
}

// Trace starts a new message with the trace level.
func Trace(ctx context.Context) *zerolog.Event {
	return addCtx(ctx, zlog.Trace(ctx))
}

// Debug starts a new message with the debug level.
func Debug(ctx context.Context) *zerolog.Event {
	return addCtx(ctx, zlog.Debug(ctx))
}

// Info starts a new message with the infor level.
func Info(ctx context.Context) *zerolog.Event {
	return addCtx(ctx, zlog.Info(ctx))
}

// Warn starts a new message with the warn level.
func Warn(ctx context.Context) *zerolog.Event {
	return addCtx(ctx, zlog.Warn(ctx))
}

// Error starts a new message with the error level.
func Error(ctx context.Context) *zerolog.Event {
	return addCtx(ctx, zlog.Error(ctx))
}
