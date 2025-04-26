package model

import (
	"reflect"
	"testing"
)

func Test_NewCube(t *testing.T) {
	tests := []struct {
		name string
		want *Cube
	}{
		{
			name: "New Cube",
			want: &Cube{
				State: [6]Face{
					// Front - white
					{
						Name:  "white",
						Index: Front,
						Stickers: [3][3]Sticker{
							{{Color: White, Index: Front_2_0_0},
								{Color: White, Index: Front_2_0_1},
								{Color: White, Index: Front_2_0_2}},
							{{Color: White, Index: Front_2_1_0},
								{Color: White, Index: Front_2_1_1},
								{Color: White, Index: Front_2_1_2}},
							{{Color: White, Index: Front_2_2_0},
								{Color: White, Index: Front_2_2_1},
								{Color: White, Index: Front_2_2_2}},
						},
					},
					// Back - yellow
					{
						Name:  "yellow",
						Index: Back,
						Stickers: [3][3]Sticker{
							{{Color: Yellow, Index: Back_0_0_0},
								{Color: Yellow, Index: Back_0_0_1},
								{Color: Yellow, Index: Back_0_0_2}},
							{{Color: Yellow, Index: Back_0_1_0},
								{Color: Yellow, Index: Back_0_1_1},
								{Color: Yellow, Index: Back_0_1_2}},
							{{Color: Yellow, Index: Back_0_2_0},
								{Color: Yellow, Index: Back_0_2_1},
								{Color: Yellow, Index: Back_0_2_2}},
						},
					},
					// Up - blue
					{
						Name:  "blue",
						Index: Up,
						Stickers: [3][3]Sticker{
							{{Color: Blue, Index: Up_0_0_2},
								{Color: Blue, Index: Up_0_1_2},
								{Color: Blue, Index: Up_0_2_2}},
							{{Color: Blue, Index: Up_1_0_2},
								{Color: Blue, Index: Up_1_1_2},
								{Color: Blue, Index: Up_1_2_2}},
							{{Color: Blue, Index: Up_2_0_2},
								{Color: Blue, Index: Up_2_1_2},
								{Color: Blue, Index: Up_2_2_2}},
						},
					},
					// Down - green
					{
						Name:  "green",
						Index: Down,
						Stickers: [3][3]Sticker{
							{{Color: Green, Index: Down_0_0_0},
								{Color: Green, Index: Down_0_1_0},
								{Color: Green, Index: Down_0_2_0}},
							{{Color: Green, Index: Down_1_0_0},
								{Color: Green, Index: Down_1_1_0},
								{Color: Green, Index: Down_1_2_0}},
							{{Color: Green, Index: Down_2_0_0},
								{Color: Green, Index: Down_2_1_0},
								{Color: Green, Index: Down_2_2_0}},
						},
					},
					// Left - red
					{
						Name:  "red",
						Index: Left,
						Stickers: [3][3]Sticker{
							{{Color: Red, Index: Left_0_0_0},
								{Color: Red, Index: Left_0_0_1},
								{Color: Red, Index: Left_0_0_2}},
							{{Color: Red, Index: Left_1_0_0},
								{Color: Red, Index: Left_1_0_1},
								{Color: Red, Index: Left_1_0_2}},
							{{Color: Red, Index: Left_2_0_0},
								{Color: Red, Index: Left_2_0_1},
								{Color: Red, Index: Left_2_0_2}},
						},
					},
					// Right - orange
					{
						Name:  "orange",
						Index: Right,
						Stickers: [3][3]Sticker{
							{{Color: Orange, Index: Right_0_2_0},
								{Color: Orange, Index: Right_0_2_1},
								{Color: Orange, Index: Right_0_2_2}},
							{{Color: Orange, Index: Right_1_2_0},
								{Color: Orange, Index: Right_1_2_1},
								{Color: Orange, Index: Right_1_2_2}},
							{{Color: Orange, Index: Right_2_2_0},
								{Color: Orange, Index: Right_2_2_1},
								{Color: Orange, Index: Right_2_2_2}},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCube(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got NewCube() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCube_RotateFront(t *testing.T) {
	tests := []struct {
		name      string
		clockwise TurningDirection
		expected  [6][3][3]Sticker
	}{
		{
			name:      "Front Face Clockwise Rotation",
			clockwise: Clockwise,
			expected: [6][3][3]Sticker{
				// Front - white (rotated clockwise)
				{
					{{Color: White, Index: Front_2_1_0},
						{Color: White, Index: Front_2_0_0},
						{Color: White, Index: Front_2_0_1}},
					{{Color: White, Index: Front_2_2_0},
						{Color: White, Index: Front_2_1_1},
						{Color: White, Index: Front_2_0_2}},
					{{Color: White, Index: Front_2_2_1},
						{Color: White, Index: Front_2_1_2},
						{Color: White, Index: Front_2_2_2}},
				},
				// Back - yellow (unchanged)
				{
					{{Color: Yellow, Index: Back_0_0_0},
						{Color: Yellow, Index: Back_0_0_1},
						{Color: Yellow, Index: Back_0_0_2}},
					{{Color: Yellow, Index: Back_0_1_0},
						{Color: Yellow, Index: Back_0_1_1},
						{Color: Yellow, Index: Back_0_1_2}},
					{{Color: Yellow, Index: Back_0_2_0},
						{Color: Yellow, Index: Back_0_2_1},
						{Color: Yellow, Index: Back_0_2_2}},
				},
				// Up - blue with bottom row changed
				{
					{{Color: Blue, Index: Up_0_0_2},
						{Color: Blue, Index: Up_0_1_2},
						{Color: Blue, Index: Up_0_2_2}},
					{{Color: Blue, Index: Up_1_0_2},
						{Color: Blue, Index: Up_1_1_2},
						{Color: Blue, Index: Up_1_2_2}},
					{{Color: Red, Index: Left_2_0_2},
						{Color: Red, Index: Left_2_0_1},
						{Color: Red, Index: Left_2_0_0}},
				},
				// Down - green with top row changed
				{
					{{Color: Orange, Index: Right_0_2_0},
						{Color: Orange, Index: Right_1_2_0},
						{Color: Orange, Index: Right_2_2_0}},
					{{Color: Green, Index: Down_1_0_0},
						{Color: Green, Index: Down_1_1_0},
						{Color: Green, Index: Down_1_2_0}},
					{{Color: Green, Index: Down_2_0_0},
						{Color: Green, Index: Down_2_1_0},
						{Color: Green, Index: Down_2_2_0}},
				},
				// Left - red with right edge changed
				{
					{{Color: Red, Index: Left_0_0_0},
						{Color: Red, Index: Left_0_0_1},
						{Color: Green, Index: Down_0_0_0}},
					{{Color: Red, Index: Left_1_0_0},
						{Color: Red, Index: Left_1_0_1},
						{Color: Green, Index: Down_0_1_0}},
					{{Color: Red, Index: Left_2_0_0},
						{Color: Red, Index: Left_2_0_1},
						{Color: Green, Index: Down_0_2_0}},
				},
				// Right - orange with left column changed
				{
					{{Color: Blue, Index: Up_2_0_2},
						{Color: Orange, Index: Right_0_2_1},
						{Color: Orange, Index: Right_0_2_2}},
					{{Color: Blue, Index: Up_2_1_2},
						{Color: Orange, Index: Right_1_2_1},
						{Color: Orange, Index: Right_1_2_2}},
					{{Color: Blue, Index: Up_2_2_2},
						{Color: Orange, Index: Right_2_2_1},
						{Color: Orange, Index: Right_2_2_2}},
				},
			},
		},
		{
			name:      "Front Face Counter-Clockwise Rotation",
			clockwise: CounterClockwise,
			expected: [6][3][3]Sticker{
				// Front - white (rotated counter-clockwise)
				{
					{{Color: White, Index: Front_2_0_2},
						{Color: White, Index: Front_2_1_2},
						{Color: White, Index: Front_2_2_2}},
					{{Color: White, Index: Front_2_0_1},
						{Color: White, Index: Front_2_1_1},
						{Color: White, Index: Front_2_2_1}},
					{{Color: White, Index: Front_2_0_0},
						{Color: White, Index: Front_2_1_0},
						{Color: White, Index: Front_2_2_0}},
				},
				// Back - yellow (unchanged)
				{
					{{Color: Yellow, Index: Back_0_0_0},
						{Color: Yellow, Index: Back_0_0_1},
						{Color: Yellow, Index: Back_0_0_2}},
					{{Color: Yellow, Index: Back_0_1_0},
						{Color: Yellow, Index: Back_0_1_1},
						{Color: Yellow, Index: Back_0_1_2}},
					{{Color: Yellow, Index: Back_0_2_0},
						{Color: Yellow, Index: Back_0_2_1},
						{Color: Yellow, Index: Back_0_2_2}},
				},
				// Up - blue with bottom row changed
				{
					{{Color: Blue, Index: Up_0_0_2},
						{Color: Blue, Index: Up_0_1_2},
						{Color: Blue, Index: Up_0_2_2}},
					{{Color: Blue, Index: Up_1_0_2},
						{Color: Blue, Index: Up_1_1_2},
						{Color: Blue, Index: Up_1_2_2}},
					{{Color: Orange, Index: Right_0_2_0},
						{Color: Orange, Index: Right_1_2_0},
						{Color: Orange, Index: Right_2_2_0}},
				},
				// Down - green with top row changed
				{
					{{Color: Red, Index: Left_0_0_2},
						{Color: Red, Index: Left_1_0_2},
						{Color: Red, Index: Left_2_0_2}},
					{{Color: Green, Index: Down_1_0_0},
						{Color: Green, Index: Down_1_1_0},
						{Color: Green, Index: Down_1_2_0}},
					{{Color: Green, Index: Down_2_0_0},
						{Color: Green, Index: Down_2_1_0},
						{Color: Green, Index: Down_2_2_0}},
				},
				// Left - red with right column changed
				{
					{{Color: Red, Index: Left_0_0_0},
						{Color: Red, Index: Left_0_0_1},
						{Color: Blue, Index: Up_2_0_2}},
					{{Color: Red, Index: Left_1_0_0},
						{Color: Red, Index: Left_1_0_1},
						{Color: Blue, Index: Up_2_1_2}},
					{{Color: Red, Index: Left_2_0_0},
						{Color: Red, Index: Left_2_0_1},
						{Color: Blue, Index: Up_2_2_2}},
				},
				// Right - orange with left column changed
				{
					{{Color: Green, Index: Down_0_0_0},
						{Color: Orange, Index: Right_0_2_1},
						{Color: Orange, Index: Right_0_2_2}},
					{{Color: Green, Index: Down_0_1_0},
						{Color: Orange, Index: Right_1_2_1},
						{Color: Orange, Index: Right_1_2_2}},
					{{Color: Green, Index: Down_0_2_0},
						{Color: Orange, Index: Right_2_2_1},
						{Color: Orange, Index: Right_2_2_2}},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize a fresh cube
			c := NewCube()

			// Apply rotation
			c.RotateFace(Front, tt.clockwise)

			// Verify the adjacent faces match expectations
			// We check each face separately to make debugging easier
			faces := []string{"Front", "Back", "Up", "Down", "Left", "Right"}
			for i := range 6 {
				// Check the stickers of the face
				if !reflect.DeepEqual(c.State[i].Stickers, tt.expected[i]) {
					t.Errorf("%s face after rotation: got %v, want %v",
						faces[i], c.State[i].Stickers, tt.expected[i])
				}
			}
		})
	}
}

// Test full cube rotation consistency
func TestCube_RotationConsistency(t *testing.T) {
	c := NewCube()

	// Perform a clockwise front rotation followed by counter-clockwise front rotation
	// The cube should return to its original state
	originalState := c.State

	// Apply clockwise rotation
	c.RotateFace(Front, Clockwise)
	// Apply counter-clockwise rotation
	c.RotateFace(Front, CounterClockwise)

	// Verify the cube returned to original state
	if !reflect.DeepEqual(c.State, originalState) {
		t.Errorf("Cube did not return to original state after CW+CCW rotations")
	}

	// Test four clockwise rotations (should return to original state)
	c = NewCube()
	originalState = c.State

	for i := 0; i < 4; i++ {
		c.RotateFace(Front, Clockwise)
	}

	if !reflect.DeepEqual(c.State, originalState) {
		t.Errorf("Cube did not return to original state after 4 clockwise rotations")
	}
}

func TestCube_RotateFrontDebug(t *testing.T) {
	// Create a cube and rotate the front face clockwise
	c := NewCube()

	// Get the initial state of the edges around the front face
	up, uborder := GetNorthFace(Front)
	upEdge := c.GetEdge(up, uborder)
	right, rborder := GetEastFace(Front)
	rightEdge := c.GetEdge(right, rborder)
	down, dborder := GetSouthFace(Front)
	downEdge := c.GetEdge(down, dborder)
	left, lborder := GetWestFace(Front)
	leftEdge := c.GetEdge(left, lborder)

	t.Logf("Before rotation:")
	t.Logf("  Up (%v) edge: %v %v %v", uborder, upEdge[0], upEdge[1], upEdge[2])
	t.Logf("  Right (%v) edge: %v %v %v", rborder, rightEdge[0], rightEdge[1], rightEdge[2])
	t.Logf("  Down (%v) edge: %v %v %v", dborder, downEdge[0], downEdge[1], downEdge[2])
	t.Logf("  Left (%v) edge: %v %v %v", lborder, leftEdge[0], leftEdge[1], leftEdge[2])

	// Rotate the front face clockwise
	c.RotateFace(Front, Clockwise)

	// Get the state after rotation
	upEdgeAfter := c.GetEdge(up, uborder)
	rightEdgeAfter := c.GetEdge(right, rborder)
	downEdgeAfter := c.GetEdge(down, dborder)
	leftEdgeAfter := c.GetEdge(left, lborder)

	t.Logf("After clockwise rotation:")
	t.Logf("  Up (%v) edge: %v %v %v", uborder, upEdgeAfter[0], upEdgeAfter[1], upEdgeAfter[2])
	t.Logf("  Right (%v) edge: %v %v %v", rborder, rightEdgeAfter[0], rightEdgeAfter[1], rightEdgeAfter[2])
	t.Logf("  Down (%v) edge: %v %v %v", dborder, downEdgeAfter[0], downEdgeAfter[1], downEdgeAfter[2])
	t.Logf("  Left (%v) edge: %v %v %v", lborder, leftEdgeAfter[0], leftEdgeAfter[1], leftEdgeAfter[2])
}
