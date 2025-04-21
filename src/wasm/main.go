//go:build js && wasm
// +build js,wasm

package main

func main() {
	// Create channel to keep program alive
	c := make(chan struct{}, 0)

	// Register JavaScript callbacks
	println("Registering callbacks...")
	registerCallbacks()
	println("WebAssembly module initialized")

	// Keep the Go program running
	<-c
}
