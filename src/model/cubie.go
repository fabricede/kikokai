package model

// ------------------------------------------
// Cubie represents a single piece of the Rubik's Cube with colors on its faces.
// ------------------------------------------
type Cubie struct {
	Colors map[FaceIndex]Color
}

// NewCubie initializes a cubie with colors based on its position (x, y, z).
func NewCubie(x, y, z int) *Cubie {
	c := &Cubie{Colors: make(map[FaceIndex]Color)}

	// The positions are in array coordinates (0,1,2)
	// Assign colors only to visible/external faces

	// Left face (x = 0) gets Red
	if x == 0 {
		c.Colors[Left] = Red
	}

	// Right face (x = 2) gets Orange
	if x == 2 {
		c.Colors[Right] = Orange
	}

	// Down face (y = 0) gets Yellow
	if y == 0 {
		c.Colors[Down] = Yellow
	}

	// Up face (y = 2) gets White
	if y == 2 {
		c.Colors[Up] = White
	}

	// Front face (z = 0) gets Green
	if z == 0 {
		c.Colors[Front] = Green
	}

	// Back face (z = 2) gets Blue
	if z == 2 {
		c.Colors[Back] = Blue
	}

	return c
}

// function to rotate a cubie clockwise
func (cu *Cubie) rotateClockwise(axis CubeCoordinate) {
	// in a cubie all faces rotate
	// clockwise around the axis
	switch {
	case axis.X == 1:
		cu.Colors[Front], cu.Colors[Right], cu.Colors[Back], cu.Colors[Left] =
			cu.Colors[Left], cu.Colors[Front], cu.Colors[Right], cu.Colors[Back]
	case axis.X == -1:
		cu.Colors[Front], cu.Colors[Right], cu.Colors[Back], cu.Colors[Left] =
			cu.Colors[Right], cu.Colors[Back], cu.Colors[Left], cu.Colors[Front]
	case axis.Y == 1:
		cu.Colors[Front], cu.Colors[Up], cu.Colors[Back], cu.Colors[Down] =
			cu.Colors[Down], cu.Colors[Front], cu.Colors[Up], cu.Colors[Back]
	case axis.Y == -1:
		cu.Colors[Front], cu.Colors[Up], cu.Colors[Back], cu.Colors[Down] =
			cu.Colors[Up], cu.Colors[Back], cu.Colors[Down], cu.Colors[Front]
	case axis.Z == 1:
		cu.Colors[Up], cu.Colors[Right], cu.Colors[Down], cu.Colors[Left] =
			cu.Colors[Left], cu.Colors[Up], cu.Colors[Right], cu.Colors[Down]
	case axis.Z == -1:
		cu.Colors[Up], cu.Colors[Right], cu.Colors[Down], cu.Colors[Left] =
			cu.Colors[Right], cu.Colors[Down], cu.Colors[Left], cu.Colors[Up]
	}
}

// function to rotate a cubie counter-clockwise
func (cu *Cubie) rotateCounterClockwise(axis CubeCoordinate) {
	// all faces rotate counter-clockwise around the axis
	switch {
	case axis.X == 1:
		cu.Colors[Front], cu.Colors[Right], cu.Colors[Back], cu.Colors[Left] =
			cu.Colors[Right], cu.Colors[Back], cu.Colors[Left], cu.Colors[Front]
	case axis.X == -1:
		cu.Colors[Front], cu.Colors[Right], cu.Colors[Back], cu.Colors[Left] =
			cu.Colors[Left], cu.Colors[Front], cu.Colors[Right], cu.Colors[Back]
	case axis.Y == 1:
		cu.Colors[Front], cu.Colors[Up], cu.Colors[Back], cu.Colors[Down] =
			cu.Colors[Up], cu.Colors[Back], cu.Colors[Down], cu.Colors[Front]
	case axis.Y == -1:
		cu.Colors[Front], cu.Colors[Up], cu.Colors[Back], cu.Colors[Down] =
			cu.Colors[Down], cu.Colors[Front], cu.Colors[Up], cu.Colors[Back]
	case axis.Z == 1:
		cu.Colors[Up], cu.Colors[Right], cu.Colors[Down], cu.Colors[Left] =
			cu.Colors[Right], cu.Colors[Down], cu.Colors[Left], cu.Colors[Up]
	case axis.Z == -1:
		cu.Colors[Up], cu.Colors[Right], cu.Colors[Down], cu.Colors[Left] =
			cu.Colors[Left], cu.Colors[Up], cu.Colors[Right], cu.Colors[Down]
	}
}
