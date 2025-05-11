package model

import (
	"log"
	"reflect"
	"testing"
)

/*
	Colors: map[FaceIndex]Color{
				Front: front,
				Right: right,
				Back:  back,
				Left:  left,
				Up:    up,
				Down:  down,
			},

=== RUN   TestDebugRotationFunctions/Debug_rotation_for_axis_(1,0,0)

	cubie_test.go:248: Initial state: map[0:0 1:1 2:2 3:3 4:4 5:5]
	cubie_test.go:252: After clockwise rotation on (1,0,0): map[0:0 1:4 2:2 3:5 4:3 5:1]
	cubie_test.go:257: After counter-clockwise rotation on (1,0,0): map[0:0 1:5 2:2 3:4 4:1 5:3]

=== RUN   TestDebugRotationFunctions/Debug_rotation_for_axis_(-1,0,0)

	cubie_test.go:248: Initial state: map[0:0 1:1 2:2 3:3 4:4 5:5]
	cubie_test.go:252: After clockwise rotation on (-1,0,0): map[0:0 1:5 2:2 3:4 4:1 5:3]
	cubie_test.go:257: After counter-clockwise rotation on (-1,0,0): map[0:0 1:4 2:2 3:5 4:3 5:1]

=== RUN   TestDebugRotationFunctions/Debug_rotation_for_axis_(0,1,0)

	cubie_test.go:248: Initial state: map[0:0 1:1 2:2 3:3 4:4 5:5]
	cubie_test.go:252: After clockwise rotation on (0,1,0): map[0:1 1:2 2:3 3:0 4:4 5:5]
	cubie_test.go:257: After counter-clockwise rotation on (0,1,0): map[0:3 1:0 2:1 3:2 4:4 5:5]

=== RUN   TestDebugRotationFunctions/Debug_rotation_for_axis_(0,-1,0)

	cubie_test.go:248: Initial state: map[0:0 1:1 2:2 3:3 4:4 5:5]
	cubie_test.go:252: After clockwise rotation on (0,-1,0): map[0:3 1:0 2:1 3:2 4:4 5:5]
	cubie_test.go:257: After counter-clockwise rotation on (0,-1,0): map[0:1 1:2 2:3 3:0 4:4 5:5]

=== RUN   TestDebugRotationFunctions/Debug_rotation_for_axis_(0,0,1)

	cubie_test.go:248: Initial state: map[0:0 1:1 2:2 3:3 4:4 5:5]
	cubie_test.go:252: After clockwise rotation on (0,0,1): map[0:5 1:1 2:4 3:3 4:0 5:2]
	cubie_test.go:257: After counter-clockwise rotation on (0,0,1): map[0:4 1:1 2:5 3:3 4:2 5:0]

=== RUN   TestDebugRotationFunctions/Debug_rotation_for_axis_(0,0,-1)

	cubie_test.go:248: Initial state: map[0:0 1:1 2:2 3:3 4:4 5:5]
	cubie_test.go:252: After clockwise rotation on (0,0,-1): map[0:4 1:1 2:5 3:3 4:2 5:0]
	cubie_test.go:257: After counter-clockwise rotation on (0,0,-1): map[0:5 1:1 2:4 3:3 4:0 5:2]
*/
func TestLayer_rotateClockwise(t *testing.T) {
	tests := []struct {
		name string
		m    Layer
		axis CubeCoordinate
		want Layer
	}{
		// TODO: Add test cases.
		{
			name: "Test Front face init state",
			m: Layer{
				{createCubie(0, 1, 2, 3, 4, 5), createCubie(0, 1, 2, 3, 4, 5), createCubie(0, 1, 2, 3, 4, 5)},
				{createCubie(0, 1, 2, 3, 4, 5), createCubie(0, 1, 2, 3, 4, 5), createCubie(0, 1, 2, 3, 4, 5)},
				{createCubie(0, 1, 2, 3, 4, 5), createCubie(0, 1, 2, 3, 4, 5), createCubie(0, 1, 2, 3, 4, 5)},
			},
			axis: FrontAxis,
			want: Layer{
				{createCubie(0, 4, 2, 5, 3, 1), createCubie(0, 4, 2, 5, 3, 1), createCubie(0, 4, 2, 5, 3, 1)},
				{createCubie(0, 4, 2, 5, 3, 1), createCubie(0, 4, 2, 5, 3, 1), createCubie(0, 4, 2, 5, 3, 1)},
				{createCubie(0, 4, 2, 5, 3, 1), createCubie(0, 4, 2, 5, 3, 1), createCubie(0, 4, 2, 5, 3, 1)},
			},
		},
		{
			name: "Test Front face same state",
			m: Layer{
				{createCubie(0, 0, 0, 0, 0, 0), createCubie(1, 1, 1, 1, 1, 1), createCubie(2, 2, 2, 2, 2, 2)},
				{createCubie(3, 3, 3, 3, 3, 3), createCubie(4, 4, 4, 4, 4, 4), createCubie(5, 5, 5, 5, 5, 5)},
				{createCubie(6, 6, 6, 6, 6, 6), createCubie(7, 7, 7, 7, 7, 7), createCubie(8, 8, 8, 8, 8, 8)},
			},
			axis: FrontAxis,
			want: Layer{
				{createCubie(6, 6, 6, 6, 6, 6), createCubie(3, 3, 3, 3, 3, 3), createCubie(0, 0, 0, 0, 0, 0)},
				{createCubie(7, 7, 7, 7, 7, 7), createCubie(4, 4, 4, 4, 4, 4), createCubie(1, 1, 1, 1, 1, 1)},
				{createCubie(8, 8, 8, 8, 8, 8), createCubie(5, 5, 5, 5, 5, 5), createCubie(2, 2, 2, 2, 2, 2)},
			},
		},
		{
			name: "Test Up face init state",
			m: Layer{
				{createCubie(0, 1, 2, 3, 4, 5), createCubie(0, 1, 2, 3, 4, 5), createCubie(0, 1, 2, 3, 4, 5)},
				{createCubie(0, 1, 2, 3, 4, 5), createCubie(0, 1, 2, 3, 4, 5), createCubie(0, 1, 2, 3, 4, 5)},
				{createCubie(0, 1, 2, 3, 4, 5), createCubie(0, 1, 2, 3, 4, 5), createCubie(0, 1, 2, 3, 4, 5)},
			},
			axis: UpAxis,
			want: Layer{
				{createCubie(1, 2, 3, 0, 4, 5), createCubie(1, 2, 3, 0, 4, 5), createCubie(1, 2, 3, 0, 4, 5)},
				{createCubie(1, 2, 3, 0, 4, 5), createCubie(1, 2, 3, 0, 4, 5), createCubie(1, 2, 3, 0, 4, 5)},
				{createCubie(1, 2, 3, 0, 4, 5), createCubie(1, 2, 3, 0, 4, 5), createCubie(1, 2, 3, 0, 4, 5)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Printf("Test %s, Axis: %v", tt.name, tt.axis)
			log.Printf("Layer before rotation:")
			// before rotation
			for i := range 3 {
				for j := range 3 {
					log.Printf("Before rotation(%d,%d): %v", i, j, tt.m[i][j].Colors)
				}
			}
			got := tt.m.rotateClockwise(tt.axis)
			// check all cubies colors
			for i := range 3 {
				for j := range 3 {
					if !reflect.DeepEqual(got[i][j].Colors, tt.want[i][j].Colors) {
						t.Errorf("Layer.rotateClockwise(%d,%d) = %v, want %v", i, j, got[i][j].Colors, tt.want[i][j].Colors)
					}
				}
			}
		})
	}
}
