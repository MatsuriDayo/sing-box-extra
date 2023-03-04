package main

import (
	"fmt"
	_ "unsafe"

	"github.com/matsuridayo/sing-box-extra/boxmain"
	_ "github.com/matsuridayo/sing-box-extra/distro/all"

	"github.com/sagernet/sing-box/constant"
)

func main() {
	fmt.Println("sing-box-extra:", constant.Version)
	fmt.Println()

	// sing-box
	boxmain.Main()
}
