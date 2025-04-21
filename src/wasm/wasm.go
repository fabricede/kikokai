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
	println("Creating cube piece at position:", x, y, z)

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
		color := colorMap[cube.State[5].Stickers[y+1][z+1].Color]
		println("Right face color at", x, y, z, ":", cube.State[5].Stickers[y+1][z+1].GetName(), "mapped to hex:", color)
		materials.Index(0).Get("color").Call("setHex", color)
	}

	// Left face (x = -1)
	if x == -1 {
		color := colorMap[cube.State[4].Stickers[y+1][1-z].Color]
		println("Left face color at", x, y, z, ":", cube.State[4].Stickers[y+1][1-z].GetName(), "mapped to hex:", color)
		materials.Index(1).Get("color").Call("setHex", color)
	}

	// Top face (y = 1)
	if y == 1 {
		color := colorMap[cube.State[2].Stickers[1-z][x+1].Color]
		println("Top face color at", x, y, z, ":", cube.State[2].Stickers[1-z][x+1].GetName(), "mapped to hex:", color)
		materials.Index(2).Get("color").Call("setHex", color)
	}

	// Bottom face (y = -1)
	if y == -1 {
		color := colorMap[cube.State[3].Stickers[z+1][x+1].Color]
		println("Bottom face color at", x, y, z, ":", cube.State[3].Stickers[z+1][x+1].GetName(), "mapped to hex:", color)
		materials.Index(3).Get("color").Call("setHex", color)
	}

	// Front face (z = 1)
	if z == 1 {
		color := colorMap[cube.State[0].Stickers[y+1][x+1].Color]
		println("Front face color at", x, y, z, ":", cube.State[0].Stickers[y+1][x+1].GetName(), "mapped to hex:", color)
		materials.Index(4).Get("color").Call("setHex", color)
	}

	// Back face (z = -1)
	if z == -1 {
		color := colorMap[cube.State[1].Stickers[y+1][1-x].Color]
		println("Back face color at", x, y, z, ":", cube.State[1].Stickers[y+1][1-x].GetName(), "mapped to hex:", color)
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

	face := model.FaceIndex(args[0].Int())
	// Convert boolean to the proper Direction type constant
	var clockwise model.TurningDirection
	if args[1].Bool() {
		clockwise = model.Clockwise
	} else {
		clockwise = model.CounterClockwise
	}

	if face < 0 || face > 5 {
		return js.ValueOf("Invalid face index")
	}

	isAnimating = true

	// Start animation
	go animateFaceRotation(face, clockwise)

	return js.ValueOf("Animation started")
}

// Animate the rotation of a face
func animateFaceRotation(face model.FaceIndex, clockwise model.TurningDirection) {
	// Log start of animation
	dirStr := "clockwise"
	if clockwise == model.CounterClockwise {
		dirStr = "counter-clockwise"
	}
	println("Starting rotation of face", int(face), dirStr)

	// Create a rotation group
	rotationGroup := group.New()
	scene.Call("add", rotationGroup)

	// Find cubes to rotate - collect all candidates first before modifying the group
	var cubesToRotate []js.Value
	children := cubeGroup.Get("children")
	length := children.Length()

	// First, identify all cubes that should rotate with this face
	for i := 0; i < length; i++ {
		child := children.Index(i)
		// Check if the cube's userData exists before accessing it
		if !child.IsUndefined() && !child.IsNull() && !child.Get("userData").IsUndefined() {
			if shouldRotateWithFace(child, face) {
				cubesToRotate = append(cubesToRotate, child)
			}
		}
	}

	println("Found", len(cubesToRotate), "cubes to rotate with face", int(face))

	// Then remove them from the cube group and add to rotation group
	for _, cubeMesh := range cubesToRotate {
		cubeGroup.Call("remove", cubeMesh)
		rotationGroup.Call("add", cubeMesh)
	}

	// Get rotation axis
	rotationAxis := getRotationAxis(face)

	// Animate rotation - setting rotation angle based on direction
	rotationAngle := float64(-0.1) // Default for clockwise
	if clockwise == model.CounterClockwise {
		rotationAngle = float64(0.1) // For counter-clockwise
	}

	totalRotation := float64(0)
	targetRotation := -3.14159 / 2 // Default for clockwise

	if clockwise == model.CounterClockwise {
		targetRotation = 3.14159 / 2 // For counter-clockwise
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
			println("Animation complete for face", int(face), "- updating cube model")

			// Make sure to iterate in reverse to avoid index issues when removing children
			for i := rotationGroup.Get("children").Get("length").Int() - 1; i >= 0; i-- {
				child := rotationGroup.Get("children").Index(i)
				if !child.IsUndefined() && !child.IsNull() {
					rotationGroup.Call("remove", child)
				}
			}

			scene.Call("remove", rotationGroup)

			// Log cube state before update
			stateBeforeJSON, _ := json.Marshal(cube.State)
			println("Cube state before update:", string(stateBeforeJSON))

			// Update the model
			cube.RotateFace(face, clockwise)

			// Log cube state after update
			stateAfterJSON, _ := json.Marshal(cube.State)
			println("Cube state after update:", string(stateAfterJSON))

			// Recreate the cube with new state
			createCube()

			// Release the animateFrame function from memory when animation is complete
			animateFrame.Release()

			isAnimating = false
			println("Animation and model update completed for face", int(face))
		}
		return nil
	})

	// Store function in funcs to prevent garbage collection during animation
	funcs = append(funcs, animateFrame)

	js.Global().Call("requestAnimationFrame", animateFrame)
}

