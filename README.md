# sing-box-extra

sing-box and some additional functionality or API

## Components

see distro/all

### boxapi

- v2ray service API (An implementation that replaces the sing-box v2ray option, the caller can set it after obtaining the router)
- golang http.Client API

### boxdns

- underlying DNS for Linux & Windows

### boxbox

Custom Box

Use this instead of `github.com/sagernet/sing-box`

`boxbox.New(Options,PlatformOptions)` is used in `libgojni.so`

### boxmain

Custom CLI tools

`boxmain.Create(jsonContent)` creates `boxbox.Box` instead of `box.Box` (used in `nekobox_core`)
