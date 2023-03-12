package adapter

import (
	"net"
)

type FakeDnsEngineI interface {
	LookupDomain(domain string) net.IP
	IsIPInPool(ip net.IP) bool
	RestoreToDomain(ip net.IP) (string, error)
}
