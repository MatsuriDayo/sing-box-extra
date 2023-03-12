package fakedns

import (
	"fmt"
	"net"

	"github.com/matsuridayo/sing-box-extra/adapter"
	"github.com/matsuridayo/sing-box-extra/boxdns/fakedns/fakeip"
)

var _ adapter.FakeDnsEngineI = (*fakednsEngine)(nil)

type fakednsEngine struct {
	pool *fakeip.Pool
}

func (fe *fakednsEngine) LookupDomain(domain string) net.IP {
	return fe.pool.Lookup(domain)[:]
}

func (fe *fakednsEngine) IsIPInPool(ip net.IP) bool {
	return fe.pool.IPNet().Contains(ip)
}

func (fe *fakednsEngine) RestoreToDomain(ip net.IP) (string, error) {
	domain, ok := fe.pool.LookBack(ip)
	if !ok {
		return "", fmt.Errorf("domain not found in cache pool")
	}
	return domain, nil
}
