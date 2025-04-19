package model

import (
	"math/rand"
)

type Sticker struct {
	Color Color
	Index StickerIndex
}

func (s *Sticker) GetColor() Color {
	return s.Color
}
func (s *Sticker) GetIndex() StickerIndex {
	return s.Index
}
func (s *Sticker) GetName() string {
	return StickerColorName[s.Index]
}

type Face struct {
	Name     string
	Index    FaceIndex
	Stickers [3][3]Sticker
}

func (f *Face) GetName() string {
	return f.Name
}
func (f *Face) GetIndex() FaceIndex {
	return f.Index
}
func (f *Face) GetColor() Color {
	return f.Stickers[1][1].GetColor()
}
func GetFaceNameFromIndex(i FaceIndex) string {
	return []string{"Front", "Back", "Up", "Down", "Left", "Right"}[i]
}

// Cube represents a Rubik's cube with 6 faces
type Cube struct {
	State [6]Face // 6 faces
}

// NewCube creates and initializes a new Rubik's cube
func NewCube() *Cube {
	c := &Cube{}
	for i, face := range c.State {
		face.Index = FaceIndex(i)
		face.Name = GetFaceNameFromIndex(face.Index)
		stickers := [3][3]Sticker{}
		// Initialize stickers for each face
		stickerIndex := 0
		for row := 0; row < 3; row++ {
			for col := 0; col < 3; col++ {
				face.Stickers[row][col] = Sticker{
					Color: Color(i),
					Index: StickerIndex(stickerIndex),
				}
				stickerIndex++
			}
		}
		face.Stickers = stickers
		c.State[i] = face
	}
	return c
}

// RotateFace rotates the specified face
func (c *Cube) RotateFace(face FaceIndex, clockwise TurningDirection) {
	RotateFace(c, face, clockwise)
}

// GetEdge returns the edge stickers of the adjacent face in the given orientation
func (c *Cube) GetEdge(face FaceIndex, orientation Orientation) (edge [3]Sticker) {
	switch orientation {
	case North:
		edge[0] = c.State[face].Stickers[0][0]
		edge[1] = c.State[face].Stickers[1][0]
		edge[2] = c.State[face].Stickers[2][0]
	case South:
		edge[0] = c.State[face].Stickers[2][0]
		edge[1] = c.State[face].Stickers[2][1]
		edge[2] = c.State[face].Stickers[2][2]
	case East:
		edge[0] = c.State[face].Stickers[0][2]
		edge[1] = c.State[face].Stickers[1][2]
		edge[2] = c.State[face].Stickers[2][2]
	case West:
		edge[0] = c.State[face].Stickers[0][0]
		edge[1] = c.State[face].Stickers[1][0]
		edge[2] = c.State[face].Stickers[2][0]
	default:
		edge[0] = c.State[face].Stickers[1][1]
		edge[1] = c.State[face].Stickers[1][1]
		edge[2] = c.State[face].Stickers[1][1]
	}
	return edge
}

// SetEdge sets the edge stickers of the adjacent face in the given orientation
func (cube *Cube) SetEdge(face FaceIndex, orientation Orientation, edge [3]Sticker) {
	// copy the stickers from the cube
	stickers := cube.State[face].Stickers
	switch orientation {
	case North:
		stickers[0][0] = edge[0]
		stickers[1][0] = edge[1]
		stickers[2][0] = edge[2]
	case South:
		stickers[2][0] = edge[0]
		stickers[2][1] = edge[1]
		stickers[2][2] = edge[2]
	case East:
		stickers[0][2] = edge[0]
		stickers[1][2] = edge[1]
		stickers[2][2] = edge[2]
	case West:
		stickers[0][0] = edge[0]
		stickers[1][0] = edge[1]
		stickers[2][0] = edge[2]
	}
	cube.State[face].Stickers = stickers
}

// Scramble applies a series of random rotations to the cube
func (c *Cube) Scramble(moves int) {
	// Apply random rotations
	for i := 0; i < moves; i++ {
		// Random face (0-5)
		face := FaceIndex(rand.Intn(6))

		// Random direction (true/false for clockwise/counter-clockwise)
		clockwise := TurningDirection(rand.Intn(2) == 1)

		// Rotate the face
		c.RotateFace(face, clockwise)
	}
}
