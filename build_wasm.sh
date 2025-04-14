#!/bin/bash

echo "Building WebAssembly module..."
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" ./static/wasm_exec.js
# Build the WebAssembly binary
GOOS=js GOARCH=wasm go build -o main.wasm ./src/wasm/wasm.go

echo "WebAssembly build complete!"

echo "Building Go application..."

go build -o ./bin/server.exe ./src/ 

echo "Go application build complete!"