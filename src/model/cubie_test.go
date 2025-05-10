package model

import (
	"maps"
	"reflect"
	"testing"
)

func TestCubie_RotateClockwise(t *testing.T) {
	tests := []struct {
		name      string
		axis      CubeCoordinate
		wantColor map[FaceIndex]Color
	}{
		{
			name: "Rotate X-axis positive",
			axis: CubeCoordinate{X: 1},
			wantColor: map[FaceIndex]Color{
				Front: White,
				Right: Blue,
				Back:  Yellow,
				Left:  Green,
				Up:    Red,
				Down:  Orange,
			},
		},
		{
			name: "Rotate X-axis negative",
			axis: CubeCoordinate{X: -1},
			wantColor: map[FaceIndex]Color{
				Front: White,
				Right: Green,
				Back:  Yellow,
				Left:  Blue,
				Up:    Orange,
				Down:  Red,
			},
		},
		{
			name: "Rotate Y-axis positive",
			axis: CubeCoordinate{Y: 1},
			wantColor: map[FaceIndex]Color{
				Front: Orange, // Right → Front (5→0)
				Right: Yellow, // Back → Right (1→5)
				Back:  Red,    // Left → Back (4→1)
				Left:  White,  // Front → Left (0→4)
				Up:    Blue,   // Unchanged (2)
				Down:  Green,  // Unchanged (3)
			},
		},
		{
			name: "Rotate Y-axis negative",
			axis: CubeCoordinate{Y: -1},
			wantColor: map[FaceIndex]Color{
				Front: Red,    // Left → Front (4→0)
				Right: White,  // Front → Right (0→5)
				Back:  Orange, // Right → Back (5→1)
				Left:  Yellow, // Back → Left (1→4)
				Up:    Blue,   // Unchanged (2)
				Down:  Green,  // Unchanged (3)
			},
		},
		{
			name: "Rotate Z-axis positive",
			axis: CubeCoordinate{Z: 1},
			wantColor: map[FaceIndex]Color{
				Front: Green,  // Down → Front (3→0)
				Right: Orange, // Unchanged (1)
				Back:  Blue,   // Up → Back (2→1)
				Left:  Red,    // Unchanged (4)
				Up:    White,  // Front → Up (0→2)
				Down:  Yellow, // Back → Down (1→3)
			},
		},
		{
			name: "Rotate Z-axis negative",
			axis: CubeCoordinate{Z: -1},
			wantColor: map[FaceIndex]Color{
				Front: Blue,   // Up → Front (2→0)
				Right: Orange, // Unchanged (1)
				Back:  Green,  // Down → Back (3→1)
				Left:  Red,    // Unchanged (4)
				Up:    Yellow, // Back → Up (1→2)
				Down:  White,  // Front → Down (0→3)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cubie := NewCubie()
			cubie.rotateClockwise(tt.axis)
			if !reflect.DeepEqual(cubie.Colors, tt.wantColor) {
				t.Errorf("rotateClockwise() got = %v, want %v", cubie.Colors, tt.wantColor)
			}
		})
	}
}