// Determine if a cube should rotate with the face
func shouldRotateWithFace(cube js.Value, face model.FaceIndex) bool {
	// Verify that cube and userData exist
	if cube.IsUndefined() || cube.IsNull() {
		println("Warning: Undefined or null cube in shouldRotateWithFace")
		return false
	}

	userData := cube.Get("userData")
	if userData.IsUndefined() || userData.IsNull() {
		println("Warning: Undefined or null userData in shouldRotateWithFace")
		return false
	}

	// Get position values, checking each value exists
	if userData.Get("posX").IsUndefined() || userData.Get("posY").IsUndefined() || userData.Get("posZ").IsUndefined() {
		println("Warning: Position values missing in userData")
		return false
	}

	posX := userData.Get("posX").Int()
	posY := userData.Get("posY").Int()
	posZ := userData.Get("posZ").Int()

	var shouldRotate bool

	switch face {
	case model.Front:
		shouldRotate = posZ == 1
	case model.Back:
		shouldRotate = posZ == -1
	case model.Up:
		shouldRotate = posY == 1
	case model.Down:
		shouldRotate = posY == -1
	case model.Left:
		shouldRotate = posX == -1
	case model.Right:
		shouldRotate = posX == 1
	default:
		shouldRotate = false
	}

	// Log selection decision with cube position information
	if shouldRotate {
		println("Selected cube for rotation at position:", posX, posY, posZ, "for face", int(face))
	}

	return shouldRotate
}

// Get the axis for rotation based on the face
func getRotationAxis(face model.FaceIndex) js.Value {
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

// Scramble the cube
func scrambleCube(this js.Value, args []js.Value) interface{} {
	if isAnimating {
		return js.ValueOf("Animation in progress")
	}

	// Scramble the cube with a standard number of random moves
	cube.Scramble(20) // Scramble with 20 random moves
	createCube()
	return js.ValueOf("Cube scrambled")
}

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
	context.Set("fillStyle", "rgba(255, 255, 255, 0.8)")
	context.Call("fillRect", 0, 0, canvas.Get("width").Int(), canvas.Get("height").Int())

	// Draw the text
	context.Set("font", "24px Arial")
	context.Set("fillStyle", color)
	context.Set("textAlign", "center")
	context.Set("textBaseline", "middle")
	context.Call("fillText", text, canvas.Get("width").Int()/2, canvas.Get("height").Int()/2)

	// Create a texture from the canvas
	texture := three.Get("CanvasTexture").New(canvas)
	material := three.Get("SpriteMaterial").New(map[string]interface{}{
		"map": texture,
	})

	// Create a sprite with the texture
	sprite := three.Get("Sprite").New(material)

	// Set the sprite position
	sprite.Get("position").Set("x", x)
	sprite.Get("position").Set("y", y)
	sprite.Get("position").Set("z", z)

	// Set the sprite scale
	sprite.Get("scale").Set("x", 0.5)
	sprite.Get("scale").Set("y", 0.25)
	sprite.Get("scale").Set("z", 1)

	// Add the sprite to the scene
	scene.Call("add", sprite)
}

// Register JavaScript callbacks with proper debug output
func registerCallbacks() {
	// Create functions that will persist (avoid garbage collection)
	initThreeSceneFunc := js.FuncOf(initThreeScene)
	getStateFunc := js.FuncOf(getState)
	rotateFaceFunc := js.FuncOf(rotateFace)
	resetCubeFunc := js.FuncOf(resetCube)
	scrambleCubeFunc := js.FuncOf(scrambleCube)

	// Register functions in the global namespace
	js.Global().Set("wasmInitThreeScene", initThreeSceneFunc)
	js.Global().Set("wasmGetState", getStateFunc)
	js.Global().Set("wasmRotateFace", rotateFaceFunc)
	js.Global().Set("wasmResetCube", resetCubeFunc)
	js.Global().Set("wasmScrambleCube", scrambleCubeFunc)

	// Add a debug function to verify registration
	debugFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return js.ValueOf("WASM functions registered successfully")
	})
	js.Global().Set("wasmDebug", debugFunc)

	// Store functions to prevent garbage collection
	// This is crucial - functions will be garbage collected if not stored
	funcs = append(funcs, initThreeSceneFunc, getStateFunc, rotateFaceFunc,
		resetCubeFunc, scrambleCubeFunc,
		debugFunc)

	// Print to console that functions are registered
	println("WASM functions registered: wasmInitThreeScene, wasmGetState, wasmRotateFace, wasmResetCube, wasmScrambleCube, wasmUpdateCubeFromState, wasmAddCoordinateAxes")
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
