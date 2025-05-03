package model

import (
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

// ------------------------------------------
// Cubie represents a single piece of the Rubik's Cube with colors on its faces.
// ------------------------------------------
type Cubie struct {
	Colors map[FaceIndex]Color
}

// NewCubie initializes a cubie with colors based on its position (x, y, z).
func NewCubie(x, y, z int) *Cubie {
	c := &Cubie{Colors: make(map[FaceIndex]Color)}
	if x == 0 {
		c.Colors[Left] = Orange
	} else if x == 2 {
		c.Colors[Right] = Red
	}
	if y == 0 {
		c.Colors[Down] = Yellow
	} else if y == 2 {
		c.Colors[Up] = White
	}
	if z == 0 {
		c.Colors[Front] = Green
	} else if z == 2 {
		c.Colors[Back] = Blue
	}
	return c
}

// -------------------------------------------
// Cube represents the Rubik's Cube as a 3x3x3 array of cubies.
// --------------------------------------------
type Cube struct {
	cubies [3][3][3]*Cubie
}

// NewCube initializes a solved Rubik's Cube.
func NewCube() *Cube {
	cube := &Cube{}
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			for z := 0; z < 3; z++ {
				cube.cubies[x][y][z] = NewCubie(x, y, z)
			}
		}
	}
	return cube
}

// GetFaceColor returns the color of a specific face and position.
func (c *Cube) GetFaceColor(face FaceIndex, x, y int) Color {
	switch face {
	case Front:
		return c.cubies[x][y][0].Colors[Front]
	case Back:
		return c.cubies[x][y][2].Colors[Back]
	case Left:
		return c.cubies[0][y][x].Colors[Left]
	case Right:
		return c.cubies[2][y][2-x].Colors[Right]
	case Up:
		return c.cubies[x][2][2-y].Colors[Up]
	case Down:
		return c.cubies[x][0][y].Colors[Down]
	default:
		return Green // Default case, shouldn't happen
	}
}

// RotateFace rotates the specified face
func (c *Cube) RotateFace(face FaceIndex, clockwise TurningDirection) {
	RotateFace(c, face, clockwise)
}

// Scramble applies a series of random rotations to the cube
func (c *Cube) Scramble(moves int) {
	// Apply random rotations
	for i := 0; i < moves; i++ {
		// Random face (0-5)
		face := FaceIndex(rand.Intn(6))

		// Random direction (true/false for clockwise/counter-clockwise)
		clockwise := TurningDirection(rand.Intn(2) == 1)

		// Rotate the face
		c.RotateFace(face, clockwise)
	}
}
