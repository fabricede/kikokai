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

// Cube represents a Rubik's cube with 6 faces, each face is a 3x3 grid
type Cube struct {
	State [6][3][3]string // 6 faces, 3x3 grid per face
}

// NewCube creates and initializes a new Rubik's cube
func NewCube() *Cube {
	c := &Cube{}
	colors := [6]string{"white", "yellow", "blue", "green", "red", "orange"}
	for i, color := range colors {
		for row := 0; row < 3; row++ {
			for col := 0; col < 3; col++ {
				c.State[i][row][col] = color
			}
		}
	}
	return c
}

// RotateFace rotates the specified face and updates adjacent faces
func (c *Cube) RotateFace(face Face, clockwise Direction) {
	// First rotate the face itself
	if clockwise {
		rotateMatrixClockwise(&c.State[face])
	} else {
		rotateMatrixCounterClockwise(&c.State[face])
	}

	// Then update the adjacent faces
	c.updateAdjacentFaces(face, clockwise)
}

// rotateMatrixClockwise rotates a 3x3 matrix clockwise in place
func rotateMatrixClockwise(matrix *[3][3]string) {
	temp := *matrix
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[j][2-i] = temp[i][j]
		}
	}
}

// rotateMatrixCounterClockwise rotates a 3x3 matrix counter-clockwise in place
func rotateMatrixCounterClockwise(matrix *[3][3]string) {
	temp := *matrix
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[2-j][i] = temp[i][j]
		}
	}
}

// updateAdjacentFaces updates the adjacent faces after a face rotation
func (c *Cube) updateAdjacentFaces(face Face, clockwise Direction) {
	switch face {
	case Front:
		c.rotateFrontAdjacents(clockwise)
	case Back:
		c.rotateBackAdjacents(clockwise)
	case Up:
		c.rotateUpAdjacents(clockwise)
	case Down:
		c.rotateDownAdjacents(clockwise)
	case Left:
		c.rotateLeftAdjacents(clockwise)
	case Right:
		c.rotateRightAdjacents(clockwise)
	}
}

// Helper functions to rotate adjacent face edges
func (c *Cube) rotateFrontAdjacents(clockwise Direction) {
	temp := [3]string{}

	// Save top row of Down face
	for i := 0; i < 3; i++ {
		temp[i] = c.State[Up][2][i]
	}

	if clockwise {
		// Up -> Right
		for i := 0; i < 3; i++ {
			c.State[Up][2][i] = c.State[Left][2-i][2]
		}

		// Left -> Down
		for i := 0; i < 3; i++ {
			c.State[Left][i][2] = c.State[Down][0][i]
		}

		// Down -> Right
		for i := 0; i < 3; i++ {
			c.State[Down][0][i] = c.State[Right][2-i][0]
		}

		// Temp (original Up) -> Right
		for i := 0; i < 3; i++ {
			c.State[Right][i][0] = temp[i]
		}
	} else {
		// Up -> Left
		for i := 0; i < 3; i++ {
			c.State[Up][2][i] = c.State[Right][i][0]
		}

		// Right -> Down
		for i := 0; i < 3; i++ {
			c.State[Right][i][0] = c.State[Down][0][2-i]
		}

		// Down -> Left
		for i := 0; i < 3; i++ {
			c.State[Down][0][i] = c.State[Left][i][2]
		}

		// Temp (original Up) -> Left
		for i := 0; i < 3; i++ {
			c.State[Left][2-i][2] = temp[i]
		}
	}
}

func (c *Cube) rotateBackAdjacents(clockwise Direction) {
	temp := [3]string{}

	// Save bottom row of Up face
	for i := 0; i < 3; i++ {
		temp[i] = c.State[Up][0][i]
	}

	if clockwise {
		// Up -> Left
		for i := 0; i < 3; i++ {
			c.State[Up][0][i] = c.State[Right][i][2]
		}

		// Right -> Down
		for i := 0; i < 3; i++ {
			c.State[Right][i][2] = c.State[Down][2][2-i]
		}

		// Down -> Left
		for i := 0; i < 3; i++ {
			c.State[Down][2][i] = c.State[Left][i][0]
		}

		// Temp (original Up) -> Left
		for i := 0; i < 3; i++ {
			c.State[Left][2-i][0] = temp[i]
		}
	} else {
		// Up -> Right
		for i := 0; i < 3; i++ {
			c.State[Up][0][i] = c.State[Left][2-i][0]
		}

		// Left -> Down
		for i := 0; i < 3; i++ {
			c.State[Left][i][0] = c.State[Down][2][i]
		}

		// Down -> Right
		for i := 0; i < 3; i++ {
			c.State[Down][2][i] = c.State[Right][2-i][2]
		}

		// Temp (original Up) -> Right
		for i := 0; i < 3; i++ {
			c.State[Right][i][2] = temp[i]
		}
	}
}

