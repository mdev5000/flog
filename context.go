package flog

import (
	"context"
	"errors"

	"github.com/mdev5000/flog/attr"
)

type contextKeyS struct{}

var (
	contextKey contextKeyS

	ErrMissingLogger = errors.New("logger not found on context")
)

func FromCtx(ctx context.Context) Logger {
	l, found := TryFromCtx(ctx)
	if !found {
		panic(ErrMissingLogger)
	}
	return l
}

func TryFromCtx(ctx context.Context) (Logger, bool) {
	l, found := TryFromCtxNoPrefix(ctx)
	if found {
		attrs := attr.FromCtx(ctx)
		l = l.PrefixAttr(attrs)
	}
	return l, found
}

func TryFromCtxNoPrefix(ctx context.Context) (Logger, bool) {
	lRaw := ctx.Value(contextKey)
	l, _ := lRaw.(Logger)
	found := l != nil
	return l, found
}

func NewCtx(ctx context.Context, t Logger) context.Context {
	return context.WithValue(ctx, contextKey, t)
}
