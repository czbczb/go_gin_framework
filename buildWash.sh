rm -rf static/main.wasm
GOOS=js GOARCH=wasm go build -o cmd/wasm/static/main.wasm cmd/wasm/wasm.go