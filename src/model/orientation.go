package model

import "fmt"

// Orientation represents the orientation of a face on the cube
// It is used to determine the position of the faces relative to each other
// For exemple, from the perspective of the cube's Front face
// North is the top face, South is the bottom face, East is the right face, and West is the left face
// Relative to Front face, it's North face is rotated 90 degrees counter-clockwise, It's South face is rotated 90 degrees clockwise
// therefore:
// - it's North border is the Est Border of the North face
// - it's East Border is the North border of the East face
// - it's South border is the West border of the South face
// - it's West border is the East border of the West face

// FaceIndex represents the order of representation of faces on the cube
type FaceIndex int

// TurningDirection for rotation
type TurningDirection bool

// CubeCoordinate represents a position in 3D space (x,y,z)
// where:
// - x-axis: Back (x=-1) to Front (x=1)
// - y-axis: Down (y=-1) to Up (y=1)
// - z-axis: Left (z=-1) to Right (z=1)
type CubeCoordinate struct {
	X, Y, Z int
}

// String returns a string representation of the coordinate
func (c CubeCoordinate) String() string {
	return fmt.Sprintf("(%d,%d,%d)", c.X, c.Y, c.Z)
}

const (
	// Face constants
	Front FaceIndex = iota
	Right
	Back
	Left
	Up
	Down

	// Turning direction
	Clockwise        TurningDirection = true
	CounterClockwise TurningDirection = false
)

// Face coordinate constants
var (
	FrontAxis = CubeCoordinate{X: 1, Y: 0, Z: 0}
	BackAxis  = CubeCoordinate{X: -1, Y: 0, Z: 0}
	UpAxis    = CubeCoordinate{X: 0, Y: 1, Z: 0}
	DownAxis  = CubeCoordinate{X: 0, Y: -1, Z: 0}
	LeftAxis  = CubeCoordinate{X: 0, Y: 0, Z: -1}
	RightAxis = CubeCoordinate{X: 0, Y: 0, Z: 1}
)

// FaceToCoordinate converts a FaceIndex to its corresponding CubeCoordinate
func FaceToCoordinate(face FaceIndex) CubeCoordinate {
	switch face {
	case Front:
		return FrontAxis
	case Back:
		return BackAxis
	case Up:
		return UpAxis
	case Down:
		return DownAxis
	case Left:
		return LeftAxis
	case Right:
		return RightAxis
	default:
		return CubeCoordinate{0, 0, 0}
	}
}

// GetCoordFromAxis transforms axis and layer to Coordinate
func GetCoordFromAxis(axis string, layer int) (face CubeCoordinate) {
	switch axis {
	case "x":
		if layer > 0 {
			face = FrontAxis
		} else {
			face = BackAxis
		}
	case "y":
		if layer > 0 {
			face = UpAxis
		} else {
			face = DownAxis
		}
	case "z":
		if layer > 0 {
			face = RightAxis
		} else {
			face = LeftAxis
		}
	default:
		panic("Invalid axis")
	}
	return face
}
