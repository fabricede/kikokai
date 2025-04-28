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

// only one x, y or z should not be 0
func GetAdjacentFace(face FaceIndex, x, y, z int) FaceIndex {
	switch face {
	case Front:
		if y == 1 {
			return Up
		} else if y == -1 {
			return Down
		} else if z == 1 {
			return Right
		} else if z == -1 {
			return Left
		}
	case Back:
		if y == 1 {
			return Up
		} else if y == -1 {
			return Down
		} else if z == 1 {
			return Right
		} else if z == -1 {
			return Left
		}
	case Up:
		if x == 1 {
			return Front
		} else if x == -1 {
			return Back
		} else if z == 1 {
			return Right
		} else if z == -1 {
			return Left
		}
	case Down:
		if x == 1 {
			return Front
		} else if x == -1 {
			return Back
		} else if z == 1 {
			return Right
		} else if z == -1 {
			return Left
		}
	case Left:
		if x == 1 {
			return Front
		} else if x == -1 {
			return Back
		} else if y == 1 {
			return Up
		} else if y == -1 {
			return Down
		}
	case Right:
		if x == 1 {
			return Front
		} else if x == -1 {
			return Back
		} else if y == 1 {
			return Up
		} else if y == -1 {
			return Down
		}
	}
	return face
}

// GetStickerCoordinate returns the row and column coordinates of a sticker on a face based on its 3D position
func GetStickerCoordinate(face FaceIndex, position CubeCoordinate) (row, col int, err error) {
	// Calculate the local coordinates on the face
	// We need to project the 3D position onto the 2D face
	switch face {
	case Front: // x = 1
		if position.X != 1 {
			return 0, 0, fmt.Errorf("position %v is not on Front face ", position)
		}
		col = 1 + position.Y // Down/South (z=-1) is row 0, Up   /North (z=1) is row 2
		row = 1 + position.Z // Left/West  (y=-1) is col 0, Right/East  (y=1) is col 2
	case Back: // x = -1
		if position.X != -1 {
			return 0, 0, fmt.Errorf("position %v is not on Back face ", position)
		}
		col = 1 + position.Y // Down /South (z=-1) is row 0, Up  /North (z=1) is row 2
		row = 1 + position.Z // Right/West  (y=-1) is col 0, Left/East  (y=1) is col 2
	case Up: // y = 1
		if position.Z != 1 {
			return 0, 0, fmt.Errorf("position %v is not on Up face ", position)
		}
		col = 1 + position.X // Back (x=-1) is row 0, Front (x=1) is row 2
		row = 1 + position.Z // Left (z=-1) is col 0, Right (z=1) is col 2
	case Down: // y = -1
		if position.Z != -1 {
			return 0, 0, fmt.Errorf("position %v is not on Down face ", position)
		}
		col = 1 + position.X // Front (x=1) is row 0, Back (x=-1) is row 2
		row = 1 + position.Z // Left (z=-1) is col 0, Right (z=1) is col 2
	case Left: // z = -1
		if position.Y != -1 {
			return 0, 0, fmt.Errorf("position %v is not on Left face ", position)
		}
		col = 1 + position.X // Back (x=-1) is col 0, Front (x=1) is col 2
		row = 1 + position.Y // Up (y=1) is row 0, Down (y=-1) is row 2
	case Right: // z = 1
		if position.Y != 1 {
			return 0, 0, fmt.Errorf("position %v is not on Right face ", position)
		}
		col = 1 + position.X // Front (x=1) is col 0, Back (x=-1) is col 2
		row = 1 + position.Y // Up (y=1) is row 0, Down (y=-1) is row 2
	}

	return row, col, nil
}

// GetCubePosition returns the 3D position of a sticker based on its face and row/column coordinates
func GetCubePosition(face FaceIndex, row, col int) CubeCoordinate {
	// Start with the face normal vector
	position := FaceToCoordinate(face)

	// Calculate the offsets from the center of the face
	rowOffset := row - 1 // row 0 is bottom, row 2 is top
	colOffset := col - 1 // col 0 is left, col 2 is right

	// Apply the offsets based on the face
	switch face {
	case Front: // x = 1
		position.Y += colOffset // Left to Right
		position.Z += rowOffset // Up to Down
	case Back: // x = -1
		position.Y += colOffset // Right to Left (mirrored)
		position.Z += rowOffset // Up to Down
	case Up: // z = 1
		position.X += rowOffset // Back to Front
		position.Z += colOffset // Left to Right
	case Down: // z = -1
		position.X += rowOffset // Front to Back (mirrored)
		position.Z += colOffset // Left to Right
	case Left: // y = -1
		position.Y += colOffset // Back to Front (mirrored)
		position.Z += rowOffset // Up to Down
	case Right: // y = 1
		position.Y += colOffset // Front to Back
		position.Z += rowOffset // Up to Down
	}

	return position
}

