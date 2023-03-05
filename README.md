# sonnet
go-jsonnet running in the browser.

Inspired by [gjson-play](https://github.com/tidwall/gjson-play).

# Build

```sh
GOOS=js GOARCH=wasm go build -o sonnet.wasm cmd/sonnet/main.go && \
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

# Test

```sh
go run cmd/server/main.go
```
