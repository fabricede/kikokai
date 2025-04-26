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
						Name:  "White",
						Index: Front,
						Stickers: [3][3]Sticker{
							{{Color: White, Index: Front_2_2_0},
								{Color: White, Index: Front_2_2_1},
								{Color: White, Index: Front_2_2_2}},
							{{Color: White, Index: Front_2_1_0},
								{Color: White, Index: Front_2_1_1},
								{Color: White, Index: Front_2_1_2}},
							{{Color: White, Index: Front_2_0_0},
								{Color: White, Index: Front_2_0_1},
								{Color: White, Index: Front_2_0_2}},
						},
					},
					// Back - yellow
					{
						Name:  "Yellow",
						Index: Back,
						Stickers: [3][3]Sticker{
							{{Color: Yellow, Index: Back_NW},
								{Color: Yellow, Index: Back_N},
								{Color: Yellow, Index: Back_NE}},
							{{Color: Yellow, Index: Back_W},
								{Color: Yellow, Index: Back_Center},
								{Color: Yellow, Index: Back_E}},
							{{Color: Yellow, Index: Back_SW},
								{Color: Yellow, Index: Back_S},
								{Color: Yellow, Index: Back_SE}},
						},
					},
					// Up - blue
					{
						Name:  "Blue",
						Index: Up,
						Stickers: [3][3]Sticker{
							{{Color: Blue, Index: Up_NW},
								{Color: Blue, Index: Up_N},
								{Color: Blue, Index: Up_NE}},
							{{Color: Blue, Index: Up_W},
								{Color: Blue, Index: Up_Center},
								{Color: Blue, Index: Up_E}},
							{{Color: Blue, Index: Up_SW},
								{Color: Blue, Index: Up_S},
								{Color: Blue, Index: Up_SE}},
						},
					},
					// Down - green
					{
						Name:  "Green",
						Index: Down,
						Stickers: [3][3]Sticker{
							{{Color: Green, Index: Down_NW},
								{Color: Green, Index: Down_N},
								{Color: Green, Index: Down_NE}},
							{{Color: Green, Index: Down_W},
								{Color: Green, Index: Down_Center},
								{Color: Green, Index: Down_E}},
							{{Color: Green, Index: Down_SW},
								{Color: Green, Index: Down_S},
								{Color: Green, Index: Down_SE}},
						},
					},
					// Left - red
					{
						Name:  "Red",
						Index: Left,
						Stickers: [3][3]Sticker{
							{{Color: Red, Index: Left_NW},
								{Color: Red, Index: Left_N},
								{Color: Red, Index: Left_NE}},
							{{Color: Red, Index: Left_W},
								{Color: Red, Index: Left_Center},
								{Color: Red, Index: Left_E}},
							{{Color: Red, Index: Left_SW},
								{Color: Red, Index: Left_S},
								{Color: Red, Index: Left_SE}},
						},
					},
					// Right - orange
					{
						Name:  "Orange",
						Index: Right,
						Stickers: [3][3]Sticker{
							{{Color: Orange, Index: Right_NW},
								{Color: Orange, Index: Right_N},
								{Color: Orange, Index: Right_NE}},
							{{Color: Orange, Index: Right_W},
								{Color: Orange, Index: Right_Center},
								{Color: Orange, Index: Right_E}},
							{{Color: Orange, Index: Right_SW},
								{Color: Orange, Index: Right_S},
								{Color: Orange, Index: Right_SE}},
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
					{{Color: White, Index: Front_2_0_0},
						{Color: White, Index: Front_2_1_0},
						{Color: White, Index: Front_2_2_0}},
					{{Color: White, Index: Front_2_0_1},
						{Color: White, Index: Front_2_1_1},
						{Color: White, Index: Front_2_2_1}},
					{{Color: White, Index: Front_2_0_2},
						{Color: White, Index: Front_2_1_2},
						{Color: White, Index: Front_2_2_2}},
				},
				// Back - yellow (unchanged)
				{
					{{Color: Yellow, Index: Back_NW},
						{Color: Yellow, Index: Back_N},
						{Color: Yellow, Index: Back_NE}},
					{{Color: Yellow, Index: Back_W},
						{Color: Yellow, Index: Back_Center},
						{Color: Yellow, Index: Back_E}},
					{{Color: Yellow, Index: Back_SW},
						{Color: Yellow, Index: Back_S},
						{Color: Yellow, Index: Back_SE}},
				},
				// Up - blue with Left column changed Red From Left south
				{
					{{Color: Red, Index: Left_SE},
						{Color: Blue, Index: Up_N},
						{Color: Blue, Index: Up_NE}},
					{{Color: Red, Index: Left_S},
						{Color: Blue, Index: Up_Center},
						{Color: Blue, Index: Up_E}},
					{{Color: Red, Index: Left_SW},
						{Color: Blue, Index: Up_S},
						{Color: Blue, Index: Up_SE}},
				},
				// Down - green with top row changed Orange From Right north
				{
					{{Color: Orange, Index: Right_SW},
						{Color: Green, Index: Down_N},
						{Color: Green, Index: Down_NE}},
					{{Color: Orange, Index: Right_W},
						{Color: Green, Index: Down_Center},
						{Color: Green, Index: Down_E}},
					{{Color: Orange, Index: Right_NW},
						{Color: Green, Index: Down_S},
						{Color: Green, Index: Down_SE}},
				},
				// Left - red with bottom row changed Green From Down west
				{
					{{Color: Red, Index: Left_NW},
						{Color: Red, Index: Left_N},
						{Color: Red, Index: Left_NE}},
					{{Color: Red, Index: Left_W},
						{Color: Red, Index: Left_Center},
						{Color: Red, Index: Left_E}},
					{{Color: Green, Index: Down_NW},
						{Color: Green, Index: Down_W},
						{Color: Green, Index: Down_SW}},
				},
				// Right - orange with Up row changed Blue From Up west
				{
					{{Color: Blue, Index: Up_NW},
						{Color: Blue, Index: Up_W},
						{Color: Blue, Index: Up_SW}},
					{{Color: Orange, Index: Right_W},
						{Color: Orange, Index: Right_Center},
						{Color: Orange, Index: Right_E}},
					{{Color: Orange, Index: Right_SW},
						{Color: Orange, Index: Right_S},
						{Color: Orange, Index: Right_SE}},
				},
			},
		},
		/*{
			name:      "Front Face Counter-Clockwise Rotation",
			clockwise: CounterClockwise,
			expected: [6][3][3]Color{
				// Front - white (rotated counter-clockwise)
				{
					{White, White, White},
					{White, White, White},
					{White, White, White},
				},
				// Back - yellow (unchanged)
				{
					{Yellow, Yellow, Yellow},
					{Yellow, Yellow, Yellow},
					{Yellow, Yellow, Yellow},
				},
				// Up - blue with bottom row changed
				{
					{Blue, Blue, Blue},
					{Blue, Blue, Blue},
					{Orange, Orange, Orange}, // From Right
				},
				// Down - green with top row changed
				{
					{Red, Red, Red}, // From Left
					{Green, Green, Green},
					{Green, Green, Green},
				},
				// Left - red with right column changed
				{
					{Red, Red, Blue},
					{Red, Red, Blue},
					{Red, Red, Blue}, // From Up
				},
				// Right - orange with left column changed
				{
					{Green, Orange, Orange},
					{Green, Orange, Orange},
					{Green, Orange, Orange}, // From Down
				},
			},
		},*/
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
