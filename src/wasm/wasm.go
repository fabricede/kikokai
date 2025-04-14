//go:build js && wasm
// +build js,wasm

package main

import (
	"encoding/json"
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

	// Color mapping
	colorMap = map[string]uint32{
		"white":  0xFFFFFF,
		"yellow": 0xFFFF00,
		"blue":   0x0000FF,
		"green":  0x00FF00,
		"red":    0xFF0000,
		"orange": 0xFFA500,
		"black":  0x000000,
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

// Initialize THREE.js scene from Go
func initThreeScene(this js.Value, args []js.Value) interface{} {
	// Set up constructor references
	setupThreeReferences()

	// Create scene
	scene = three.Get("Scene").New()
	scene.Call("add", setupLighting())

	// Create camera
	camera = three.Get("PerspectiveCamera").New(75, 600/500, 0.1, 1000)
	// Fix position setting - position is an object with x, y, z properties
	camera.Get("position").Set("x", 4)
	camera.Get("position").Set("y", 4)
	camera.Get("position").Set("z", 4)
	camera.Call("lookAt", 0, 0, 0)

	// Create renderer
	renderer = three.Get("WebGLRenderer").New(map[string]interface{}{
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
	animationCallback = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
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

// Create the 3D cube from state
func createCube() {
	// Clear existing cube
	for cubeGroup.Get("children").Get("length").Int() > 0 {
		cubeGroup.Call("remove", cubeGroup.Get("children").Index(0))
	}

	// Create small cubes
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				// Skip center cube
				if x == 0 && y == 0 && z == 0 {
					continue
				}

				createCubePiece(x, y, z)
			}
		}
	}
}

// Create a single cube piece
func createCubePiece(x, y, z int) {
	// Create geometry
	geometry := box.New(cubeSize, cubeSize, cubeSize)

	// Create materials array (right, left, top, bottom, front, back)
	materials := js.Global().Get("Array").New(6)

	for i := 0; i < 6; i++ {
		material := three.Get("MeshStandardMaterial").New(map[string]interface{}{
			"color": 0x000000,
		})
		materials.SetIndex(i, material)
	}

	// Right face (x = 1)
	if x == 1 {
		color := colorMap[cube.State[5][y+1][z+1]]
		materials.Index(0).Get("color").Call("setHex", color)
	}

	// Left face (x = -1)
	if x == -1 {
		color := colorMap[cube.State[4][y+1][1-z]]
		materials.Index(1).Get("color").Call("setHex", color)
	}

	// Top face (y = 1)
	if y == 1 {
		color := colorMap[cube.State[2][1-z][x+1]]
		materials.Index(2).Get("color").Call("setHex", color)
	}

	// Bottom face (y = -1)
	if y == -1 {
		color := colorMap[cube.State[3][z+1][x+1]]
		materials.Index(3).Get("color").Call("setHex", color)
	}

	// Front face (z = 1)
	if z == 1 {
		color := colorMap[cube.State[0][y+1][x+1]]
		materials.Index(4).Get("color").Call("setHex", color)
	}

	// Back face (z = -1)
	if z == -1 {
		color := colorMap[cube.State[1][y+1][1-x]]
		materials.Index(5).Get("color").Call("setHex", color)
	}

	// Create mesh with materials
	cubeMesh := mesh.New(geometry, materials)
	// Set position directly as properties instead of using set() method
	cubeMesh.Get("position").Set("x", float64(x)*(cubeSize+gap))
	cubeMesh.Get("position").Set("y", float64(y)*(cubeSize+gap))
	cubeMesh.Get("position").Set("z", float64(z)*(cubeSize+gap))

	// Add position as custom properties for animation selection
	userData := js.Global().Get("Object").New()
	userData.Set("posX", x)
	userData.Set("posY", y)
	userData.Set("posZ", z)
	cubeMesh.Set("userData", userData)

	// Add to cube group
	cubeGroup.Call("add", cubeMesh)
}

// Get the state of the cube
func getState(this js.Value, args []js.Value) interface{} {
	stateJSON, _ := json.Marshal(cube.State)
	return js.ValueOf(string(stateJSON))
}

// Rotate a face of the cube
func rotateFace(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 || isAnimating {
		return js.ValueOf("Invalid arguments or animation in progress")
	}

	face := model.Face(args[0].Int())
	clockwise := model.Direction(args[1].Bool())

	if face < 0 || face > 5 {
		return js.ValueOf("Invalid face index")
	}

	isAnimating = true

	// Start animation
	go animateFaceRotation(face, clockwise)

	return js.ValueOf("Animation started")
}

// Animate the rotation of a face
func animateFaceRotation(face model.Face, clockwise model.Direction) {
	// Create a rotation group
	rotationGroup := group.New()
	scene.Call("add", rotationGroup)

	// Find cubes to rotate
	var cubesToRotate []js.Value

	children := cubeGroup.Get("children")
	length := children.Length()

	for i := 0; i < length; i++ {
		cube := children.Index(i)
		if shouldRotateWithFace(cube, face) {
			cubesToRotate = append(cubesToRotate, cube)
			cubeGroup.Call("remove", cube)
			rotationGroup.Call("add", cube)
		}
	}

	// Get rotation axis
	rotationAxis := getRotationAxis(face)

	// Animate rotation
	rotationAngle := float64(0.1)
	if !clockwise {
		rotationAngle = -rotationAngle
	}

	totalRotation := float64(0)
	targetRotation := 3.14159 / 2 // 90 degrees in radians

	if !clockwise {
		targetRotation = -targetRotation
	}

	// Set up animation callback
	var animateFrame js.Func
	animateFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if float64(js.Global().Get("Math").Call("abs", totalRotation).Float()) < float64(js.Global().Get("Math").Call("abs", targetRotation).Float()) {
			// Continue animation
			rotationGroup.Call("rotateOnAxis", rotationAxis, rotationAngle)
			totalRotation += rotationAngle
			js.Global().Call("requestAnimationFrame", animateFrame)
		} else {
			// Animation complete - cleanup
			for rotationGroup.Get("children").Get("length").Int() > 0 {
				rotationGroup.Call("remove", rotationGroup.Get("children").Index(0))
			}
			scene.Call("remove", rotationGroup)

			// Update the model
			cube.RotateFace(face, clockwise)

			// Recreate the cube with new state
			createCube()

			isAnimating = false
		}
		return nil
	})

	js.Global().Call("requestAnimationFrame", animateFrame)
}

