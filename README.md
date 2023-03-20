# sing-box-extra

sing-box and some additional functionality or API

## Components

see distro/all

### boxapi

- v2ray service API (An implementation that replaces the sing-box v2ray option, the caller can set it after obtaining the router)
- golang http.Client API

### boxdns

- underlying DNS for Linux & Windows
- FakeDNS server (Note: The `auth_user` will be set to `fakedns` for fakeIP connections. When must query the real IP (e.g. Wireguard), please add rules containing this attribute to correctly finish the necessary DNS query.)

### boxbox

Custom Box

Use this instead of `github.com/sagernet/sing-box`

### boxmain

Custom CLI tools

Creates `boxbox.Box` instead of `box.Box`

### boxroute

Custom Router

It's used if you create `boxbox.Box`
