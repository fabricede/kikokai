package model

import (
	"log"
	"math/rand"
)

// SharedCube is the global cube instance to be shared across servers
var SharedCube *Cube
var initialCube *Cube

func init() {
	// Initialize a new cube
	SharedCube = NewCube()
	initialCube = NewCube() // used to get stickers position
}

// ResetCube resets the cube to its initial state
func ResetCube() {
	SharedCube = NewCube()
}

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
	return FaceColorName[i]
}

// Cube represents a Rubik's cube with 6 faces
type Cube struct {
	State [6]Face // 6 faces
}

// NewCube creates and initializes a new Rubik's cube
func NewCube() *Cube {
	c := &Cube{}
	for i := range c.State {
		face := &c.State[i]
		face.Index = FaceIndex(i)
		face.Name = GetFaceNameFromIndex(face.Index)

		// Initialize stickers for each face
		stickerCount := 0
		for row := range 3 {
			for col := range 3 {
				currentStickerIndex := StickerIndex(i*9 + stickerCount)
				colorName := StickerColorName[currentStickerIndex]
				log.Printf("Initializing sticker %d,%d on face %s : color %s", row, col, face.Name, colorName)

				// Assign colors based on the face index
				face.Stickers[row][col] = Sticker{
					Color: Color(i),            // Use the face index as the color
					Index: currentStickerIndex, // Use the calculated sticker index
				}
				stickerCount++
			}
		}
	}
	return c
}

// GetSttickerIndex returns the sticker index of the specified face and position
func (c *Cube) GetStickerIndex(face FaceIndex, row, col int) StickerIndex {
	if face < 0 || face >= 6 {
		log.Println("Invalid face index")
		return -1
	}
	if row < 0 || row >= 3 || col < 0 || col >= 3 {
		log.Println("Invalid row or column index")
		return -1
	}
	return c.State[face].Stickers[row][col].Index
}

// GetSticker returns the sticker at the given 3D position
func (cube *Cube) GetSticker(face FaceIndex, position CubeCoordinate) (Sticker, error) {
	// Get the row and column on that face
	row, col, err := GetStickerCoordinate(face, position)
	if err != nil {
		return Sticker{}, err
	}
	// Return the sticker at that position
	return cube.State[face].Stickers[row][col], nil
}

// SetSticker sets the sticker at the given 3D position
func (cube *Cube) SetSticker(face FaceIndex, position CubeCoordinate, sticker Sticker) error {
	// Get the row and column on that face
	row, col, err := GetStickerCoordinate(face, position)
	if err != nil {
		return err
	}
	// Set the sticker at that position
	cube.State[face].Stickers[row][col] = sticker
	return nil
}

// RotateFace rotates the specified face
func (c *Cube) RotateFace(face FaceIndex, clockwise TurningDirection) {
	RotateFace(c, face, clockwise)
}

// GetEdge returns the edge stickers of the face in the given axis in cubecoordinate{(-1/+1),0,0}
func (c *Cube) GetEdge(face FaceIndex, axis CubeCoordinate) (edge [3]Sticker) {
	faceCoord := FaceToCoordinate(face)
	x := faceCoord.X | axis.X
	y := faceCoord.Y | axis.Y
	z := faceCoord.Z | axis.Z
	switch {
	case x == 0:
		for i := range 3 {
			edge[i], _ = c.GetSticker(face, CubeCoordinate{X: i - 1, Y: y, Z: z})
		}
	case y == 0:
		for i := range 3 {
			edge[i], _ = c.GetSticker(face, CubeCoordinate{X: x, Y: i - 1, Z: z})
		}
	case z == 0:
		for i := range 3 {
			edge[i], _ = c.GetSticker(face, CubeCoordinate{X: x, Y: y, Z: i - 1})
		}
	}
	return edge
}

// SetEdge sets the edge stickers of the face in the given axis in cubecoordinate{(-1/+1),0,0}
func (c *Cube) SetEdge(face FaceIndex, axis CubeCoordinate, edge [3]Sticker) {
	faceCoord := FaceToCoordinate(face)
	x := faceCoord.X | axis.X
	y := faceCoord.Y | axis.Y
	z := faceCoord.Z | axis.Z
	switch {
	case x == 0:
		for i := range 3 {
			c.SetSticker(face, CubeCoordinate{X: i - 1, Y: y, Z: z}, edge[i])
		}
	case y == 0:
		for i := range 3 {
			c.SetSticker(face, CubeCoordinate{X: x, Y: i - 1, Z: z}, edge[i])
		}
	case z == 0:
		for i := range 3 {
			c.SetSticker(face, CubeCoordinate{X: x, Y: y, Z: i - 1}, edge[i])
		}
	}
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
