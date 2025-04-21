package model

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

type Orientation int

// TurningDirection for rotation
type TurningDirection bool

// CubeCoordinate represents a position in 3D space (x,y,z)
// where:
// - x-axis: Back (x=-1) to Front (x=1)
// - y-axis: Left (y=-1) to Right (y=1)
// - z-axis: Down (z=-1) to Up (z=1)
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

	// Orientation constants
	North Orientation = iota
	East
	Center
	West
	South

	// Turning direction
	Clockwise        TurningDirection = true
	CounterClockwise TurningDirection = false
)

// Face coordinate constants
var (
	FrontCoord = CubeCoordinate{X: 1, Y: 0, Z: 0}
	BackCoord  = CubeCoordinate{X: -1, Y: 0, Z: 0}
	UpCoord    = CubeCoordinate{X: 0, Y: 0, Z: 1}
	DownCoord  = CubeCoordinate{X: 0, Y: 0, Z: -1}
	LeftCoord  = CubeCoordinate{X: 0, Y: -1, Z: 0}
	RightCoord = CubeCoordinate{X: 0, Y: 1, Z: 0}
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

// CoordinateToFace converts a CubeCoordinate to its corresponding FaceIndex
func CoordinateToFace(coord CubeCoordinate) FaceIndex {
	switch {
	case coord.X == 1 && coord.Y == 0 && coord.Z == 0:
		return Front
	case coord.X == -1 && coord.Y == 0 && coord.Z == 0:
		return Back
	case coord.X == 0 && coord.Y == 0 && coord.Z == 1:
		return Up
	case coord.X == 0 && coord.Y == 0 && coord.Z == -1:
		return Down
	case coord.X == 0 && coord.Y == -1 && coord.Z == 0:
		return Left
	case coord.X == 0 && coord.Y == 1 && coord.Z == 0:
		return Right
	default:
		// Invalid coordinate, return Front as default
		return Front
	}
}

// GetNorthFace returns the face that is North to the given face when the cube's face is facing us
func GetNorthFace(face FaceIndex) (FaceIndex, Orientation) {
	switch face {
	case Front:
		return Up, West
	case Up:
		return Left, East
	case Back:
		return Down, West
	case Down:
		return Right, West
	case Left:
		return Back, East
	case Right:
		return Front, East
	default:
		return face, Center
	}
}

// GetSouthFace returns the face that is South to the given face when the cube's face is facing us
func GetSouthFace(face FaceIndex) (FaceIndex, Orientation) {
	switch face {
	case Front:
		return Down, East
	case Up:
		return Right, East
	case Back:
		return Up, East
	case Down:
		return Left, East
	case Left:
		return Front, West
	case Right:
		return Back, West
	default:
		return face, Center
	}
}

// GetEastFace returns the face that is East to the given face when the cube's face is facing us
func GetEastFace(face FaceIndex) (FaceIndex, Orientation) {
	switch face {
	case Front:
		return Right, North
	case Up:
		return Back, South
	case Back:
		return Left, North
	case Down:
		return Front, South
	case Left:
		return Down, South
	case Right:
		return Up, South
	default:
		return face, Center
	}
}

// GetWestFace returns the face that is West to the given face when the cube's face is facing us
func GetWestFace(face FaceIndex) (FaceIndex, Orientation) {
	switch face {
	case Front:
		return Left, South
	case Up:
		return Front, North
	case Back:
		return Right, South
	case Down:
		return Back, North
	case Left:
		return Up, North
	case Right:
		return Down, North
	default:
		return face, Center
	}
}

// GetNorthEdge returns the edge that is North to the given face when the cube's face is facing us
func GetNorthEdge(cube *Cube, face FaceIndex) (edge [3]Sticker) {
	north, border := GetNorthFace(face)
	return cube.GetEdge(north, border)
}

// GetSouthEdge returns the edge that is South to the given face when the cube's face is facing us
func GetSouthEdge(cube *Cube, face FaceIndex) (edge [3]Sticker) {
	south, border := GetSouthFace(face)
	return cube.GetEdge(south, border)
}

// GetEastEdge returns the edge that is East to the given face when the cube's face is facing us
func GetEastEdge(cube *Cube, face FaceIndex) (edge [3]Sticker) {
	east, border := GetEastFace(face)
	return cube.GetEdge(east, border)
}

// GetWestEdge returns the edge that is West to the given face when the cube's face is facing us
func GetWestEdge(cube *Cube, face FaceIndex) (edge [3]Sticker) {
	west, border := GetWestFace(face)
	return cube.GetEdge(west, border)
}

// SetNorthEdge sets the edge that is Notrh to the given face when the cube's face is facing us
func SetNorthEdge(cube *Cube, face FaceIndex, edge [3]Sticker) {
	north, border := GetNorthFace(face)
	cube.SetEdge(north, border, edge)
}

// SetSouthEdge sets the edge that is South to the given face when the cube's face is facing us
func SetSouthEdge(cube *Cube, face FaceIndex, edge [3]Sticker) {
	south, border := GetSouthFace(face)
	cube.SetEdge(south, border, edge)
}

// SetEastEdge sets the edge that is East to the given face when the cube's face is facing us
func SetEastEdge(cube *Cube, face FaceIndex, edge [3]Sticker) {
	east, border := GetEastFace(face)
	cube.SetEdge(east, border, edge)
}

// SetWestEdge sets the edge that is West to the given face when the cube's face is facing us
func SetWestEdge(cube *Cube, face FaceIndex, edge [3]Sticker) {
	west, border := GetWestFace(face)
	cube.SetEdge(west, border, edge)
}
