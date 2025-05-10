//go:build js && wasm
// +build js,wasm

package main

import (
	"encoding/json"
	"kikokai/src/model"
	"math"
	"syscall/js"
)

// Rotate a face of the cube
func rotateFace(this js.Value, args []js.Value) any {
	// Enhanced parameter validation
	if isAnimating {
		println("Animation already in progress, ignoring rotation request")
		return js.ValueOf("Animation in progress")
	}

	if len(args) < 2 {
		println("Error: Not enough arguments to rotateFace, expected 2, got", len(args))
		return js.ValueOf("Invalid arguments: expected face index and clockwise boolean")
	}

	// Verify face parameter exists and is valid
	if args[0].IsUndefined() || args[0].IsNull() {
		println("Error: Face parameter is undefined or null")
		return js.ValueOf("Invalid face parameter")
	}

	face := model.FaceIndex(args[0].Int())

	if face < 0 || face > 5 {
		println("Error: Invalid face index:", int(face))
		return js.ValueOf("Invalid face index")
	}

	// Verify clockwise parameter exists and is a boolean
	var clockwise model.TurningDirection
	if args[1].IsUndefined() || args[1].IsNull() {
		println("Warning: Clockwise parameter is undefined or null, defaulting to clockwise")
		clockwise = model.Clockwise
	} else {
		// Explicitly log the raw value first
		println("Raw clockwise value received:", args[1].String(), "as bool:", args[1].Bool())

		// Use Bool() only when we know the value is defined
		if args[1].Bool() {
			clockwise = model.Clockwise
			println("Setting direction to CLOCKWISE")
		} else {
			clockwise = model.CounterClockwise
			println("Setting direction to COUNTER-CLOCKWISE")
		}
	}

	println("Starting rotation of face", int(face), "with clockwise value:", clockwise == model.Clockwise)
	isAnimating = true

	// Start animation
	go animateFaceRotation(face, clockwise)

	return js.ValueOf("Animation started")
}

// Animate the rotation of a face
func animateFaceRotation(face model.FaceIndex, clockwise model.TurningDirection) {
	// Log start of animation
	dirStr := "clockwise"
	dirValue := 1
	if clockwise == model.CounterClockwise {
		dirStr = "counter-clockwise"
		dirValue = 0
	}
	println("Starting rotation of face", int(face), dirStr, "direction value:", dirValue)

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
	rotationAxis := model.FaceToCoordinate(face)

	// Define rotation parameters based on direction
	var rotationAngle float64
	var targetRotation float64

	// Always use positive values and adjust sign based on direction
	if clockwise == model.Clockwise {
		rotationAngle = -0.1          // Negative for clockwise
		targetRotation = -math.Pi / 2 // -90 degrees
	} else {
		rotationAngle = 0.1          // Positive for counter-clockwise
		targetRotation = math.Pi / 2 // +90 degrees
	}

	totalRotation := float64(0)

	// Set up animation callback
	var animateFrame js.Func
	animateFrame = js.FuncOf(func(this js.Value, args []js.Value) any {
		// Check if we've reached the target rotation amount
		if math.Abs(totalRotation) < math.Abs(targetRotation) {
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
			stateBeforeJSON, _ := json.Marshal(cube.Cubies)
			println("Cube state before update:", string(stateBeforeJSON))

			// Update the model
			cube.RotateAxis(rotationAxis, clockwise)

			// Log cube state after update
			stateAfterJSON, _ := json.Marshal(cube.Cubies)
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
func getRotationAxis(face model.CubeCoordinate) js.Value {
	return vector3.New(face.X, face.Y, face.Z)
}
