package fakedns

import (
	"context"
	"fmt"
	"net"
	"net/netip"

	"github.com/matsuridayo/sing-box-extra/hooks"
	"github.com/sagernet/sing-box/adapter"
	"github.com/sagernet/sing/common/network"
)

func init() {
	hooks.RegisterRouterHook(func(r adapter.Router) adapter.Router {
		return &RouterWithFakeIP{r}
	})
}

type RouterWithFakeIP struct {
	adapter.Router
}

func (r *RouterWithFakeIP) RouteConnection(ctx context.Context, conn net.Conn, metadata adapter.InboundContext) error {
	h := hooks.Ctx(ctx)
	if h != nil && h.FakeEngine != nil {
		ip := metadata.Destination.IPAddr().IP
		if h.FakeEngine.IsIPInPool(ip) {
			if d, err := h.FakeEngine.RestoreToDomain(ip); err == nil {
				metadata.Destination.Fqdn = d
				metadata.Destination.Addr = netip.Addr{}
			} else {
				return fmt.Errorf("fakeip RestoreToDomain failed: %v", err)
			}
		}
	}
	return r.Router.RouteConnection(ctx, conn, metadata)
}

func (r *RouterWithFakeIP) RoutePacketConnection(ctx context.Context, conn network.PacketConn, metadata adapter.InboundContext) error {
	h := hooks.Ctx(ctx)
	if h != nil && h.FakeEngine != nil {
		ip := metadata.Destination.IPAddr().IP
		if h.FakeEngine.IsIPInPool(ip) {
			return fmt.Errorf("attempt to use fakeip udp")
		}
	}
	return r.Router.RoutePacketConnection(ctx, conn, metadata)
}
