package model

// Matrix5x5 represents a 5x5 matrix
type Matrix5x5 [5][5]int

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
func (m *Matrix5x5) Init(c *Cube, face Face) {
	// Copy the face's state into the matrix center
	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			m[i][j] = int(c.State[face][i-1][j-1])
		}
	}
	// Get the north edge
	northEdge := GetNorthEdge(c, face)
	// Copy the north edge to the matrix
	m[0][1] = int(northEdge[0])
	m[0][2] = int(northEdge[1])
	m[0][3] = int(northEdge[2])
	// Get the south edge
	southEdge := GetSouthEdge(c, face)
	// Copy the south edge to the matrix
	m[4][1] = int(southEdge[0])
	m[4][2] = int(southEdge[1])
	m[4][3] = int(southEdge[2])
	// Get the west edge
	westEdge := GetWestEdge(c, face)
	// Copy the west edge to the matrix
	m[1][0] = int(westEdge[0])
	m[2][0] = int(westEdge[1])
	m[3][0] = int(westEdge[2])
	// Get the east edge
	eastEdge := GetEastEdge(c, face)
	// Copy the east edge to the matrix
	m[1][4] = int(eastEdge[0])
	m[2][4] = int(eastEdge[1])
	m[3][4] = int(eastEdge[2])
}

// SetCube updates the cube's state from the matrix
func (m *Matrix5x5) SetCube(c *Cube, face Face) {
	// Copy the matrix back to the cube's face
	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			c.State[face][i-1][j-1] = Color(m[i][j])
		}
	}
	// Copy the north edge back to the cube
	northEdge := [3]Color{Color(m[0][1]), Color(m[0][2]), Color(m[0][3])}
	SetNorthEdge(c, face, northEdge)
	// Copy the south edge back to the cube
	southEdge := [3]Color{Color(m[4][1]), Color(m[4][2]), Color(m[4][3])}
	SetSouthEdge(c, face, southEdge)
	// Copy the west edge back to the cube
	westEdge := [3]Color{Color(m[1][0]), Color(m[2][0]), Color(m[3][0])}
	SetWestEdge(c, face, westEdge)
	// Copy the east edge back to the cube
	eastEdge := [3]Color{Color(m[1][4]), Color(m[2][4]), Color(m[3][4])}
	SetEastEdge(c, face, eastEdge)
}

// RotateFace rotates a specific face of the cube
func RotateFace(c *Cube, face Face, clockwise Direction) {
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
