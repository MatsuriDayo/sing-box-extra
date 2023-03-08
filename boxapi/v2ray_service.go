package boxapi

import (
	"context"
	"net"

	"github.com/sagernet/sing-box/adapter"
	"github.com/sagernet/sing-box/experimental/v2rayapi"
	"github.com/sagernet/sing-box/option"
	"github.com/sagernet/sing/common/network"
)

type SbV2rayServer struct {
	ss *SbV2rayStatsService
}

func NewSbV2rayServer(options option.V2RayStatsServiceOptions) *SbV2rayServer {
	return &SbV2rayServer{
		ss: &SbV2rayStatsService{v2rayapi.NewStatsService(options)},
	}
}

func (s *SbV2rayServer) Start() error                            { return nil }
func (s *SbV2rayServer) Close() error                            { return nil }
func (s *SbV2rayServer) StatsService() adapter.V2RayStatsService { return s.ss }

func (s *SbV2rayServer) QueryStats(name string) int64 {
	req := &v2rayapi.GetStatsRequest{
		Name:   name,
		Reset_: true,
	}
	resp, err := s.ss.GetStats(context.TODO(), req)
	if err == nil {
		return resp.Stat.Value
	}
	return 0
}

type SbV2rayStatsService struct {
	*v2rayapi.StatsService
}

func (s *SbV2rayStatsService) RoutedConnection(inbound string, outbound string, user string, conn net.Conn) net.Conn {
	// TODO track
	return s.StatsService.RoutedConnection(inbound, outbound, user, conn)
}

func (s *SbV2rayStatsService) RoutedPacketConnection(inbound string, outbound string, user string, conn network.PacketConn) network.PacketConn {
	return s.StatsService.RoutedPacketConnection(inbound, outbound, user, conn)
}
