package model

import (
	"reflect"
	"testing"
)

func TestCube_rotateFrontAdjacents(t *testing.T) {
	type fields struct {
		State [6][3][3]Color
	}
	type args struct {
		clockwise Direction
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected [6][3][3]Color
	}{
		{
			name: "Front Face Clockwise",
			fields: fields{
				State: [6][3][3]Color{
					// Front - white
					{
						{White, White, White},
						{White, White, White},
						{White, White, White},
					},
					// Back - yellow
					{
						{Yellow, Yellow, Yellow},
						{Yellow, Yellow, Yellow},
						{Yellow, Yellow, Yellow},
					},
					// Up - blue
					{
						{Blue, Blue, Blue},
						{Blue, Blue, Blue},
						{Blue, Blue, Blue},
					},
					// Down - green
					{
						{Green, Green, Green},
						{Green, Green, Green},
						{Green, Green, Green},
					},
					// Left - red
					{
						{Red, Red, Red},
						{Red, Red, Red},
						{Red, Red, Red},
					},
					// Right - orange
					{
						{Orange, Orange, Orange},
						{Orange, Orange, Orange},
						{Orange, Orange, Orange},
					},
				},
			},
			args: args{clockwise: Clockwise},
			expected: [6][3][3]Color{
				// Front - white (unchanged)
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
					{Red, Red, Red}, // Updated from Left
					{Blue, Blue, Blue},
					{Blue, Blue, Blue},
				},
				// Down - green with top row changed
				{
					{Green, Green, Green}, // Updated from Right
					{Green, Green, Green},
					{Orange, Orange, Orange},
				},
				// Left - red with right column changed
				{
					{Green, Red, Red},
					{Green, Red, Red},
					{Green, Red, Red}, // Updated from Down
				},
				// Right - orange with left column changed
				{
					{Orange, Orange, Blue},
					{Orange, Orange, Blue},
					{Orange, Orange, Blue}, // Updated from Up
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cube{
				State: tt.fields.State,
			}
			c.rotateFrontAdjacents(tt.args.clockwise)

			// Compare the result with expected state
			if !reflect.DeepEqual(c.State, tt.expected) {
				t.Errorf("rotateFrontAdjacents() got = %v, want %v", c.State, tt.expected)
			}
		})
	}
}

func TestCube_RotateFront(t *testing.T) {
	tests := []struct {
		name      string
		clockwise Direction
		expected  [6][3][3]Color
	}{
		{
			name:      "Front Face Clockwise Rotation",
			clockwise: Clockwise,
			expected: [6][3][3]Color{
				// Front - white (rotated clockwise)
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
					{Red, Red, Red}, // From Left
				},
				// Down - green with top row changed
				{
					{Orange, Orange, Orange}, // From Right
					{Green, Green, Green},
					{Green, Green, Green},
				},
				// Left - red with right column changed
				{
					{Red, Red, Green},
					{Red, Red, Green},
					{Red, Red, Green}, // From Down
				},
				// Right - orange with left column changed
				{
					{Blue, Orange, Orange},
					{Blue, Orange, Orange},
					{Blue, Orange, Orange}, // From Up
				},
			},
		},
		{
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
			for i := 0; i < 6; i++ {
				if !reflect.DeepEqual(c.State[i], tt.expected[i]) {
					t.Errorf("%s face after rotation: got %v, want %v",
						faces[i], c.State[i], tt.expected[i])
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
