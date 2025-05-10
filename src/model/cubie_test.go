package model

import (
	"reflect"
	"testing"
)

func TestCubie_RotateClockwise(t *testing.T) {
	tests := []struct {
		name      string
		cubie     *Cubie
		axis      CubeCoordinate
		wantColor map[FaceIndex]Color
	}{
		{
			name: "Rotate X-axis positive",
			cubie: &Cubie{
				Colors: map[FaceIndex]Color{
					Front: Green,
					Right: Orange,
					Back:  Blue,
					Left:  Red,
					Up:    White,
					Down:  Yellow,
				},
			},
			axis: CubeCoordinate{X: 1},
			wantColor: map[FaceIndex]Color{
				Front: Red,
				Right: Green,
				Back:  Orange,
				Left:  Blue,
				Up:    White,
				Down:  Yellow,
			},
		},
		{
			name: "Rotate X-axis negative",
			cubie: &Cubie{
				Colors: map[FaceIndex]Color{
					Front: Green,
					Right: Orange,
					Back:  Blue,
					Left:  Red,
					Up:    White,
					Down:  Yellow,
				},
			},
			axis: CubeCoordinate{X: -1},
			wantColor: map[FaceIndex]Color{
				Front: Orange,
				Right: Blue,
				Back:  Red,
				Left:  Green,
				Up:    White,
				Down:  Yellow,
			},
		},
		{
			name: "Rotate Y-axis positive",
			cubie: &Cubie{
				Colors: map[FaceIndex]Color{
					Front: Green,
					Right: Orange,
					Back:  Blue,
					Left:  Red,
					Up:    White,
					Down:  Yellow,
				},
			},
			axis: CubeCoordinate{Y: 1},
			wantColor: map[FaceIndex]Color{
				Front: Yellow,
				Right: Orange,
				Back:  White,
				Left:  Red,
				Up:    Green,
				Down:  Blue,
			},
		},
		{
			name: "Rotate Y-axis negative",
			cubie: &Cubie{
				Colors: map[FaceIndex]Color{
					Front: Green,
					Right: Orange,
					Back:  Blue,
					Left:  Red,
					Up:    White,
					Down:  Yellow,
				},
			},
			axis: CubeCoordinate{Y: -1},
			wantColor: map[FaceIndex]Color{
				Front: White,
				Right: Orange,
				Back:  Yellow,
				Left:  Red,
				Up:    Blue,
				Down:  Green,
			},
		},
		{
			name: "Rotate Z-axis positive",
			cubie: &Cubie{
				Colors: map[FaceIndex]Color{
					Front: Green,
					Right: Orange,
					Back:  Blue,
					Left:  Red,
					Up:    White,
					Down:  Yellow,
				},
			},
			axis: CubeCoordinate{Z: 1},
			wantColor: map[FaceIndex]Color{
				Front: Green,
				Right: White,
				Back:  Blue,
				Left:  Yellow,
				Up:    Red,
				Down:  Orange,
			},
		},
		{
			name: "Rotate Z-axis negative",
			cubie: &Cubie{
				Colors: map[FaceIndex]Color{
					Front: Green,
					Right: Orange,
					Back:  Blue,
					Left:  Red,
					Up:    White,
					Down:  Yellow,
				},
			},
			axis: CubeCoordinate{Z: -1},
			wantColor: map[FaceIndex]Color{
				Front: Green,
				Right: Yellow,
				Back:  Blue,
				Left:  White,
				Up:    Orange,
				Down:  Red,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cubie.rotateClockwise(tt.axis)
			if !reflect.DeepEqual(tt.cubie.Colors, tt.wantColor) {
				t.Errorf("rotateClockwise() got = %v, want %v", tt.cubie.Colors, tt.wantColor)
			}
		})
	}
}

