package model

// Matrix5x5 represents a 5x5 matrix
type Matrix5x5 [5][5]Sticker

func SetMatrix(init [5][5]StickerIndex) (m *Matrix5x5) {
	m = &Matrix5x5{}
	for i := range 5 {
		for j := range 5 {
			m[i][j] = Sticker{Index: init[i][j]}
		}
	}
	return m
}

// RotateClockwise rotates the 5x5 matrix 90 degrees clockwise
func (m Matrix5x5) RotateClockwise() Matrix5x5 {
	var result Matrix5x5
	for i := range 5 {
		for j := range 5 {
			result[j][4-i] = m[i][j]
		}
	}
	return result
}

// RotateCounterClockwise rotates the 5x5 matrix 90 degrees counter-clockwise
func (m Matrix5x5) RotateCounterClockwise() Matrix5x5 {
	var result Matrix5x5
	for i := range 5 {
		for j := range 5 {
			result[4-j][i] = m[i][j]
		}
	}
	return result
}

// Init initializes the matrix from the cube's face point of view
func (m *Matrix5x5) Init(c *Cube, face FaceIndex) {

	// Copy the face's state into the matrix center
	for i := range 3 {
		for j := range 3 {
			m[i+1][j+1] = c.State[face].Stickers[i][j]
		}
	}

	faceCoord := FaceToCoordinate(face)
	var xcol, ycol, yrow, zrow int
	if faceCoord.X != 0 {
		ycol = 1
		zrow = 1
	} else {
		xcol = 1
		yrow = 1
	}
	// Get the adjacent face on row = 2
	r2Face := GetAdjacentFace(face, 0, 1*yrow, 1*zrow)
	// Get the r2 edge
	r2Edge := c.GetEdge(r2Face, faceCoord)
	// Copy the r2 edge to the matrix
	m[0][1] = r2Edge[0]
	m[0][2] = r2Edge[1]
	m[0][3] = r2Edge[2]
	// Get the adjacent face on row = 0
	r0Face := GetAdjacentFace(face, 0, -1*yrow, -1*zrow)
	// Get the r0 edge
	r0Edge := c.GetEdge(r0Face, faceCoord)
	// Copy the r0 edge to the matrix
	m[4][1] = r0Edge[0]
	m[4][2] = r0Edge[1]
	m[4][3] = r0Edge[2]
	// Get the adjacent face on col = 0
	c0Face := GetAdjacentFace(face, -1*xcol, -1*ycol, 0)
	// Get the r0 edge
	c0Edge := c.GetEdge(c0Face, faceCoord)
	// Copy the c0 edge to the matrix
	m[1][0] = c0Edge[0]
	m[2][0] = c0Edge[1]
	m[3][0] = c0Edge[2]
	// Get the adjacent face on col = 2
	c2Face := GetAdjacentFace(face, 1*xcol, 1*ycol, 0)
	// Get the r0 edge
	c2Edge := c.GetEdge(c2Face, faceCoord)
	m[1][4] = c2Edge[0]
	m[2][4] = c2Edge[1]
	m[3][4] = c2Edge[2]
}

// SetCube updates the cube's state from the matrix
func (m *Matrix5x5) SetCube(c *Cube, face FaceIndex) {
	// Copy the matrix back to the cube's face
	for i := range 3 {
		for j := range 3 {
			c.State[face].Stickers[i][j] = m[i+1][j+1]
		}
	}
	faceCoord := FaceToCoordinate(face)
	var xcol, ycol, yrow, zrow int
	if faceCoord.X != 0 {
		ycol = 1
		zrow = 1
	} else {
		xcol = 1
		yrow = 1
	}
	// Get the adjacent face on row = 2
	r2Face := GetAdjacentFace(face, 0, 1*yrow, 1*zrow)
	// Set the r2 edge
	r2Edge := [3]Sticker{m[0][1], m[0][2], m[0][3]}
	c.SetEdge(r2Face, faceCoord, r2Edge)
	// Get the adjacent face on row = 0
	r0Face := GetAdjacentFace(face, 0, -1*yrow, -1*zrow)
	// Set the r0 edge
	r0Edge := [3]Sticker{m[4][1], m[4][2], m[4][3]}
	c.SetEdge(r0Face, faceCoord, r0Edge)
	// Get the adjacent face on col = 0
	c0Face := GetAdjacentFace(face, -1*xcol, -1*ycol, 0)
	// Set the c0 edge
	c0Edge := [3]Sticker{m[1][0], m[2][0], m[3][0]}
	c.SetEdge(c0Face, faceCoord, c0Edge)
	// Get the adjacent face on col = 2
	c2Face := GetAdjacentFace(face, 1*xcol, 1*ycol, 0)
	// Set the c2 edge
	c2Edge := [3]Sticker{m[1][4], m[2][4], m[3][4]}
	c.SetEdge(c2Face, faceCoord, c2Edge)
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
