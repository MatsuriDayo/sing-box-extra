package hooks

import (
	"context"

	"github.com/matsuridayo/sing-box-extra/adapter"
)

// access some useful resources from hookContextValue

type hookContextKey struct{}

type hookContextValue struct {
	Router     *routerPrivateWrapper
	FakeEngine adapter.FakeDnsEngineI // TODO multi or v6
}

func HookCtx(ctx context.Context) context.Context {
	return context.WithValue(ctx, (*hookContextKey)(nil), &hookContextValue{})
}

func Ctx(ctx context.Context) *hookContextValue {
	return ctx.Value((*hookContextKey)(nil)).(*hookContextValue)
}
