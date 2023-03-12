package boxapi

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/matsuridayo/sing-box-extra/boxbox"
	"github.com/sagernet/sing-box/common/dialer"
	"github.com/sagernet/sing/common/metadata"
	"github.com/sagernet/sing/common/network"
)

func GetProxyHttpClient(box *boxbox.Box) *http.Client {
	var d network.Dialer

	if box != nil {
		d = dialer.NewRouter(box.Router())
	}

	dialContext := func(ctx context.Context, network, addr string) (net.Conn, error) {
		return d.DialContext(ctx, network, metadata.ParseSocksaddr(addr))
	}

	transport := &http.Transport{
		TLSHandshakeTimeout:   time.Second * 3,
		ResponseHeaderTimeout: time.Second * 3,
	}

	if d != nil {
		transport.DialContext = dialContext
	}

	client := &http.Client{
		Transport: transport,
	}

	return client
}
