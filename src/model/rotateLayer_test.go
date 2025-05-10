package model

import (
	"maps"
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
			var matrix Layer
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

// TestFaceToAxisLayer tests the conversion from face to axis+layer parameters
func TestFaceToAxisLayer(t *testing.T) {
	testCases := []struct {
		name      string
		face      FaceIndex
		clockwise TurningDirection
		wantAxis  CubeCoordinate
		wantLayer int
		wantClock bool
	}{
		{
			name:      "Front face clockwise",
			face:      Front,
			clockwise: Clockwise,
			wantAxis:  CubeCoordinate{Z: 1},
			wantLayer: 2,
			wantClock: true,
		},
		{
			name:      "Back face clockwise",
			face:      Back,
			clockwise: Clockwise,
			wantAxis:  CubeCoordinate{Z: 1},
			wantLayer: 0,
			wantClock: false, // Inverted for back face
		},
		{
			name:      "Up face counter-clockwise",
			face:      Up,
			clockwise: CounterClockwise,
			wantAxis:  CubeCoordinate{Y: 1},
			wantLayer: 2,
			wantClock: false,
		},
		{
			name:      "Left face counter-clockwise",
			face:      Left,
			clockwise: CounterClockwise,
			wantAxis:  CubeCoordinate{X: 1},
			wantLayer: 0,
			wantClock: true, // Inverted for left face
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			axis, layer, isClock := faceToAxisLayerDirection(tc.face, tc.clockwise)

			if axis.X != tc.wantAxis.X || axis.Y != tc.wantAxis.Y || axis.Z != tc.wantAxis.Z {
				t.Errorf("faceToAxisLayerDirection(%v, %v) got axis = %v, want %v", tc.face, tc.clockwise, axis, tc.wantAxis)
			}

			if layer != tc.wantLayer {
				t.Errorf("faceToAxisLayerDirection(%v, %v) got layer = %d, want %d", tc.face, tc.clockwise, layer, tc.wantLayer)
			}

			if isClock != tc.wantClock {
				t.Errorf("faceToAxisLayerDirection(%v, %v) got clockwise = %v, want %v", tc.face, tc.clockwise, isClock, tc.wantClock)
			}
		})
	}
}

// TestCubieRotation tests that cubies' orientations are correctly transformed during rotation
func TestCubieRotation(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name      string
		face      FaceIndex
		direction TurningDirection
		position  [3]int // [x, y, z] of the cubie to check
		// map of face -> face indicating which initial face's color should be on which face after rotation
		// e.g., {Up: Right} means the color initially on Up face should be on Right face after rotation
		colorTransformations map[FaceIndex]FaceIndex
		// face that should remain unchanged during this rotation
		unchangedFaces []FaceIndex
	}{
		{
			name:      "Front face clockwise rotation",
			face:      Front,
			direction: Clockwise,
			position:  [3]int{2, 0, 2}, // top-right cubie on Front face
			colorTransformations: map[FaceIndex]FaceIndex{
				Up:    Left,
				Right: Up,
				Down:  Right,
				Left:  Down,
			},
			unchangedFaces: []FaceIndex{Front, Back},
		},
		{
			name:      "Up face clockwise rotation",
			face:      Up,
			direction: Clockwise,
			position:  [3]int{2, 0, 0}, // top-right cubie on Up face
			colorTransformations: map[FaceIndex]FaceIndex{
				Front: Right,
				Right: Back,
				Back:  Left,
				Left:  Front,
			},
			unchangedFaces: []FaceIndex{Up, Down},
		},
		{
			name:      "Right face counter-clockwise rotation",
			face:      Right,
			direction: CounterClockwise,
			position:  [3]int{2, 0, 2}, // top-front cubie on Right face
			colorTransformations: map[FaceIndex]FaceIndex{
				Up:    Back,
				Front: Up,
				Down:  Front,
				Back:  Down,
			},
			unchangedFaces: []FaceIndex{Right, Left},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new cube for this test
			cube := NewCube()

			// Get the test cubie
			testCubie := cube.Cubies[tc.position[0]][tc.position[1]][tc.position[2]]

			initialColors := map[FaceIndex]Color{}
			// Copy the initial colors
			maps.Copy(initialColors, testCubie.Colors)

			// Apply the rotation to the face
			cube.RotateFace(tc.face, tc.direction)

			// Check that the unchanged faces remain unchanged
			for _, face := range tc.unchangedFaces {
				if got, want := testCubie.Colors[face], initialColors[face]; got != want {
					t.Errorf("%d(%s) color changed after rotation: got %v, want %v", face, FaceColorName[face], got, want)
				}
			}

			// Check color transformations
			for newFace, originalFace := range tc.colorTransformations {
				if got, want := testCubie.Colors[newFace], initialColors[originalFace]; got != want {
					t.Errorf("%d(%s) face color incorrect after rotation: got %v, want %v (should be color from %d(%s))",
						newFace, FaceColorName[newFace], got, want, originalFace, FaceColorName[originalFace])
				}
			}
		})
	}
}

// TestAxisRotation tests the axis-based rotation directly
func TestAxisRotation(t *testing.T) {
	testCases := []struct {
		name      string
		axis      CubeCoordinate
		layer     int
		clockwise bool
		// Starting position coordinates to check
		startPos [3]int
		// Expected ending position coordinates after rotation
		endPos [3]int
	}{
		{
			name:      "X-Axis Layer 0 Clockwise",
			axis:      CubeCoordinate{X: 1},
			layer:     0,
			clockwise: true,
			startPos:  [3]int{0, 0, 0},
			endPos:    [3]int{0, 2, 0},
		},
		{
			name:      "Y-Axis Layer 2 Counter-Clockwise",
			axis:      CubeCoordinate{Y: 1},
			layer:     2,
			clockwise: false,
			startPos:  [3]int{0, 2, 0},
			endPos:    [3]int{0, 2, 2},
		},
		{
			name:      "Z-Axis Layer 1 Clockwise",
			axis:      CubeCoordinate{Z: 1},
			layer:     1,
			clockwise: true,
			startPos:  [3]int{0, 0, 1},
			endPos:    [3]int{2, 0, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new cube
			cube := NewCube()

			// Store the pointer to the initial cubie we want to track
			initialCubie := cube.Cubies[tc.startPos[0]][tc.startPos[1]][tc.startPos[2]]
			if initialCubie == nil {
				t.Fatalf("No cubie at initial position %v", tc.startPos)
			}

			// Get a reference to this cubie
			initialPtr := initialCubie

			// Apply rotation
			cube.RotateAxis(tc.axis, tc.layer, tc.clockwise)

			// Validate if the cubie at endPos is the same instance as our initialCubie
			endCubie := cube.Cubies[tc.endPos[0]][tc.endPos[1]][tc.endPos[2]]
			if endCubie == nil {
				t.Fatalf("No cubie at end position %v", tc.endPos)
			}

			if initialPtr != endCubie {
				t.Errorf("Expected cubie at %v to move to %v, but found a different cubie",
					tc.startPos, tc.endPos)
			}
		})
	}
}
