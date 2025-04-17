package model

// Face represents the position of a face on the cube
type Face int

// Direction for rotation
type Direction bool

const (
	// Face constants
	Front Face = iota
	Back
	Up
	Down
	Left
	Right

	// Direction constants
	Clockwise        Direction = true
	CounterClockwise Direction = false
)

// GetNorthFace returns the face that is up to the given face when the cube's face is facing us
func GetNorthFace(face Face) Face {
	switch face {
	case Front:
		return Up
	case Up:
		return Back
	case Back:
		return Up
	case Down:
		return Back
	case Left:
		return Up
	case Right:
		return Up
	default:
		return face
	}
}

// GetSouthFace returns the face that is down to the given face when the cube's face is facing us
func GetSouthFace(face Face) Face {
	switch face {
	case Front:
		return Down
	case Up:
		return Front
	case Back:
		return Down
	case Down:
		return Back
	case Left:
		return Down
	case Right:
		return Down
	default:
		return face
	}
}

// GetEastFace returns the face that is right to the given face when the cube's face is facing us
func GetEastFace(face Face) Face {
	switch face {
	case Front:
		return Right
	case Up:
		return Right
	case Back:
		return Left
	case Down:
		return Back
	case Left:
		return Front
	case Right:
		return Back
	default:
		return face
	}
}

// GetWestFace returns the face that is left to the given face when the cube's face is facing us
func GetWestFace(face Face) Face {
	switch face {
	case Front:
		return Left
	case Up:
		return Up
	case Back:
		return Right
	case Down:
		return Down
	case Left:
		return Back
	case Right:
		return Front
	default:
		return face
	}
}

// GetNorthFace returns the edge that is up to the given face when the cube's face is facing us
func GetNorthEdge(cube *Cube, face Face) (edge [3]Color) {
	north := GetNorthFace(face)
	edge[0] = cube.State[north][0][0]
	edge[1] = cube.State[north][1][0]
	edge[2] = cube.State[north][2][0]
	return edge
}

// GetSouthFace returns the edge that is down to the given face when the cube's face is facing us
func GetSouthEdge(cube *Cube, face Face) (edge [3]Color) {
	south := GetSouthFace(face)
	edge[0] = cube.State[south][2][0]
	edge[1] = cube.State[south][2][1]
	edge[2] = cube.State[south][2][2]
	return edge
}

// GetEastFace returns the edge that is right to the given face when the cube's face is facing us
func GetEastEdge(cube *Cube, face Face) (edge [3]Color) {
	east := GetEastFace(face)
	edge[0] = cube.State[east][0][2]
	edge[1] = cube.State[east][1][2]
	edge[2] = cube.State[east][2][2]
	return edge
}

// GetWestFace returns the edge that is left to the given face when the cube's face is facing us
func GetWestEdge(cube *Cube, face Face) (edge [3]Color) {
	west := GetWestFace(face)
	edge[0] = cube.State[west][0][0]
	edge[1] = cube.State[west][1][0]
	edge[2] = cube.State[west][2][0]
	return edge
}

// SetNorthEdge sets the edge that is up to the given face when the cube's face is facing us
func SetNorthEdge(cube *Cube, face Face, edge [3]Color) {
	north := GetNorthFace(face)
	cube.State[north][0][0] = edge[0]
	cube.State[north][1][0] = edge[1]
	cube.State[north][2][0] = edge[2]
}

// SetSouthEdge sets the edge that is down to the given face when the cube's face is facing us
func SetSouthEdge(cube *Cube, face Face, edge [3]Color) {
	south := GetSouthFace(face)
	cube.State[south][2][0] = edge[0]
	cube.State[south][2][1] = edge[1]
	cube.State[south][2][2] = edge[2]
}

// SetEastEdge sets the edge that is right to the given face when the cube's face is facing us
func SetEastEdge(cube *Cube, face Face, edge [3]Color) {
	east := GetEastFace(face)
	cube.State[east][0][2] = edge[0]
	cube.State[east][1][2] = edge[1]
	cube.State[east][2][2] = edge[2]
}

// SetWestEdge sets the edge that is left to the given face when the cube's face is facing us
func SetWestEdge(cube *Cube, face Face, edge [3]Color) {
	west := GetWestFace(face)
	cube.State[west][0][0] = edge[0]
	cube.State[west][1][0] = edge[1]
	cube.State[west][2][0] = edge[2]
}
