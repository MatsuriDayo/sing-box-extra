# sing-box-extra

sing-box and some additional functionality or API

## Components

see distro/all

### boxbox

Use this instead of Box

### boxapi

- v2ray service API
- golang http.Client API

### boxdns

- underlying DNS for Linux & Windows
- FakeDNS server (Note: The `auth_user` will be set to `fakedns` for fakeIP connections. When must query the real IP (e.g. Wireguard), please add rules containing this attribute to correctly finish the necessary DNS query.)

### boxroute

Custom Router

### boxmain

CLI tools
