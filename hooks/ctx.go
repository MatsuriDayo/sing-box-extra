package hooks

import (
	"context"

	"github.com/matsuridayo/sing-box-extra/adapter"
)

// access some useful resources from hookContextValue

type HookContextKey struct{}

type HookContextValue struct {
	Router     *routerPrivateWrapper
	FakeEngine adapter.FakeDnsEngineI // TODO multi or v6
}

func Ctx(ctx context.Context) *HookContextValue {
	if h, ok := ctx.Value((*HookContextKey)(nil)).(*HookContextValue); ok {
		return h
	}
	return nil
}
