package boxapi

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/matsuridayo/sing-box-extra/boxbox"

	"github.com/sagernet/sing-box/common/dialer"
	"github.com/sagernet/sing/common/metadata"
	N "github.com/sagernet/sing/common/network"
)

func GetProxyHttpClient(box *boxbox.Box) *http.Client {
	transport := &http.Transport{
		TLSHandshakeTimeout:   time.Second * 3,
		ResponseHeaderTimeout: time.Second * 3,
	}

	if box != nil {
		transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			router := box.Router()
			conn, err := dialer.NewRouter(router).DialContext(ctx, network, metadata.ParseSocksaddr(addr))
			if err != nil {
				return nil, err
			}
			if vs := router.V2RayServer(); vs != nil {
				if ss, ok := vs.StatsService().(*SbStatsService); ok {
					conn = ss.RoutedConnectionInternal("", router.DefaultOutbound(N.NetworkName(network)).Tag(), "", conn, false)
				}
			}
			return conn, nil
		}
	}

	client := &http.Client{
		Transport: transport,
	}

	return client
}
