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

	// Convert from -1,0,1 coordinates to 0,1,2 array indices
	cubeX := x + 1
	cubeY := y + 1
	cubeZ := z + 1

	// Debug info
	println("Creating cubie at visual coordinates:", x, y, z, "array indices:", cubeX, cubeY, cubeZ)

	// Check if this cubie exists and has colors
	cubie := cube.Cubies[cubeX][cubeY][cubeZ]
	if cubie == nil {
		println("Warning: No cubie at position", x, y, z, "(converted to", cubeX, cubeY, cubeZ, ")")
		return
	}

	// Log all colors for this cubie for debugging
	println("Colors for cubie at", cubeX, cubeY, cubeZ, ":")
	for face, color := range cubie.Colors {
		println("  Face", face, "has color", color)
	}

	// Right face (x = 1)
	if x == 1 && cubie.Colors != nil {
		if color, ok := cubie.Colors[model.Right]; ok {
			hexColor := colorMap[color]
			println("Setting RIGHT face color at", x, y, z, "to", color, "(hex:", hexColor, ")")
			materials.Index(0).Get("color").Call("setHex", hexColor)
		} else {
			println("No RIGHT face color found for cubie at", x, y, z)
		}
	}

	// Left face (x = -1)
	if x == -1 && cubie.Colors != nil {
		if color, ok := cubie.Colors[model.Left]; ok {
			hexColor := colorMap[color]
			println("Setting LEFT face color at", x, y, z, "to", color, "(hex:", hexColor, ")")
			materials.Index(1).Get("color").Call("setHex", hexColor)
		} else {
			println("No LEFT face color found for cubie at", x, y, z)
		}
	}

	// Top face (y = 1)
	if y == 1 && cubie.Colors != nil {
		if color, ok := cubie.Colors[model.Up]; ok {
			hexColor := colorMap[color]
			println("Setting UP face color at", x, y, z, "to", color, "(hex:", hexColor, ")")
			materials.Index(2).Get("color").Call("setHex", hexColor)
		} else {
			println("No UP face color found for cubie at", x, y, z)
		}
	}

	// Bottom face (y = -1)
	if y == -1 && cubie.Colors != nil {
		if color, ok := cubie.Colors[model.Down]; ok {
			hexColor := colorMap[color]
			println("Setting DOWN face color at", x, y, z, "to", color, "(hex:", hexColor, ")")
			materials.Index(3).Get("color").Call("setHex", hexColor)
		} else {
			println("No DOWN face color found for cubie at", x, y, z)
		}
	}

	// Front face (z = 1)
	if z == 1 && cubie.Colors != nil {
		if color, ok := cubie.Colors[model.Back]; ok {
			hexColor := colorMap[color]
			println("Setting FRONT face color at", x, y, z, "to", color, "(hex:", hexColor, ")")
			materials.Index(4).Get("color").Call("setHex", hexColor)
		} else {
			println("No FRONT face color found for cubie at", x, y, z)
		}
	}

	// Back face (z = -1)
	if z == -1 && cubie.Colors != nil {
		if color, ok := cubie.Colors[model.Front]; ok {
			hexColor := colorMap[color]
			println("Setting BACK face color at", x, y, z, "to", color, "(hex:", hexColor, ")")
			materials.Index(5).Get("color").Call("setHex", hexColor)
		} else {
			println("No BACK face color found for cubie at", x, y, z)
		}
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
	stateJSON, _ := json.Marshal(cube.Cubies)
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
	var cubies [3][3][3]*model.Cubie
	err := json.Unmarshal([]byte(stateJSON), &cubies)
	if err != nil {
		println("Error parsing cube state:", err.Error())
		return js.ValueOf("Error: Invalid state format")
	}

	// Update the cube state
	cube.Cubies = cubies

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
