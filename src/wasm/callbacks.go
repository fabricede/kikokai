//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"
)

// Register JavaScript callbacks with proper debug output
func registerCallbacks() {
	// Create functions that will persist (avoid garbage collection)
	initThreeSceneFunc := js.FuncOf(initThreeScene)
	getStateFunc := js.FuncOf(getState)
	rotateFaceFunc := js.FuncOf(rotateFace)
	resetCubeFunc := js.FuncOf(resetCube)
	scrambleCubeFunc := js.FuncOf(scrambleCube)
	addCoordinateAxesFunc := js.FuncOf(addCoordinateAxes)

	// Register functions in the global namespace
	js.Global().Set("wasmInitThreeScene", initThreeSceneFunc)
	js.Global().Set("wasmGetState", getStateFunc)
	js.Global().Set("wasmRotateFace", rotateFaceFunc)
	js.Global().Set("wasmResetCube", resetCubeFunc)
	js.Global().Set("wasmScrambleCube", scrambleCubeFunc)
	js.Global().Set("wasmAddCoordinateAxes", addCoordinateAxesFunc)

	// Add a debug function to verify registration
	debugFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		return js.ValueOf("WASM functions registered successfully")
	})
	js.Global().Set("wasmDebug", debugFunc)

	// Store functions to prevent garbage collection
	// This is crucial - functions will be garbage collected if not stored
	funcs = append(funcs, initThreeSceneFunc, getStateFunc, rotateFaceFunc,
		resetCubeFunc, scrambleCubeFunc, addCoordinateAxesFunc,
		debugFunc)

	// Print to console that functions are registered
	println("WASM functions registered: wasmInitThreeScene, wasmGetState, wasmRotateFace, wasmResetCube, wasmScrambleCube, wasmAddCoordinateAxes")
}
