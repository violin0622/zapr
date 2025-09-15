package ctx

import (
	"context"

	"github.com/go-logr/logr"
)

type (
	fieldkey struct{}
	vKey     struct{}
)

var (
	fkey fieldkey
	vkey vKey
)

func WithValues(ctx context.Context, fields ...any) context.Context {
	return context.WithValue(ctx, fkey, append(Values(ctx), fields...))
}

func Values(ctx context.Context) []any {
	if values, ok := ctx.Value(fkey).([]any); ok {
		return values
	}
	return nil
}

func WithV(ctx context.Context, v int) context.Context {
	return context.WithValue(ctx, vkey, V(ctx)+v)
}

func V(ctx context.Context) int {
	if v, ok := ctx.Value(vkey).(int); ok {
		return v
	}
	return 0
}

func Extract(ctx context.Context, log logr.Logger) logr.Logger {
	return log.V(V(ctx)).WithValues(Values(ctx)...)
}
