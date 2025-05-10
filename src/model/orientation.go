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
	Back
	Up
	Down
	Left
	Right

	// Turning direction
	Clockwise        TurningDirection = true
	CounterClockwise TurningDirection = false
)

// Face coordinate constants
var (
	FrontCoord = CubeCoordinate{X: 1, Y: 0, Z: 0}
	BackCoord  = CubeCoordinate{X: -1, Y: 0, Z: 0}
	UpCoord    = CubeCoordinate{X: 0, Y: 1, Z: 0}
	DownCoord  = CubeCoordinate{X: 0, Y: -1, Z: 0}
	LeftCoord  = CubeCoordinate{X: 0, Y: 0, Z: -1}
	RightCoord = CubeCoordinate{X: 0, Y: 0, Z: 1}
)

// FaceToCoordinate converts a FaceIndex to its corresponding CubeCoordinate
func FaceToCoordinate(face FaceIndex) CubeCoordinate {
	switch face {
	case Front:
		return FrontCoord
	case Back:
		return BackCoord
	case Up:
		return UpCoord
	case Down:
		return DownCoord
	case Left:
		return LeftCoord
	case Right:
		return RightCoord
	default:
		return CubeCoordinate{0, 0, 0}
	}
}

// GetCoordFromAxis transforms axis and layer to Coordinate
func GetCoordFromAxis(axis string, layer int) (face CubeCoordinate) {
	switch axis {
	case "x":
		if layer > 0 {
			face = FrontCoord
		} else {
			face = BackCoord
		}
	case "y":
		if layer > 0 {
			face = UpCoord
		} else {
			face = DownCoord
		}
	case "z":
		if layer > 0 {
			face = RightCoord
		} else {
			face = LeftCoord
		}
	default:
		panic("Invalid axis")
	}
	return face
}
