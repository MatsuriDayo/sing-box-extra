package boxmain

import (
	"os"
	"time"

	"github.com/sagernet/sing-box/log"
)

func DisableColor() {
	disableColor = true
	log.SetStdLogger(log.NewFactory(log.Formatter{BaseTime: time.Now(), DisableColors: true}, os.Stderr, nil).Logger())
}
