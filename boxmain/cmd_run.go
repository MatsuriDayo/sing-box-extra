package boxmain

import (
	"context"
	"io"
	"os"
	"os/signal"
	runtimeDebug "runtime/debug"
	"syscall"

	"github.com/matsuridayo/libneko/protect_server"
	"github.com/matsuridayo/sing-box-extra/boxbox"

	"github.com/sagernet/sing-box/log"
	"github.com/sagernet/sing-box/option"
	E "github.com/sagernet/sing/common/exceptions"

	"github.com/spf13/cobra"
)

var commandRun = &cobra.Command{
	Use:   "run",
	Short: "Run service",
	Run: func(cmd *cobra.Command, args []string) {
		if protectListenPath != "" {
			// for v2ray now
			go protect_server.ServeProtect(protectListenPath, true, protectFwMark, nil)
		}

		err := run()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	mainCommand.AddCommand(commandRun)
}

func readConfig(nekoConfigContent []byte) (option.Options, error) {
	var (
		configContent []byte
		err           error
	)
	if nekoConfigContent == nil {
		if configPath == "stdin" {
			configContent, err = io.ReadAll(os.Stdin)
		} else {
			configContent, err = os.ReadFile(configPath)
		}
	} else {
		configContent = nekoConfigContent
	}
	if err != nil {
		return option.Options{}, E.Cause(err, "read config")
	}
	var options option.Options
	err = options.UnmarshalJSON(configContent)
	if err != nil {
		return option.Options{}, E.Cause(err, "decode config")
	}
	return options, nil
}

func Create(nekoConfigContent []byte, forceDisableColor bool) (*boxbox.Box, context.CancelFunc, error) {
	options, err := readConfig(nekoConfigContent)
	if err != nil {
		return nil, nil, err
	}
	if disableColor || forceDisableColor {
		if options.Log == nil {
			options.Log = &option.LogOptions{}
		}
		options.Log.DisableColor = true
	}
	ctx, cancel := context.WithCancel(context.Background())
	instance, err := boxbox.New(ctx, options, nil)
	if err != nil {
		cancel()
		return nil, nil, E.Cause(err, "create service")
	}
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	defer func() {
		signal.Stop(osSignals)
		close(osSignals)
	}()

	go func() {
		_, loaded := <-osSignals
		if loaded {
			cancel()
		}
	}()
	err = instance.Start()
	if err != nil {
		cancel()
		return nil, nil, E.Cause(err, "start service")
	}
	return instance, cancel, nil
}

func run() error {
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	defer signal.Stop(osSignals)
	for {
		instance, cancel, err := Create(nil, false)
		if err != nil {
			return err
		}
		runtimeDebug.FreeOSMemory()
		for {
			osSignal := <-osSignals
			if osSignal == syscall.SIGHUP {
				err = check()
				if err != nil {
					log.Error(E.Cause(err, "reload service"))
					continue
				}
			}
			cancel()
			instance.Close()
			if osSignal != syscall.SIGHUP {
				return nil
			}
			break
		}
	}
}
