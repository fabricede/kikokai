//go:build js && wasm
// +build js,wasm

package main

import (
	"encoding/json"
	"kikokai/src/model"
	"syscall/js"
)

// Global cube instance
var cube *model.Cube

func init() {
	// Initialize a new cube
	cube = model.NewCube()
}

// WebAssembly exported functions
func getState(this js.Value, args []js.Value) interface{} {
	stateJSON, _ := json.Marshal(cube.State)
	return js.ValueOf(string(stateJSON))
}

func rotateFace(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return js.ValueOf("Invalid arguments")
	}

	face := model.Face(args[0].Int())
	clockwise := model.Direction(args[1].Bool())

	if face < 0 || face > 5 {
		return js.ValueOf("Invalid face index")
	}

	cube.RotateFace(face, clockwise)
	return js.ValueOf("ok")
}

func registerCallbacks() {
	js.Global().Set("getState", js.FuncOf(getState))
	js.Global().Set("rotateFace", js.FuncOf(rotateFace))
}

func main() {
	c := make(chan struct{}, 0)

	// Register JavaScript callbacks
	registerCallbacks()

	println("WebAssembly module initialized")

	// Keep the Go program running
	<-c
}
