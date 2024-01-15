package common

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type ctxKey int

const Key ctxKey = 1

// Values represent state for each request.
type Values struct {
	TraceID    string
	Tracer     trace.Tracer
	Now        time.Time
	StatusCode int
}

// AddSpan adds a OpenTelemetry span to the trace and context.
func AddSpan(ctx context.Context, spanName string, keyValues ...attribute.KeyValue) (context.Context, trace.Span) {
	v, ok := ctx.Value(Key).(*Values)
	if !ok || v.Tracer == nil {
		return ctx, trace.SpanFromContext(ctx)
	}

	ctx, span := v.Tracer.Start(ctx, spanName)
	for _, kv := range keyValues {
		span.SetAttributes(kv)
	}

	return ctx, span
}
