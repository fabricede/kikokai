package shared

import "kikokai/src/model"

// Cube is the global cube instance to be shared across servers
var Cube *model.Cube

func init() {
	// Initialize a new cube
	Cube = model.NewCube()
}

// ResetCube resets the cube to its initial state
func ResetCube() {
	Cube = model.NewCube()
}
