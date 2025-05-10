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

/*
X-Axis Faces:
Front face is at X=1 (positive X-axis) white
Back face is at X=-1 (negative X-axis) yellow

Y-Axis Faces:
Up face is at Y=1 (positive Y-axis) blue
Down face is at Y=-1 (negative Y-axis) green

Z-Axis Faces:
Right face is at Z=1 (positive Z-axis) orange
Left face is at Z=-1 (negative Z-axis) red
*/

// function to rotate a cubie clockwise
func (cu *Cubie) rotateClockwise(axis CubeCoordinate) {
	// in a cubie all faces rotate
	// clockwise around the axis
	switch {
	case axis.X == 1:
		cu.Colors[Left], cu.Colors[Up], cu.Colors[Right], cu.Colors[Down] =
			cu.Colors[Down], cu.Colors[Left], cu.Colors[Up], cu.Colors[Right]
	case axis.X == -1:
		cu.Colors[Left], cu.Colors[Up], cu.Colors[Right], cu.Colors[Down] =
			cu.Colors[Up], cu.Colors[Right], cu.Colors[Down], cu.Colors[Left]
	case axis.Y == 1:
		cu.Colors[Front], cu.Colors[Left], cu.Colors[Back], cu.Colors[Right] =
			cu.Colors[Right], cu.Colors[Front], cu.Colors[Left], cu.Colors[Back]
	case axis.Y == -1:
		cu.Colors[Front], cu.Colors[Left], cu.Colors[Back], cu.Colors[Right] =
			cu.Colors[Left], cu.Colors[Back], cu.Colors[Right], cu.Colors[Front]
	case axis.Z == 1:
		cu.Colors[Front], cu.Colors[Up], cu.Colors[Back], cu.Colors[Down] =
			cu.Colors[Down], cu.Colors[Front], cu.Colors[Up], cu.Colors[Back]
	case axis.Z == -1:
		cu.Colors[Front], cu.Colors[Up], cu.Colors[Back], cu.Colors[Down] =
			cu.Colors[Up], cu.Colors[Back], cu.Colors[Down], cu.Colors[Front]
	}
}

// function to rotate a cubie counter-clockwise
func (cu *Cubie) rotateCounterClockwise(axis CubeCoordinate) {
	cu.rotateClockwise(CubeCoordinate{X: -axis.X, Y: -axis.Y, Z: -axis.Z})
}
