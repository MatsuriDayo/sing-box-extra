package hooks

import (
	"context"

	"github.com/sagernet/sing-box/adapter"
	"github.com/sagernet/sing-box/route"
)

type routerPrivateWrapper struct {
	adapter.Router
	sbr *route.Router
}

func (r *routerPrivateWrapper) Initialize(inbounds []adapter.Inbound, outbounds []adapter.Outbound, defaultOutbound func() adapter.Outbound) error {
	return r.sbr.Initialize(inbounds, outbounds, defaultOutbound)
}

func RegisterRouterHook(r func(adapter.Router) adapter.Router) {
	hookRouters = append(hookRouters, r)
}

var hookRouters = make([]func(adapter.Router) adapter.Router, 0)

func HookRouter(ctx context.Context, r adapter.Router) *routerPrivateWrapper {
	sbr := r.(*route.Router)
	for _, f := range hookRouters {
		r = f(r)
	}
	hkr := &routerPrivateWrapper{
		Router: r,
		sbr:    sbr,
	}
	Ctx(ctx).Router = hkr
	return hkr
}
