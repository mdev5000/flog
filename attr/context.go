package attr

import (
	"context"
)

type attrKeyS struct{}

var attrCtxKey attrKeyS

type Chain struct {
	Attrs []Attr
	Prev  *Chain
}

// CtxPrefix adds new attributes that will be logged going forward
func CtxPrefix(ctx context.Context, attrs ...Attr) context.Context {
	existingRaw := ctx.Value(attrCtxKey)
	existing, _ := existingRaw.(*Chain)
	if existing == nil {
		return context.WithValue(ctx, attrCtxKey, &Chain{Attrs: attrs})
	}
	return context.WithValue(ctx, attrCtxKey, &Chain{
		Attrs: attrs,
		Prev:  existing,
	})
}

func FromCtx(ctx context.Context) *Chain {
	existingRaw := ctx.Value(attrCtxKey)
	existing, _ := existingRaw.(*Chain)
	if existing == nil {
		return nil
	}
	return existing
}