func TestCubie_RotateCounterClockwise(t *testing.T) {
	tests := []struct {
		name      string
		axis      CubeCoordinate
		wantColor map[FaceIndex]Color
	}{
		{
			name: "Rotate X-axis positive counter-clockwise",
			axis: CubeCoordinate{X: 1},
			wantColor: map[FaceIndex]Color{
				Front: White,  // Unchanged (0)
				Right: Green,  // Unchanged (1)
				Back:  Yellow, // Unchanged (1)
				Left:  Blue,   // Up → Left (2→4)
				Up:    Orange, // Right → Up (5→2)
				Down:  Red,    // Left → Down (4→3)
			},
		},
		{
			name: "Rotate X-axis negative counter-clockwise",
			axis: CubeCoordinate{X: -1},
			wantColor: map[FaceIndex]Color{
				Front: White,  // Unchanged
				Right: Blue,   // Down → Right
				Back:  Yellow, // Unchanged
				Left:  Green,  // Up → Left
				Up:    Red,    // Left → Up
				Down:  Orange, // Right → Down
			},
		},
		{
			name: "Rotate Y-axis positive counter-clockwise",
			axis: CubeCoordinate{Y: 1},
			wantColor: map[FaceIndex]Color{
				Front: Red,    // Left → Front (4→0)
				Right: White,  // Front → Right (0→5)
				Back:  Orange, // Right → Back (5→1)
				Left:  Yellow, // Back → Left (1→4)
				Up:    Blue,   // Unchanged (2)
				Down:  Green,  // Unchanged (3)
			},
		},
		{
			name: "Rotate Y-axis negative counter-clockwise",
			axis: CubeCoordinate{Y: -1},
			wantColor: map[FaceIndex]Color{
				Front: Orange, // Right → Front (5→0)
				Right: Yellow, // Back → Right (1→5)
				Back:  Red,    // Left → Back (4→1)
				Left:  White,  // Front → Left (0→4)
				Up:    Blue,   // Unchanged (2)
				Down:  Green,  // Unchanged (3)
			},
		},
		{
			name: "Rotate Z-axis positive counter-clockwise",
			axis: CubeCoordinate{Z: 1},
			wantColor: map[FaceIndex]Color{
				Front: Blue,   // Up → Front (2→0)
				Right: Orange, // Unchanged (1)
				Back:  Green,  // Down → Back (3→1)
				Left:  Red,    // Unchanged (4)
				Up:    Yellow, // Back → Up (1→2)
				Down:  White,  // Front → Down (0→3)
			},
		},
		{
			name: "Rotate Z-axis negative counter-clockwise",
			axis: CubeCoordinate{Z: -1},
			wantColor: map[FaceIndex]Color{
				Front: Green,  // Down → Front (3→0)
				Right: Orange, // Unchanged (1)
				Back:  Blue,   // Up → Back (2→1)
				Left:  Red,    // Unchanged (4)
				Up:    White,  // Front → Up (0→2)
				Down:  Yellow, // Back → Down (1→3)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cubie := NewCubie()
			cubie.rotateCounterClockwise(tt.axis)
			if !reflect.DeepEqual(cubie.Colors, tt.wantColor) {
				t.Errorf("rotateCounterClockwise() got = %v, want %v", cubie.Colors, tt.wantColor)
			}
		})
	}
}

// TestRotateIdentity verifies that rotating clockwise and then counter-clockwise (or vice versa) returns to the original state
func TestRotateIdentity(t *testing.T) {
	axes := []CubeCoordinate{
		{X: 1}, {X: -1},
		{Y: 1}, {Y: -1},
		{Z: 1}, {Z: -1},
	}

	for _, axis := range axes {
		t.Run("Clockwise then Counter-clockwise on axis "+axis.String(), func(t *testing.T) {
			// Create a cubie with all faces colored
			original := NewCubie()
			originalColors := maps.Clone(original.Colors)

			// Rotate clockwise then counter-clockwise
			original.rotateClockwise(axis)
			original.rotateCounterClockwise(axis)

			// Should be back to original state
			if !reflect.DeepEqual(original.Colors, originalColors) {
				t.Errorf("Identity rotation failed for axis %v. Got %v, want %v",
					axis, original.Colors, originalColors)
			}
		})

		t.Run("Counter-clockwise then Clockwise on axis "+axis.String(), func(t *testing.T) {
			// Create a cubie with all faces colored
			original := NewCubie()
			originalColors := maps.Clone(original.Colors)

			// Rotate counter-clockwise then clockwise
			original.rotateCounterClockwise(axis)
			original.rotateClockwise(axis)

			// Should be back to original state
			if !reflect.DeepEqual(original.Colors, originalColors) {
				t.Errorf("Identity rotation failed for axis %v. Got %v, want %v",
					axis, original.Colors, originalColors)
			}
		})
	}
}

// DebugRotationFunctions helps diagnose the actual behavior of rotations
func TestDebugRotationFunctions(t *testing.T) {
	axes := []CubeCoordinate{
		{X: 1}, {X: -1},
		{Y: 1}, {Y: -1},
		{Z: 1}, {Z: -1},
	}

	for _, axis := range axes {
		t.Run("Debug rotation for axis "+axis.String(), func(t *testing.T) {
			// Create a cubie with standard colors
			cubie := NewCubie()

			// Print initial state
			t.Logf("Initial state: %v", cubie.Colors)

			// Rotate clockwise and print
			cubie.rotateClockwise(axis)
			t.Logf("After clockwise rotation on %v: %v", axis, cubie.Colors)

			// Create another cubie for counter-clockwise
			cubie2 := NewCubie()
			cubie2.rotateCounterClockwise(axis)
			t.Logf("After counter-clockwise rotation on %v: %v", axis, cubie2.Colors)
		})
	}
}