func (c *Cube) rotateUpAdjacents(clockwise Direction) {
	temp := [3]string{}

	// Save top row of Front face
	for i := 0; i < 3; i++ {
		temp[i] = c.State[Front][0][i]
	}

	if clockwise {
		// Front -> Right
		for i := 0; i < 3; i++ {
			c.State[Front][0][i] = c.State[Right][0][i]
		}

		// Right -> Back
		for i := 0; i < 3; i++ {
			c.State[Right][0][i] = c.State[Back][0][i]
		}

		// Back -> Left
		for i := 0; i < 3; i++ {
			c.State[Back][0][i] = c.State[Left][0][i]
		}

		// Temp (original Front) -> Left
		for i := 0; i < 3; i++ {
			c.State[Left][0][i] = temp[i]
		}
	} else {
		// Front -> Left
		for i := 0; i < 3; i++ {
			c.State[Front][0][i] = c.State[Left][0][i]
		}

		// Left -> Back
		for i := 0; i < 3; i++ {
			c.State[Left][0][i] = c.State[Back][0][i]
		}

		// Back -> Right
		for i := 0; i < 3; i++ {
			c.State[Back][0][i] = c.State[Right][0][i]
		}

		// Temp (original Front) -> Right
		for i := 0; i < 3; i++ {
			c.State[Right][0][i] = temp[i]
		}
	}
}

func (c *Cube) rotateDownAdjacents(clockwise Direction) {
	temp := [3]string{}

	// Save bottom row of Front face
	for i := 0; i < 3; i++ {
		temp[i] = c.State[Front][2][i]
	}

	if clockwise {
		// Front -> Left
		for i := 0; i < 3; i++ {
			c.State[Front][2][i] = c.State[Left][2][i]
		}

		// Left -> Back
		for i := 0; i < 3; i++ {
			c.State[Left][2][i] = c.State[Back][2][i]
		}

		// Back -> Right
		for i := 0; i < 3; i++ {
			c.State[Back][2][i] = c.State[Right][2][i]
		}

		// Temp (original Front) -> Right
		for i := 0; i < 3; i++ {
			c.State[Right][2][i] = temp[i]
		}
	} else {
		// Front -> Right
		for i := 0; i < 3; i++ {
			c.State[Front][2][i] = c.State[Right][2][i]
		}

		// Right -> Back
		for i := 0; i < 3; i++ {
			c.State[Right][2][i] = c.State[Back][2][i]
		}

		// Back -> Left
		for i := 0; i < 3; i++ {
			c.State[Back][2][i] = c.State[Left][2][i]
		}

		// Temp (original Front) -> Left
		for i := 0; i < 3; i++ {
			c.State[Left][2][i] = temp[i]
		}
	}
}

func (c *Cube) rotateLeftAdjacents(clockwise Direction) {
	temp := [3]string{}

	// Save leftmost column of Front face
	for i := 0; i < 3; i++ {
		temp[i] = c.State[Front][i][0]
	}

	if clockwise {
		// Front -> Up
		for i := 0; i < 3; i++ {
			c.State[Front][i][0] = c.State[Down][i][0]
		}

		// Down -> Back
		for i := 0; i < 3; i++ {
			c.State[Down][i][0] = c.State[Back][2-i][2]
		}

		// Back -> Up
		for i := 0; i < 3; i++ {
			c.State[Back][i][2] = c.State[Up][2-i][0]
		}

		// Temp (original Front) -> Up
		for i := 0; i < 3; i++ {
			c.State[Up][i][0] = temp[i]
		}
	} else {
		// Front -> Down
		for i := 0; i < 3; i++ {
			c.State[Front][i][0] = c.State[Up][i][0]
		}

		// Up -> Back
		for i := 0; i < 3; i++ {
			c.State[Up][i][0] = c.State[Back][2-i][2]
		}

		// Back -> Down
		for i := 0; i < 3; i++ {
			c.State[Back][i][2] = c.State[Down][2-i][0]
		}

		// Temp (original Front) -> Down
		for i := 0; i < 3; i++ {
			c.State[Down][i][0] = temp[i]
		}
	}
}

func (c *Cube) rotateRightAdjacents(clockwise Direction) {
	temp := [3]string{}

	// Save rightmost column of Front face
	for i := 0; i < 3; i++ {
		temp[i] = c.State[Front][i][2]
	}

	if clockwise {
		// Front -> Down
		for i := 0; i < 3; i++ {
			c.State[Front][i][2] = c.State[Up][i][2]
		}

		// Up -> Back
		for i := 0; i < 3; i++ {
			c.State[Up][i][2] = c.State[Back][2-i][0]
		}

		// Back -> Down
		for i := 0; i < 3; i++ {
			c.State[Back][i][0] = c.State[Down][2-i][2]
		}

		// Temp (original Front) -> Down
		for i := 0; i < 3; i++ {
			c.State[Down][i][2] = temp[i]
		}
	} else {
		// Front -> Up
		for i := 0; i < 3; i++ {
			c.State[Front][i][2] = c.State[Down][i][2]
		}

		// Down -> Back
		for i := 0; i < 3; i++ {
			c.State[Down][i][2] = c.State[Back][2-i][0]
		}

		// Back -> Up
		for i := 0; i < 3; i++ {
			c.State[Back][i][0] = c.State[Up][2-i][2]
		}

		// Temp (original Front) -> Up
		for i := 0; i < 3; i++ {
			c.State[Up][i][2] = temp[i]
		}
	}
}
