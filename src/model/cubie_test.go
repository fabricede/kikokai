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
				Front: Red,
				Right: White,
				Back:  Orange,
				Left:  Yellow,
				Up:    Blue,
				Down:  Green,
			},
		},
		{
			name: "Rotate X-axis negative",
			axis: CubeCoordinate{X: -1},
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
			name: "Rotate Y-axis positive",
			axis: CubeCoordinate{Y: 1},
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
			name: "Rotate Y-axis negative",
			axis: CubeCoordinate{Y: -1},
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
			name: "Rotate Z-axis positive",
			axis: CubeCoordinate{Z: 1},
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
			name: "Rotate Z-axis negative",
			axis: CubeCoordinate{Z: -1},
			wantColor: map[FaceIndex]Color{
				Front: White,
				Right: Green,
				Back:  Yellow,
				Left:  Blue,
				Up:    Orange,
				Down:  Red,
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
				Front: Orange,
				Right: Yellow,
				Back:  Red,
				Left:  White,
				Up:    Blue,
				Down:  Green,
			},
		},
		{
			name: "Rotate X-axis negative counter-clockwise",
			axis: CubeCoordinate{X: -1},
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
			name: "Rotate Y-axis positive counter-clockwise",
			axis: CubeCoordinate{Y: 1},
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
			name: "Rotate Y-axis negative counter-clockwise",
			axis: CubeCoordinate{Y: -1},
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
			name: "Rotate Z-axis positive counter-clockwise",
			axis: CubeCoordinate{Z: 1},
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
			name: "Rotate Z-axis negative counter-clockwise",
			axis: CubeCoordinate{Z: -1},
			wantColor: map[FaceIndex]Color{
				Front: White,
				Right: Blue,
				Back:  Yellow,
				Left:  Green,
				Up:    Red,
				Down:  Orange,
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
