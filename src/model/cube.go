package model

import (
	"math/rand"
	"time"
)

// Cube represents a Rubik's cube with 6 faces, each face is a 3x3 grid
type Cube struct {
	State [6][3][3]Color // 6 faces, 3x3 grid per face, using Color enum instead of string
}

// NewCube creates and initializes a new Rubik's cube
func NewCube() *Cube {
	c := &Cube{}
	colors := [6]Color{White, Yellow, Blue, Green, Red, Orange}
	for i, color := range colors {
		for row := 0; row < 3; row++ {
			for col := 0; col < 3; col++ {
				c.State[i][row][col] = color
			}
		}
	}
	return c
}

// RotateFace rotates the specified face
func (c *Cube) RotateFace(face Face, clockwise Direction) {
	RotateFace(c, face, clockwise)
}

// Scramble applies a series of random rotations to the cube
func (c *Cube) Scramble(moves int) {
	// Seed the random number generator if not seeded already
	rand.Seed(time.Now().UnixNano())

	// Apply random rotations
	for i := 0; i < moves; i++ {
		// Random face (0-5)
		face := Face(rand.Intn(6))

		// Random direction (true/false for clockwise/counter-clockwise)
		clockwise := Direction(rand.Intn(2) == 1)

		// Rotate the face
		c.RotateFace(face, clockwise)
	}
}