// Determine if a cube should rotate with the face
func shouldRotateWithFace(cube js.Value, face model.Face) bool {
	userData := cube.Get("userData")
	posX := userData.Get("posX").Int()
	posY := userData.Get("posY").Int()
	posZ := userData.Get("posZ").Int()

	switch face {
	case model.Front:
		return posZ == 1
	case model.Back:
		return posZ == -1
	case model.Up:
		return posY == 1
	case model.Down:
		return posY == -1
	case model.Left:
		return posX == -1
	case model.Right:
		return posX == 1
	default:
		return false
	}
}

// Get the axis for rotation based on the face
func getRotationAxis(face model.Face) js.Value {
	switch face {
	case model.Front:
		return vector3.New(0, 0, 1)
	case model.Back:
		return vector3.New(0, 0, -1)
	case model.Up:
		return vector3.New(0, 1, 0)
	case model.Down:
		return vector3.New(0, -1, 0)
	case model.Left:
		return vector3.New(-1, 0, 0)
	case model.Right:
		return vector3.New(1, 0, 0)
	default:
		return vector3.New(0, 1, 0)
	}
}

// Reset the cube
func resetCube(this js.Value, args []js.Value) interface{} {
	if isAnimating {
		return js.ValueOf("Animation in progress")
	}

	cube = model.NewCube()
	createCube()
	return js.ValueOf("Cube reset")
}

// Register JavaScript callbacks with proper debug output
func registerCallbacks() {
	// Create functions that will persist (avoid garbage collection)
	initThreeSceneFunc := js.FuncOf(initThreeScene)
	getStateFunc := js.FuncOf(getState)
	rotateFaceFunc := js.FuncOf(rotateFace)
	resetCubeFunc := js.FuncOf(resetCube)

	// Register functions in the global namespace
	js.Global().Set("wasmInitThreeScene", initThreeSceneFunc)
	js.Global().Set("wasmGetState", getStateFunc)
	js.Global().Set("wasmRotateFace", rotateFaceFunc)
	js.Global().Set("wasmResetCube", resetCubeFunc)

	// Add a debug function to verify registration
	debugFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return js.ValueOf("WASM functions registered successfully")
	})
	js.Global().Set("wasmDebug", debugFunc)

	// Store functions to prevent garbage collection
	// This is crucial - functions will be garbage collected if not stored
	funcs = append(funcs, initThreeSceneFunc, getStateFunc, rotateFaceFunc, resetCubeFunc, debugFunc)

	// Print to console that functions are registered
	println("WASM functions registered: wasmInitThreeScene, wasmGetState, wasmRotateFace, wasmResetCube")
}

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
