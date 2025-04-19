package model

// Matrix5x5 represents a 5x5 matrix
type Matrix5x5 [5][5]Sticker

func SetMatrix(init [5][5]StickerIndex) (m *Matrix5x5) {
	m = &Matrix5x5{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			m[i][j] = Sticker{Index: init[i][j]}
		}
	}
	return m
}

// RotateClockwise rotates the 5x5 matrix 90 degrees clockwise
func (m Matrix5x5) RotateClockwise() Matrix5x5 {
	var result Matrix5x5
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			result[j][4-i] = m[i][j]
		}
	}
	return result
}

// RotateCounterClockwise rotates the 5x5 matrix 90 degrees counter-clockwise
func (m Matrix5x5) RotateCounterClockwise() Matrix5x5 {
	var result Matrix5x5
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			result[4-j][i] = m[i][j]
		}
	}
	return result
}

// Init initializes the matrix from the cube's face point of view
func (m *Matrix5x5) Init(c *Cube, face FaceIndex) {
	// Copy the face's state into the matrix center
	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			m[i][j] = c.State[face].Stickers[i-1][j-1]
		}
	}
	// Get the north edge
	northEdge := GetNorthEdge(c, face)
	// Copy the north edge to the matrix
	m[0][1] = northEdge[0]
	m[0][2] = northEdge[1]
	m[0][3] = northEdge[2]
	// Get the south edge
	southEdge := GetSouthEdge(c, face)
	// Copy the south edge to the matrix
	m[4][1] = southEdge[0]
	m[4][2] = southEdge[1]
	m[4][3] = southEdge[2]
	// Get the west edge
	westEdge := GetWestEdge(c, face)
	// Copy the west edge to the matrix
	m[1][0] = westEdge[0]
	m[2][0] = westEdge[1]
	m[3][0] = westEdge[2]
	// Get the east edge
	eastEdge := GetEastEdge(c, face)
	// Copy the east edge to the matrix
	m[1][4] = eastEdge[0]
	m[2][4] = eastEdge[1]
	m[3][4] = eastEdge[2]
}

// SetCube updates the cube's state from the matrix
func (m *Matrix5x5) SetCube(c *Cube, face FaceIndex) {
	// Copy the matrix back to the cube's face
	stickers := c.State[face].Stickers
	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			c.State[face].Stickers[i-1][j-1] = m[i][j]
		}
	}
	c.State[face].Stickers = stickers
	// Copy the north edge back to the cube
	northEdge := [3]Sticker{m[0][1], m[0][2], m[0][3]}
	SetNorthEdge(c, face, northEdge)
	// Copy the south edge back to the cube
	southEdge := [3]Sticker{m[4][1], m[4][2], m[4][3]}
	SetSouthEdge(c, face, southEdge)
	// Copy the west edge back to the cube
	westEdge := [3]Sticker{m[1][0], m[2][0], m[3][0]}
	SetWestEdge(c, face, westEdge)
	// Copy the east edge back to the cube
	eastEdge := [3]Sticker{m[1][4], m[2][4], m[3][4]}
	SetEastEdge(c, face, eastEdge)
}

// RotateFace rotates a specific face of the cube
func RotateFace(c *Cube, face FaceIndex, clockwise TurningDirection) {
	var m Matrix5x5
	m.Init(c, face)
	// Rotate the matrix
	if clockwise {
		m = m.RotateClockwise()
	} else {
		m = m.RotateCounterClockwise()
	}
	m.SetCube(c, face)
}
