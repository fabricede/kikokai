//go:build js && wasm
// +build js,wasm

package main

import (
	"kikokai/src/model"
	"syscall/js"
)

// Global variables
var (
	cube        *model.Cube
	scene       js.Value
	renderer    js.Value
	camera      js.Value
	cubeGroup   js.Value
	isAnimating bool

	// Constants
	cubeSize float64 = 1
	gap      float64 = 0.05

	// Color mapping from Color enum to hex values
	colorMap = map[model.Color]uint32{
		model.White:  0xFFFFFF,
		model.Yellow: 0xFFFF00,
		model.Blue:   0x0000FF,
		model.Green:  0x00FF00,
		model.Red:    0xFF0000,
		model.Orange: 0xFFA500,
	}

	// THREE.js references
	three      js.Value
	vector3    js.Value
	threeColor js.Value
	box        js.Value
	mesh       js.Value
	group      js.Value

	// Functions to prevent garbage collection
	funcs []js.Func
)

func init() {
	// Initialize a new cube
	cube = model.NewCube()
}

// Set up Three.js references
func setupThreeReferences() {
	// Get THREE global
	three = js.Global().Get("THREE")

	// Store constructor references
	vector3 = three.Get("Vector3")
	threeColor = three.Get("Color")
	group = three.Get("Group")
	box = three.Get("BoxGeometry")
	mesh = three.Get("Mesh")
}
