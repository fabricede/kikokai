package model

// Layer represents a 3x3 matrix of Cubies for a specific layer (face) of the cube.
type Layer [3][3]*Cubie

// Legacy method required for tests
func (m *Layer) init(c *Cube, axis CubeCoordinate) {
	// Get the cubies for the specified face
	// The face is a 3x3 grid of cubies
	for i := range 3 {
		for j := range 3 {
			var x, y, z int
			switch axis {
			case UpCoord:
				// Up face (y = 2)
				x, y, z = j, 2, i
			case DownCoord:
				// Down face (y = 0)
				x, y, z = j, 0, 2-i
			case FrontCoord:
				// Front face (z = 2)
				x, y, z = j, i, 2
			case BackCoord:
				// Back face (z = 0)
				x, y, z = j, i, 0
			case LeftCoord:
				// Left face (x = 0)
				x, y, z = 0, i, j
			case RightCoord:
				// Right face (x = 2)
				x, y, z = 2, i, 2-j
			}
			// Assign the cubie to the matrix position
			m[i][j] = c.Cubies[x][y][z]
		}
	}
}

// rotateClockwise rotates the 3x3 matrix 90 degrees clockwise and change cubies orientation
func (m Layer) rotateClockwise(axis CubeCoordinate) Layer {
	var result Layer

	for i := range 3 {
		for j := range 3 {
			result[j][2-i] = m[i][j]
			result[j][2-i].rotateClockwise(axis)
		}
	}
	return result
}

// rotateCounterClockwise rotates the 3x3 matrix 90 degrees counter-clockwise and change cubies orientation
func (m Layer) rotateCounterClockwise(axis CubeCoordinate) Layer {
	var result Layer

	for i := range 3 {
		for j := range 3 {
			result[2-j][i] = m[i][j]
			result[2-j][i].rotateCounterClockwise(axis)
		}
	}
	return result
}

func (layer *Layer) setLayer(c *Cube, axis CubeCoordinate) {
	// copies the layer of the cube to a matrix
	for i := range 3 {
		for j := range 3 {
			// Map the face coordinates to the cube coordinates
			var x, y, z int
			switch axis {
			case UpCoord:
				x, y, z = j, 0, i
			case DownCoord:
				x, y, z = j, 2, 2-i
			case FrontCoord:
				x, y, z = j, i, 2
			case BackCoord:
				x, y, z = j, i, 0
			case LeftCoord:
				x, y, z = 0, i, j
			case RightCoord:
				x, y, z = 2, i, 2-j
			}

			// Assign the cubie to the matrix position
			c.Cubies[x][y][z] = layer[i][j]
		}
	}

}
