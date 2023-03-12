package boxmain

import (
	"context"

	"github.com/matsuridayo/sing-box-extra/boxbox"
	"github.com/sagernet/sing-box/log"

	"github.com/spf13/cobra"
)

var commandCheck = &cobra.Command{
	Use:   "check",
	Short: "Check configuration",
	Run: func(cmd *cobra.Command, args []string) {
		err := check()
		if err != nil {
			log.Fatal(err)
		}
	},
	Args: cobra.NoArgs,
}

func init() {
	mainCommand.AddCommand(commandCheck)
}

func check() error {
	options, err := readConfig(nil)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	_, err = boxbox.New(ctx, options, nil)
	cancel()
	return err
}
