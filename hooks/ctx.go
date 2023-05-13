package hooks

import (
	"context"

	"github.com/matsuridayo/sing-box-extra/adapter"
	adapter2 "github.com/sagernet/sing-box/adapter"
)

// access some useful resources from hookContextValue

type HookContextKey struct{}

type HookContextValue struct {
	Context    context.Context
	Router     *routerPrivateWrapper
	FakeEngine adapter.FakeDnsEngineI // TODO multi or v6
	BlockOut   adapter2.Outbound
}

func Ctx(ctx context.Context) *HookContextValue {
	if h, ok := ctx.Value((*HookContextKey)(nil)).(*HookContextValue); ok {
		return h
	}
	return nil
}
