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
			case UpAxis:
				// Up face (y = 2)
				x, y, z = i, 2, j
			case DownAxis:
				// Down face (y = 0)
				x, y, z = i, 0, j
			case FrontAxis:
				// Front face (x = 2)
				x, y, z = 2, i, j
			case BackAxis:
				// Back face (x = 0)
				x, y, z = 0, i, j
			case LeftAxis:
				// Left face (z = 0)
				x, y, z = i, j, 0
			case RightAxis:
				// Right face (z = 2)
				x, y, z = i, j, 2
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
			result[j][2-i] = (m[i][j]).rotateClockwise(axis)
		}
	}
	return result
}

// rotateCounterClockwise rotates the 3x3 matrix 90 degrees counter-clockwise and change cubies orientation
func (m Layer) rotateCounterClockwise(axis CubeCoordinate) Layer {
	var result Layer

	for i := range 3 {
		for j := range 3 {
			result[2-j][i] = m[i][j].rotateCounterClockwise(axis)
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
			case UpAxis:
				x, y, z = i, 2, j
			case DownAxis:
				x, y, z = i, 0, j
			case FrontAxis:
				x, y, z = 2, i, j
			case BackAxis:
				x, y, z = 0, i, j
			case LeftAxis:
				x, y, z = i, j, 0
			case RightAxis:
				x, y, z = i, j, 2
			}

			// Assign the matrix position to the cube
			c.Cubies[x][y][z] = layer[i][j]
		}
	}

}
