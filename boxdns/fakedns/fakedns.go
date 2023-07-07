package fakedns

import (
	"context"
	"fmt"
	"net"
	"net/netip"
	"strings"

	"github.com/matsuridayo/sing-box-extra/boxdns/fakedns/fakeip"
	"github.com/matsuridayo/sing-box-extra/hooks"
	"github.com/miekg/dns"
	D "github.com/sagernet/sing-dns"
	"github.com/sagernet/sing/common/logger"
	N "github.com/sagernet/sing/common/network"
)

func init() {
	D.RegisterTransport([]string{"fakedns"}, createFakeDNSTransport)
}

func createFakeDNSTransport(name string, ctx context.Context, logger logger.ContextLogger, dialer N.Dialer, link string) (D.Transport, error) {
	link = strings.TrimPrefix(link, "fakedns://")
	_, ipnet, err := net.ParseCIDR(link)
	if err != nil {
		return nil, fmt.Errorf("parse cidr: %v", err)
	}
	//
	pool, err := fakeip.New(fakeip.Options{
		IPNet: ipnet,
		Size:  1000,
	})
	if err != nil {
		return nil, fmt.Errorf("create fakeip pool: %v", err)
	}
	//
	fe := &fakednsEngine{pool}
	t := &fakednsTransport{name, fe}
	if c := hooks.Ctx(ctx); c != nil {
		if c.FakeEngine != nil {
			return nil, fmt.Errorf("can only contain one fakedns server")
		}
		c.FakeEngine = fe // No router at this time
	}
	return t, nil
}

type fakednsTransport struct {
	name string
	fe   *fakednsEngine
}

func (t *fakednsTransport) Name() string { return t.name }
func (t *fakednsTransport) Start() error { return nil }
func (t *fakednsTransport) Close() error { return nil }
func (t *fakednsTransport) Raw() bool    { return false }
func (t *fakednsTransport) Exchange(ctx context.Context, message *dns.Msg) (*dns.Msg, error) {
	return nil, D.ErrNoRawSupport
}
func (t *fakednsTransport) Lookup(ctx context.Context, domain string, strategy D.DomainStrategy) (ips []netip.Addr, err error) {
	ip := t.fe.LookupDomain(domain)
	nip, _ := netip.AddrFromSlice(ip[:])
	return []netip.Addr{nip}, nil
}
