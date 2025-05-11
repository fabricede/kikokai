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
			axis: FrontAxis,
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
			axis: BackAxis,
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
			axis: UpAxis,
			wantColor: map[FaceIndex]Color{
				Front: Orange,
				Right: Yellow,
				Back:  Red,
				Left:  White,
				Up:    Blue,
				Down:  Green,
			},
		},
		{
			name: "Rotate Y-axis negative",
			axis: DownAxis,
			wantColor: map[FaceIndex]Color{
				Front: Red,
				Right: White,
				Back:  Orange,
				Left:  Yellow,
				Up:    Blue,
				Down:  Green,
			},
		},
		{
			name: "Rotate Z-axis positive",
			axis: RightAxis,
			wantColor: map[FaceIndex]Color{
				Front: Green,
				Right: Orange,
				Back:  Blue,
				Left:  Red,
				Up:    White,
				Down:  Yellow,
			},
		},
		{
			name: "Rotate Z-axis negative",
			axis: LeftAxis,
			wantColor: map[FaceIndex]Color{
				Front: Blue,
				Right: Orange,
				Back:  Green,
				Left:  Red,
				Up:    Yellow,
				Down:  White,
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
			axis: FrontAxis,
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
			name: "Rotate X-axis negative counter-clockwise",
			axis: BackAxis,
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
			name: "Rotate Y-axis positive counter-clockwise",
			axis: UpAxis,
			wantColor: map[FaceIndex]Color{
				Front: Red,
				Right: White,
				Back:  Orange,
				Left:  Yellow,
				Up:    Blue,
				Down:  Green,
			},
		},
		{
			name: "Rotate Y-axis negative counter-clockwise",
			axis: DownAxis,
			wantColor: map[FaceIndex]Color{
				Front: Orange,
				Right: Yellow,
				Back:  Red,
				Left:  White,
				Up:    Blue,
				Down:  Green,
			},
		},
		{
			name: "Rotate Z-axis positive counter-clockwise",
			axis: RightAxis,
			wantColor: map[FaceIndex]Color{
				Front: Blue,
				Right: Orange,
				Back:  Green,
				Left:  Red,
				Up:    Yellow,
				Down:  White,
			},
		},
		{
			name: "Rotate Z-axis negative counter-clockwise",
			axis: LeftAxis,
			wantColor: map[FaceIndex]Color{
				Front: Green,
				Right: Orange,
				Back:  Blue,
				Left:  Red,
				Up:    White,
				Down:  Yellow,
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
		FrontAxis, BackAxis,
		UpAxis, DownAxis,
		RightAxis, LeftAxis,
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
		FrontAxis, BackAxis,
		UpAxis, DownAxis,
		RightAxis, LeftAxis,
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
