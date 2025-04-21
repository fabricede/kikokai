//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"
)

// Initialize THREE.js scene from Go
func initThreeScene(this js.Value, args []js.Value) any {
	// Set up constructor references
	setupThreeReferences()

	// Create scene
	scene = three.Get("Scene").New()
	scene.Call("add", setupLighting())

	// Create camera
	camera = three.Get("PerspectiveCamera").New(75, 600/500, 0.1, 1000)
	camera.Get("position").Set("x", 4)
	camera.Get("position").Set("y", 4)
	camera.Get("position").Set("z", 4)
	camera.Call("lookAt", 0, 0, 0)

	// Create renderer
	renderer = three.Get("WebGLRenderer").New(map[string]any{
		"antialias": true,
	})
	renderer.Call("setSize", 600, 500)

	// Attach to DOM element
	document := js.Global().Get("document")
	container := document.Call("getElementById", "cubeCanvas")
	container.Call("appendChild", renderer.Get("domElement"))

	// Set up controls
	controls := three.Get("OrbitControls").New(camera, renderer.Get("domElement"))
	controls.Set("enableDamping", true)
	controls.Set("dampingFactor", 0.05)

	// Set background color
	scene.Set("background", threeColor.New(0xf0f0f0))

	// Create cube group
	cubeGroup = group.New()
	scene.Call("add", cubeGroup)

	// Create initial cube
	createCube()

	// Set up animation loop
	var animationCallback js.Func
	animationCallback = js.FuncOf(func(this js.Value, args []js.Value) any {
		js.Global().Call("requestAnimationFrame", animationCallback)
		controls.Call("update")
		renderer.Call("render", scene, camera)
		return nil
	})
	js.Global().Call("requestAnimationFrame", animationCallback)

	return nil
}

// Set up lighting for the scene
func setupLighting() js.Value {
	lightGroup := group.New()

	// Add ambient light
	ambientLight := three.Get("AmbientLight").New(0xffffff, 0.6)
	lightGroup.Call("add", ambientLight)

	// Add directional light
	directionalLight := three.Get("DirectionalLight").New(0xffffff, 0.6)
	directionalLight.Get("position").Set("x", 10)
	directionalLight.Get("position").Set("y", 20)
	directionalLight.Get("position").Set("z", 30)
	lightGroup.Call("add", directionalLight)

	return lightGroup
}
