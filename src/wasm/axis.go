//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"
)

// Create a text label for an axis
func createAxisLabel(text string, x, y, z float64, color string) {
	// Create a canvas for the label
	document := js.Global().Get("document")
	canvas := document.Call("createElement", "canvas")
	context := canvas.Call("getContext", "2d")

	// Set canvas dimensions
	canvas.Set("width", 64)
	canvas.Set("height", 32)

	// Draw the background
	context.Set("fillStyle", "rgba(100, 98, 98, 0.86)")
	context.Call("fillRect", 0, 0, canvas.Get("width").Int(), canvas.Get("height").Int())

	// Draw the text
	context.Set("font", "24px Arial")
	context.Set("fillStyle", color)
	context.Set("textAlign", "center")
	context.Set("textBaseline", "middle")
	context.Call("fillText", text, canvas.Get("width").Int()/2, canvas.Get("height").Int()/2)

	// Create a texture from the canvas
	texture := three.Get("CanvasTexture").New(canvas)
	material := three.Get("SpriteMaterial").New(map[string]any{
		"map": texture,
	})

	// Create a sprite with the texture
	sprite := three.Get("Sprite").New(material)

	// Set the sprite position
	sprite.Get("position").Set("z", z)
	sprite.Get("position").Set("y", y)
	sprite.Get("position").Set("x", x)

	// Set the sprite scale
	sprite.Get("scale").Set("z", 0.5)
	sprite.Get("scale").Set("y", 0.25)
	sprite.Get("scale").Set("x", 1)

	// Add the sprite to the scene
	scene.Call("add", sprite)
}

// Add coordinate axes to the scene
func addCoordinateAxes(this js.Value, args []js.Value) any {
	if scene.IsUndefined() || scene.IsNull() {
		println("Scene not initialized")
		return js.ValueOf("Scene not initialized")
	}

	// Default axis length
	axisLength := 3.0
	if len(args) > 0 && !args[0].IsUndefined() && !args[0].IsNull() {
		axisLength = args[0].Float()
	}

	println("Adding coordinate axes with length:", axisLength)

	// Create axes helper
	axesHelper := three.Get("AxesHelper").New(axisLength)
	scene.Call("add", axesHelper)

	// Create text labels for axes
	createAxisLabel("Z", axisLength+0.2, 0, 0, "#ff0000") // Red for Z axis
	createAxisLabel("Y", 0, axisLength+0.2, 0, "#00ff00") // Green for Y axis
	createAxisLabel("X", 0, 0, axisLength+0.2, "#ffff00") // Yellow for X axis

	return js.ValueOf("Coordinate axes added")
}