func TestCubie_RotateCounterClockwise(t *testing.T) {
	tests := []struct {
		name      string
		cubie     *Cubie
		axis      CubeCoordinate
		wantColor map[FaceIndex]Color
	}{
		{
			name: "Rotate X-axis positive counter-clockwise",
			cubie: &Cubie{
				Colors: map[FaceIndex]Color{
					Front: Green,
					Right: Orange,
					Back:  Blue,
					Left:  Red,
					Up:    White,
					Down:  Yellow,
				},
			},
			axis: CubeCoordinate{X: 1},
			wantColor: map[FaceIndex]Color{
				Front: Orange,
				Right: Blue,
				Back:  Red,
				Left:  Green,
				Up:    White,
				Down:  Yellow,
			},
		},
		{
			name: "Rotate X-axis negative counter-clockwise",
			cubie: &Cubie{
				Colors: map[FaceIndex]Color{
					Front: Green,
					Right: Orange,
					Back:  Blue,
					Left:  Red,
					Up:    White,
					Down:  Yellow,
				},
			},
			axis: CubeCoordinate{X: -1},
			wantColor: map[FaceIndex]Color{
				Front: Red,
				Right: Green,
				Back:  Orange,
				Left:  Blue,
				Up:    White,
				Down:  Yellow,
			},
		},
		{
			name: "Rotate Y-axis positive counter-clockwise",
			cubie: &Cubie{
				Colors: map[FaceIndex]Color{
					Front: Green,
					Right: Orange,
					Back:  Blue,
					Left:  Red,
					Up:    White,
					Down:  Yellow,
				},
			},
			axis: CubeCoordinate{Y: 1},
			wantColor: map[FaceIndex]Color{
				Front: White,
				Right: Orange,
				Back:  Yellow,
				Left:  Red,
				Up:    Blue,
				Down:  Green,
			},
		},
		{
			name: "Rotate Y-axis negative counter-clockwise",
			cubie: &Cubie{
				Colors: map[FaceIndex]Color{
					Front: Green,
					Right: Orange,
					Back:  Blue,
					Left:  Red,
					Up:    White,
					Down:  Yellow,
				},
			},
			axis: CubeCoordinate{Y: -1},
			wantColor: map[FaceIndex]Color{
				Front: Yellow,
				Right: Orange,
				Back:  White,
				Left:  Red,
				Up:    Green,
				Down:  Blue,
			},
		},
		{
			name: "Rotate Z-axis positive counter-clockwise",
			cubie: &Cubie{
				Colors: map[FaceIndex]Color{
					Front: Green,
					Right: Orange,
					Back:  Blue,
					Left:  Red,
					Up:    White,
					Down:  Yellow,
				},
			},
			axis: CubeCoordinate{Z: 1},
			wantColor: map[FaceIndex]Color{
				Front: Green,
				Right: Yellow,
				Back:  Blue,
				Left:  White,
				Up:    Orange,
				Down:  Red,
			},
		},
		{
			name: "Rotate Z-axis negative counter-clockwise",
			cubie: &Cubie{
				Colors: map[FaceIndex]Color{
					Front: Green,
					Right: Orange,
					Back:  Blue,
					Left:  Red,
					Up:    White,
					Down:  Yellow,
				},
			},
			axis: CubeCoordinate{Z: -1},
			wantColor: map[FaceIndex]Color{
				Front: Green,
				Right: White,
				Back:  Blue,
				Left:  Yellow,
				Up:    Red,
				Down:  Orange,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cubie.rotateCounterClockwise(tt.axis)
			if !reflect.DeepEqual(tt.cubie.Colors, tt.wantColor) {
				t.Errorf("rotateCounterClockwise() got = %v, want %v", tt.cubie.Colors, tt.wantColor)
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
			original := &Cubie{
				Colors: map[FaceIndex]Color{
					Front: Green,
					Right: Orange,
					Back:  Blue,
					Left:  Red,
					Up:    White,
					Down:  Yellow,
				},
			}

			// Deep copy original colors for comparison later
			originalColors := make(map[FaceIndex]Color)
			for k, v := range original.Colors {
				originalColors[k] = v
			}

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
			original := &Cubie{
				Colors: map[FaceIndex]Color{
					Front: Green,
					Right: Orange,
					Back:  Blue,
					Left:  Red,
					Up:    White,
					Down:  Yellow,
				},
			}

			// Deep copy original colors for comparison later
			originalColors := make(map[FaceIndex]Color)
			for k, v := range original.Colors {
				originalColors[k] = v
			}

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
