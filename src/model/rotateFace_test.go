package model

import (
	"testing"
)

// TestMatrixInit tests the Matrix3x3.init function
func TestMatrixInit(t *testing.T) {
	// Create a new cube
	cube := NewCube()

	// Test all faces
	testCases := []struct {
		name string
		face FaceIndex
		// Expected positions and colors after initialization
		// For each matrix position [i][j], we expect a cubie at 3D position [x][y][z]
		expectedPositions [][3]int
	}{
		{
			name: "Up Face",
			face: Up,
			// For Up face, we expect cubies from the top layer (y=0)
			// Ordered as they would appear in the 3x3 matrix
			expectedPositions: [][3]int{
				{0, 0, 0}, {1, 0, 0}, {2, 0, 0}, // row 0
				{0, 0, 1}, {1, 0, 1}, {2, 0, 1}, // row 1
				{0, 0, 2}, {1, 0, 2}, {2, 0, 2}, // row 2
			},
		},
		{
			name: "Down Face",
			face: Down,
			// For Down face, we expect cubies from the bottom layer (y=2)
			expectedPositions: [][3]int{
				{0, 2, 2}, {1, 2, 2}, {2, 2, 2}, // row 0
				{0, 2, 1}, {1, 2, 1}, {2, 2, 1}, // row 1
				{0, 2, 0}, {1, 2, 0}, {2, 2, 0}, // row 2
			},
		},
		{
			name: "Front Face",
			face: Front,
			// For Front face, we expect cubies from the front layer (z=2)
			expectedPositions: [][3]int{
				{0, 0, 2}, {1, 0, 2}, {2, 0, 2}, // row 0
				{0, 1, 2}, {1, 1, 2}, {2, 1, 2}, // row 1
				{0, 2, 2}, {1, 2, 2}, {2, 2, 2}, // row 2
			},
		},
		{
			name: "Back Face",
			face: Back,
			// For Back face, we expect cubies from the back layer (z=0)
			expectedPositions: [][3]int{
				{0, 0, 0}, {1, 0, 0}, {2, 0, 0}, // row 0
				{0, 1, 0}, {1, 1, 0}, {2, 1, 0}, // row 1
				{0, 2, 0}, {1, 2, 0}, {2, 2, 0}, // row 2
			},
		},
		{
			name: "Left Face",
			face: Left,
			// For Left face, we expect cubies from the left layer (x=0)
			expectedPositions: [][3]int{
				{0, 0, 0}, {0, 0, 1}, {0, 0, 2}, // row 0
				{0, 1, 0}, {0, 1, 1}, {0, 1, 2}, // row 1
				{0, 2, 0}, {0, 2, 1}, {0, 2, 2}, // row 2
			},
		},
		{
			name: "Right Face",
			face: Right,
			// For Right face, we expect cubies from the right layer (x=2)
			expectedPositions: [][3]int{
				{2, 0, 2}, {2, 0, 1}, {2, 0, 0}, // row 0
				{2, 1, 2}, {2, 1, 1}, {2, 1, 0}, // row 1
				{2, 2, 2}, {2, 2, 1}, {2, 2, 0}, // row 2
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var matrix Matrix3x3
			matrix.init(cube, tc.face)

			// Verify each position in the matrix
			posIndex := 0
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					// Expected 3D coordinates in the cube
					expectedX := tc.expectedPositions[posIndex][0]
					expectedY := tc.expectedPositions[posIndex][1]
					expectedZ := tc.expectedPositions[posIndex][2]
					posIndex++

					// The cubie in the matrix should be the same as the one at the expected position in the cube
					expected := cube.Cubies[expectedX][expectedY][expectedZ]
					actual := matrix[i][j]

					if expected != actual {
						t.Errorf("Position [%d][%d]: expected cubie at [%d][%d][%d], got a different cubie",
							i, j, expectedX, expectedY, expectedZ)
					}
				}
			}
		})
	}
}

// TestCubieRotation tests that cubies' orientations are correctly transformed during rotation
func TestCubieRotation(t *testing.T) {
	// Test case: Front face clockwise rotation
	t.Run("Front face clockwise rotation", func(t *testing.T) {
		// Create a new cube for this test
		cube := NewCube()

		// Check initial colors for a few key cubies
		// Top-right cubie on Front face (position [2][2][2])
		topRightFront := cube.Cubies[2][2][2]

		initialColors := map[FaceIndex]Color{}
		// Copy the initial colors
		for face, color := range topRightFront.Colors {
			initialColors[face] = color
		}

		// Apply a clockwise rotation to the Front face
		RotateFace(cube, Front, Clockwise)

		// After rotation, the cubie's colors should be transformed
		// The color that was on Up should now be on Right
		// The color that was on Right should now be on Down
		// The color that was on Down should now be on Left
		// The color that was on Left should now be on Up

		// The Front face color should stay the same
		if got, want := topRightFront.Colors[Front], initialColors[Front]; got != want {
			t.Errorf("Front face color changed after rotation: got %v, want %v", got, want)
		}

		// Check color transformations for a corner cubie
		if got, want := topRightFront.Colors[Up], initialColors[Left]; got != want {
			t.Errorf("Up face color incorrect after rotation: got %v, want %v", got, want)
		}

		if got, want := topRightFront.Colors[Right], initialColors[Up]; got != want {
			t.Errorf("Right face color incorrect after rotation: got %v, want %v", got, want)
		}

		if got, want := topRightFront.Colors[Down], initialColors[Right]; got != want {
			t.Errorf("Down face color incorrect after rotation: got %v, want %v", got, want)
		}

		if got, want := topRightFront.Colors[Left], initialColors[Down]; got != want {
			t.Errorf("Left face color incorrect after rotation: got %v, want %v", got, want)
		}
	})
}
