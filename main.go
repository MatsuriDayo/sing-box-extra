package main

import (
	"fmt"
	_ "unsafe"

	"github.com/matsuridayo/sing-box-extra/boxbox"
	"github.com/matsuridayo/sing-box-extra/boxmain"
	_ "github.com/matsuridayo/sing-box-extra/distro/all"
)

func main() {
	fmt.Println("sing-box-extra:", boxbox.Version)
	fmt.Println()

	// sing-box
	boxmain.Main()
}
