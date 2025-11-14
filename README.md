# LAN RootCA

```sh
go build -ldflags "-s -w" -o ca.exe main.go
```

## Windows Install

```sh
certutil -addstore -f "ROOT" root-ca.crt
```
