package boxapi

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/matsuridayo/sing-box-extra/boxbox"

	"github.com/sagernet/sing-box/adapter"
	"github.com/sagernet/sing/common"
	"github.com/sagernet/sing/common/metadata"
)

func GetProxyHttpClient(box *boxbox.Box) *http.Client {
	transport := &http.Transport{
		TLSHandshakeTimeout:   time.Second * 3,
		ResponseHeaderTimeout: time.Second * 3,
	}

	if box != nil {
		transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			down, up := net.Pipe()
			// stats is in RouteConnection
			go func() {
				box.Router().RouteConnection(ctx, up, adapter.InboundContext{
					Inbound:     "go-http",
					Destination: metadata.ParseSocksaddr(addr),
				})
				common.Close(down, up)
			}()
			return down, nil
		}
	}

	client := &http.Client{
		Transport: transport,
	}

	return client
}
