# LAN RootCA

```sh
go build -ldflags "-s -w" -o ca.exe main.go
```

## Windows Install

```sh
certutil -addstore -f "ROOT" root-ca.crt
```

## RELEASE

```sh
git tag v0.0.1
git push origin --tags
go list -m github.com/n0z0/crootca@v0.0.1
```
