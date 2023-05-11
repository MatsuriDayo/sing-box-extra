package boxbox

import "github.com/sagernet/sing-box/experimental/clashapi"

func (s *Box) closeClashApi() error {
	// Close() may timeout, close early to prevent listen port
	if c, ok := s.router.ClashServer().(*clashapi.Server); ok {
		return c.StopServer()
	}
	return nil
}
