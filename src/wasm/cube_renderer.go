//go:build js && wasm
// +build js,wasm

package main

import (
	"encoding/json"
	"kikokai/src/model"
	"syscall/js"
)

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
		material := three.Get("MeshStandardMaterial").New(map[string]any{
			"color": 0x111111, // Dark gray for non-visible sides
		})
		materials.SetIndex(i, material)
	}

	// Right face (x = 1)
	if x == 1 {
		colorIdx := cube.State[model.Right].Stickers[y+1][z+1].Color
		color := colorMap[colorIdx]
		println("Right face color at", x, y, z, ":", colorIdx, "mapped to hex:", color)
		materials.Index(0).Get("color").Call("setHex", color)
	}

	// Left face (x = -1)
	if x == -1 {
		colorIdx := cube.State[model.Left].Stickers[y+1][1-z].Color
		color := colorMap[colorIdx]
		println("Left face color at", x, y, z, ":", colorIdx, "mapped to hex:", color)
		materials.Index(1).Get("color").Call("setHex", color)
	}

	// Top face (y = 1)
	if y == 1 {
		colorIdx := cube.State[model.Up].Stickers[1-z][x+1].Color
		color := colorMap[colorIdx]
		println("Top face color at", x, y, z, ":", colorIdx, "mapped to hex:", color)
		materials.Index(2).Get("color").Call("setHex", color)
	}

	// Bottom face (y = -1)
	if y == -1 {
		colorIdx := cube.State[model.Down].Stickers[z+1][x+1].Color
		color := colorMap[colorIdx]
		println("Bottom face color at", x, y, z, ":", colorIdx, "mapped to hex:", color)
		materials.Index(3).Get("color").Call("setHex", color)
	}

	// Front face (z = 1)
	if z == 1 {
		colorIdx := cube.State[model.Front].Stickers[y+1][x+1].Color
		color := colorMap[colorIdx]
		println("Front face color at", x, y, z, ":", colorIdx, "mapped to hex:", color)
		materials.Index(4).Get("color").Call("setHex", color)
	}

	// Back face (z = -1)
	if z == -1 {
		colorIdx := cube.State[model.Back].Stickers[y+1][1-x].Color
		color := colorMap[colorIdx]
		println("Back face color at", x, y, z, ":", colorIdx, "mapped to hex:", color)
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
func getState(this js.Value, args []js.Value) any {
	stateJSON, _ := json.Marshal(cube.State)
	return js.ValueOf(string(stateJSON))
}

// Update the cube state from a JSON string
func updateCubeFromState(this js.Value, args []js.Value) any {
	if isAnimating {
		return js.ValueOf("Animation in progress")
	}

	if len(args) < 1 {
		return js.ValueOf("Error: Missing state parameter")
	}

	stateJSON := args[0].String()

	// Parse the JSON string into cube state
	var state [6]model.Face
	err := json.Unmarshal([]byte(stateJSON), &state)
	if err != nil {
		println("Error parsing cube state:", err.Error())
		return js.ValueOf("Error: Invalid state format")
	}

	// Update the cube state
	cube.State = state

	// Rebuild the cube visualization
	createCube()

	return js.ValueOf("Cube state updated")
}

// Reset the cube
func resetCube(this js.Value, args []js.Value) any {
	if isAnimating {
		return js.ValueOf("Animation in progress")
	}

	cube = model.NewCube()
	createCube()
	return js.ValueOf("Cube reset")
}

// Scramble the cube
func scrambleCube(this js.Value, args []js.Value) any {
	if isAnimating {
		return js.ValueOf("Animation in progress")
	}

	// Scramble the cube with a standard number of random moves
	cube.Scramble(20) // Scramble with 20 random moves
	createCube()
	return js.ValueOf("Cube scrambled")
}
