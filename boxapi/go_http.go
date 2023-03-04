package boxapi

import (
	"context"
	"net"
	"net/http"
	"reflect"
	"time"
	"unsafe"

	box "github.com/sagernet/sing-box"
	"github.com/sagernet/sing-box/adapter"
	"github.com/sagernet/sing-box/common/dialer"
	"github.com/sagernet/sing/common/metadata"
	"github.com/sagernet/sing/common/network"
)

func GetProxyHttpClient(box *box.Box) *http.Client {
	var d network.Dialer

	if box != nil {
		router_ := reflect.Indirect(reflect.ValueOf(box)).FieldByName("router")
		router_ = reflect.NewAt(router_.Type(), unsafe.Pointer(router_.UnsafeAddr())).Elem()
		if router, ok := router_.Interface().(adapter.Router); ok {
			d = dialer.NewRouter(router)
		}
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
