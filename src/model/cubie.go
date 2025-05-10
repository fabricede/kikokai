package model

// ------------------------------------------
// Cubie represents a single piece of the Rubik's Cube with colors on its faces.
// ------------------------------------------
type Cubie struct {
	Colors map[FaceIndex]Color
}

// createTestCubie creates a new cubie with the standard color setup for testing
func NewCubie() *Cubie {
	return createCubie(White, Orange, Yellow, Red, Blue, Green)
}

// createTestCubie creates a new cubie with the standard color setup for testing
func createCubie(front, right, back, left, up, down Color) *Cubie {
	return &Cubie{
		Colors: map[FaceIndex]Color{
			Front: front,
			Right: right,
			Back:  back,
			Left:  left,
			Up:    up,
			Down:  down,
		},
	}
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
