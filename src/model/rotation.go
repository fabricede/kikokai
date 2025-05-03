package model

// Matrix5x5 represents a 5x5 matrix
type Matrix5x5 [5][5]Color

func SetMatrix(init [5][5]Color) (m *Matrix5x5) {
	m = &Matrix5x5{}
	for i := range 5 {
		for j := range 5 {
			m[i][j] = init[i][j]
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

// RotateFace rotates a specific face of the cube
func RotateFace(c *Cube, face FaceIndex, clockwise TurningDirection) {
	var m Matrix5x5
	//m.Init(c, face)
	// Rotate the matrix
	if clockwise {
		m = m.RotateClockwise()
	} else {
		m = m.RotateCounterClockwise()
	}
	//m.SetCube(c, face)
}
