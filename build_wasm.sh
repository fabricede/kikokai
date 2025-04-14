#!/bin/bash

echo "Building WebAssembly module..."
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" ./static/wasm_exec.js

# Clean any previous build
rm -f ./static/main.wasm

# Build the WebAssembly binary directly to the static directory
GOOS=js GOARCH=wasm go build -o ./static/main.wasm ./src/wasm/wasm.go

# Verify the build
if [ -f "./static/main.wasm" ]; then
    echo "WebAssembly build successful!"
    ls -la ./static/main.wasm
else
    echo "WebAssembly build failed!"
    exit 1
fi

echo "Building Go application..."
go build -o ./bin/server.exe ./src/ 

echo "Go application build complete!"