package model

// Matrix3x3 represents a 3x3 matrix
type Matrix3x3 [3][3]*Cubie

func (m *Matrix3x3) init(c *Cube, face FaceIndex) {
	// Get the cubies for the specified face
	// The face is a 3x3 grid of cubies
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// Map the face coordinates to the cube coordinates
			var x, y, z int
			switch face {
			case Up:
				x, y, z = j, 0, i
			case Down:
				x, y, z = j, 2, 2-i
			case Front:
				x, y, z = j, i, 2
			case Back:
				x, y, z = j, i, 0
			case Left:
				x, y, z = 0, i, j
			case Right:
				x, y, z = 2, i, 2-j
			}

			// Assign the cubie to the matrix position
			m[i][j] = c.Cubies[x][y][z]
		}
	}
}

// rotateClockwise rotates the 3x3 matrix 90 degrees clockwise
func (m Matrix3x3) rotateClockwise() Matrix3x3 {
	var result Matrix3x3
	for i := range 3 {
		for j := range 3 {
			result[j][2-i] = m[i][j]
		}
	}
	return result
}

// rotateCounterClockwise rotates the 3x3 matrix 90 degrees counter-clockwise
func (m Matrix3x3) rotateCounterClockwise() Matrix3x3 {
	var result Matrix3x3
	for i := range 3 {
		for j := range 3 {
			result[2-j][i] = m[i][j]
		}
	}
	return result
}

// Add a method to apply the rotated matrix back to the cube
func (m *Matrix3x3) SetCube(c *Cube, face FaceIndex) {
	for i := range 3 {
		for j := range 3 {
			// Map the face coordinates back to the cube coordinates
			var x, y, z int
			switch face {
			case Up:
				x, y, z = j, 0, i
			case Down:
				x, y, z = j, 2, 2-i
			case Front:
				x, y, z = j, i, 2
			case Back:
				x, y, z = j, i, 0
			case Left:
				x, y, z = 0, i, j
			case Right:
				x, y, z = 2, i, 2-j
			}

			// Get the cubie that will be placed at this position
			cubie := m[i][j]

			// Before setting the cubie in its new position, we need to transform its colors
			// to reflect the rotation of the cubie itself

			// Create a new colors map to hold the transformed colors
			transformedColors := make(map[FaceIndex]Color)

			// The color of the face we're rotating remains the same
			if color, exists := cubie.Colors[face]; exists {
				transformedColors[face] = color
			}

			// Transform the colors on adjacent faces based on the rotation
			switch face {
			case Front:
				// For Front face rotation:
				// Up -> Right -> Down -> Left -> Up
				if color, exists := cubie.Colors[Up]; exists {
					transformedColors[Right] = color
				}
				if color, exists := cubie.Colors[Right]; exists {
					transformedColors[Down] = color
				}
				if color, exists := cubie.Colors[Down]; exists {
					transformedColors[Left] = color
				}
				if color, exists := cubie.Colors[Left]; exists {
					transformedColors[Up] = color
				}

			case Back:
				// For Back face rotation:
				// Up -> Left -> Down -> Right -> Up
				if color, exists := cubie.Colors[Up]; exists {
					transformedColors[Left] = color
				}
				if color, exists := cubie.Colors[Left]; exists {
					transformedColors[Down] = color
				}
				if color, exists := cubie.Colors[Down]; exists {
					transformedColors[Right] = color
				}
				if color, exists := cubie.Colors[Right]; exists {
					transformedColors[Up] = color
				}

			case Up:
				// For Up face rotation:
				// Front -> Right -> Back -> Left -> Front
				if color, exists := cubie.Colors[Front]; exists {
					transformedColors[Right] = color
				}
				if color, exists := cubie.Colors[Right]; exists {
					transformedColors[Back] = color
				}
				if color, exists := cubie.Colors[Back]; exists {
					transformedColors[Left] = color
				}
				if color, exists := cubie.Colors[Left]; exists {
					transformedColors[Front] = color
				}

			case Down:
				// For Down face rotation:
				// Front -> Left -> Back -> Right -> Front
				if color, exists := cubie.Colors[Front]; exists {
					transformedColors[Left] = color
				}
				if color, exists := cubie.Colors[Left]; exists {
					transformedColors[Back] = color
				}
				if color, exists := cubie.Colors[Back]; exists {
					transformedColors[Right] = color
				}
				if color, exists := cubie.Colors[Right]; exists {
					transformedColors[Front] = color
				}

			case Left:
				// For Left face rotation:
				// Up -> Front -> Down -> Back -> Up
				if color, exists := cubie.Colors[Up]; exists {
					transformedColors[Front] = color
				}
				if color, exists := cubie.Colors[Front]; exists {
					transformedColors[Down] = color
				}
				if color, exists := cubie.Colors[Down]; exists {
					transformedColors[Back] = color
				}
				if color, exists := cubie.Colors[Back]; exists {
					transformedColors[Up] = color
				}

			case Right:
				// For Right face rotation:
				// Up -> Back -> Down -> Front -> Up
				if color, exists := cubie.Colors[Up]; exists {
					transformedColors[Back] = color
				}
				if color, exists := cubie.Colors[Back]; exists {
					transformedColors[Down] = color
				}
				if color, exists := cubie.Colors[Down]; exists {
					transformedColors[Front] = color
				}
				if color, exists := cubie.Colors[Front]; exists {
					transformedColors[Up] = color
				}
			}

			// Replace the cubie's color map with the transformed one
			cubie.Colors = transformedColors

			// Update the cube with the transformed cubie
			c.Cubies[x][y][z] = cubie
		}
	}
}

// RotateFace rotates a specific face of the cube
func RotateFace(c *Cube, face FaceIndex, clockwise TurningDirection) {
	var m Matrix3x3
	m.init(c, face)
	// Rotate the matrix
	if clockwise {
		m = m.rotateClockwise()
	} else {
		m = m.rotateCounterClockwise()
	}
	m.SetCube(c, face)
}
