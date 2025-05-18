package model

import (
	"encoding/json"
	"math/rand"
)

// SharedCube is the global cube instance to be shared across servers
var SharedCube *Cube
var initialCube *Cube

func init() {
	// Initialize a new cube
	SharedCube = NewCube()
	initialCube = NewCube() // used to get stickers position
}

// ResetCube resets the cube to its initial state
func ResetCube() {
	SharedCube = NewCube()
}

// -------------------------------------------
// Cube represents the Rubik's Cube as a 3x3x3 array of cubies.
// --------------------------------------------
type Cube struct {
	Cubies [3][3][3]*Cubie
}

// NewCube initializes a solved Rubik's Cube.
func NewCube() *Cube {
	cube := &Cube{}
	for x := range 3 {
		for y := range 3 {
			for z := range 3 {
				cube.Cubies[x][y][z] = NewCubie()
			}
		}
	}
	return cube
}

// RotateAxis rotates a slice of the cube around a specified axis
func (c *Cube) RotateAxis(axis CubeCoordinate, clockwise TurningDirection) {
	// copies the layer of the cube to a matrix
	var layer Layer
	layer.init(c, axis)
	// rotates the layer
	if clockwise {
		layer = layer.rotateClockwise(axis)
	} else {
		layer = layer.rotateCounterClockwise(axis)
	}
	// copies the layer back to the cube
	layer.setLayer(c, axis)
}

// Scramble applies a series of random rotations to the cube
func (c *Cube) Scramble(moves int) {
	// Apply random rotations
	for range moves {
		// Random axis (0-5)
		face := FaceIndex(rand.Intn(6))
		axis := FaceToCoordinate(face)

		// Random direction (true/false for clockwise/counter-clockwise)
		clockwise := TurningDirection(rand.Intn(2) == 1)

		// Rotate the face
		c.RotateAxis(axis, clockwise)
	}
}

// ToReadableJSON returns a human-readable JSON representation of the cube's state.
func (c *Cube) ToReadableJSON() (string, error) {
	cubeState := make([][][]map[string]string, 3)
	for x := 0; x < 3; x++ {
		cubeState[x] = make([][]map[string]string, 3)
		for y := 0; y < 3; y++ {
			cubeState[x][y] = make([]map[string]string, 3)
			for z := 0; z < 3; z++ {
				cubie := c.Cubies[x][y][z]
				if cubie != nil {
					faceColors := make(map[string]string)
					// Only include outer faces
					if x == 0 {
						faceColors["back"] = colorToName(cubie.Colors[Back])
					}
					if x == 2 {
						faceColors["front"] = colorToName(cubie.Colors[Front])
					}
					if y == 0 {
						faceColors["down"] = colorToName(cubie.Colors[Down])
					}
					if y == 2 {
						faceColors["up"] = colorToName(cubie.Colors[Up])
					}
					if z == 0 {
						faceColors["left"] = colorToName(cubie.Colors[Left])
					}
					if z == 2 {
						faceColors["right"] = colorToName(cubie.Colors[Right])
					}
					cubeState[x][y][z] = faceColors
				} else {
					cubeState[x][y][z] = nil
				}
			}
		}
	}
	jsonBytes, err := json.MarshalIndent(cubeState, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// colorToName maps a Color to its string name
func colorToName(color Color) string {
	switch color {
	case White:
		return "white"
	case Orange:
		return "orange"
	case Yellow:
		return "yellow"
	case Red:
		return "red"
	case Blue:
		return "blue"
	case Green:
		return "green"
	default:
		return "unknown"
	}
}
