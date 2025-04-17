package model

// Cube represents a Rubik's cube with 6 faces, each face is a 3x3 grid
type Cube struct {
	State [6][3][3]Color // 6 faces, 3x3 grid per face, using Color enum instead of string
}

// NewCube creates and initializes a new Rubik's cube
func NewCube() *Cube {
	c := &Cube{}
	colors := [6]Color{White, Yellow, Blue, Green, Red, Orange}
	for i, color := range colors {
		for row := 0; row < 3; row++ {
			for col := 0; col < 3; col++ {
				c.State[i][row][col] = color
			}
		}
	}
	return c
}

// RotateFace rotates the specified face
func (c *Cube) RotateFace(face Face, clockwise Direction) {
	RotateFace(c, face, true)
}