// GetAdjacentXEdge returns the adjacent edge on the X direction for a given face and X coordinate
// face could not already be on the x axis
func GetAdjacentXEdge(cube *Cube, face FaceIndex, x int) (adjacentFace FaceIndex, edge [3]Sticker, err error) {
	// first determine the adjacent face based on the x coordinate
	adjacentFace = GetAdjacentFace(face, x, 0, 0)
	switch face {
	case Up: // (y=1)
		for i := range 3 {
			edge[i], err = cube.GetSticker(adjacentFace, CubeCoordinate{X: x, Y: 1, Z: i - 1})
		}
	case Down: // (y=-1)
		for i := range 3 {
			edge[i], err = cube.GetSticker(adjacentFace, CubeCoordinate{X: x, Y: -1, Z: i - 1})
		}
	case Left: // (z=-1)
		for i := range 3 {
			edge[i], err = cube.GetSticker(adjacentFace, CubeCoordinate{X: x, Y: i - 1, Z: -1})
		}
	case Right: // (z=1)
		for i := range 3 {
			edge[i], err = cube.GetSticker(adjacentFace, CubeCoordinate{X: x, Y: i - 1, Z: 1})
		}
	default:
		return face, [3]Sticker{}, fmt.Errorf("invalid face %d", face)
	}

	return adjacentFace, edge, nil
}

// GetAdjacentYEdge returns the adjacent edge on the Y direction for a given face and Y coordinate
// face could not already be on the y axis
func GetAdjacentYEdge(cube *Cube, face FaceIndex, y int) (adjacentFace FaceIndex, edge [3]Sticker, err error) {
	// fisrt determine the adjacent face based on the y coordinate
	adjacentFace = GetAdjacentFace(face, 0, y, 0)
	switch face {
	case Front: // (x=1)
		for i := range 3 {
			edge[i], err = cube.GetSticker(adjacentFace, CubeCoordinate{X: 1, Y: y, Z: i - 1})
		}
	case Back: // (x=-1)
		for i := range 3 {
			edge[i], err = cube.GetSticker(adjacentFace, CubeCoordinate{X: -1, Y: y, Z: i - 1})
		}
	case Left: // (z=-1)
		for i := range 3 {
			edge[i], err = cube.GetSticker(adjacentFace, CubeCoordinate{X: i - 1, Y: y, Z: -1})
		}
	case Right: // (z=1)
		for i := range 3 {
			edge[i], err = cube.GetSticker(adjacentFace, CubeCoordinate{X: i - 1, Y: y, Z: 1})
		}
	default:
		return face, [3]Sticker{}, fmt.Errorf("invalid face %d", face)
	}

	return adjacentFace, edge, nil
}

// GetAdjacentZEdge returns the adjacent edge on the Z direction for a given face and Z coordinate
// face could not already be on the z axis
func GetAdjacentZEdge(cube *Cube, face FaceIndex, z int) (adjacentFace FaceIndex, edge [3]Sticker, err error) {
	// fisrt determine the adjacent face based on the z coordinate
	adjacentFace = GetAdjacentFace(face, 0, 0, z)
	switch face {
	case Front: // (x=1)
		for i := range 3 {
			edge[i], err = cube.GetSticker(adjacentFace, CubeCoordinate{X: 1, Y: i - 1, Z: z})
		}
	case Back: // (x=-1)
		for i := range 3 {
			edge[i], err = cube.GetSticker(adjacentFace, CubeCoordinate{X: -1, Y: i - 1, Z: z})
		}
	case Up: // (y=1)
		for i := range 3 {
			edge[i], err = cube.GetSticker(adjacentFace, CubeCoordinate{X: i - 1, Y: 1, Z: z})
		}
	case Down: // (y=-1)
		for i := range 3 {
			edge[i], err = cube.GetSticker(adjacentFace, CubeCoordinate{X: i - 1, Y: -1, Z: z})
		}
	default:
		return face, [3]Sticker{}, fmt.Errorf("invalid face %d", face)
	}

	return adjacentFace, edge, nil
}
