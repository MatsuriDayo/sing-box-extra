package boxroute

import (
	"context"
	"fmt"
	"net"
	"net/netip"

	"github.com/matsuridayo/sing-box-extra/hooks"
	"github.com/sagernet/sing-box/adapter"
	dns "github.com/sagernet/sing-dns"
	"github.com/sagernet/sing/common/network"
)

func init() {
	hooks.RegisterRouterHook(func(r adapter.Router) adapter.Router {
		return &Router{r}
	})
}

type Router struct {
	adapter.Router
}

func (r *Router) RouteConnection(ctx context.Context, conn net.Conn, metadata adapter.InboundContext) error {
	h := hooks.Ctx(ctx)
	if h != nil && h.FakeEngine != nil {
		ip := metadata.Destination.IPAddr().IP
		if h.FakeEngine.IsIPInPool(ip) {
			if d, err := h.FakeEngine.RestoreToDomain(ip); err == nil {
				metadata.User = "fakedns"
				metadata.OriginDestination = metadata.Destination
				metadata.Destination.Fqdn = d
				metadata.Destination.Addr = netip.Addr{}
			} else {
				return fmt.Errorf("fakeip RestoreToDomain failed: %v", err)
			}
		}
	}
	return r.Router.RouteConnection(ctx, conn, metadata)
}

func (r *Router) RoutePacketConnection(ctx context.Context, conn network.PacketConn, metadata adapter.InboundContext) error {
	h := hooks.Ctx(ctx)
	if h != nil && h.FakeEngine != nil {
		ip := metadata.Destination.IPAddr().IP
		if h.FakeEngine.IsIPInPool(ip) {
			//attempt to use fakeip udp
			return nil
		}
	}
	return r.Router.RoutePacketConnection(ctx, conn, metadata)
}

func (r *Router) Lookup(ctx context.Context, domain string, strategy dns.DomainStrategy) ([]netip.Addr, error) {
	ctx, metadata := adapter.AppendContext(ctx)
	if metadata.User == "fakedns" && hooks.TransportNameFromContext(ctx) != "" {
		// avoid dns loopback
		metadata.User = ""
	}
	return r.Router.Lookup(ctx, domain, strategy)
}

func (r *Router) LookupDefault(ctx context.Context, domain string) ([]netip.Addr, error) {
	return r.Lookup(ctx, domain, dns.DomainStrategyAsIS)
}
