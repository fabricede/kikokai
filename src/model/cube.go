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

// ------------------------------------------------
// Sticker represents a sticker on the Rubik's cube
// ------------------------------------------------
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

// ------------------------------------------
// Face represents a face of the Rubik's cube
// ------------------------------------------
type Face struct {
	Index    FaceIndex
	Stickers [3][3]Sticker
}

func (f *Face) GetName() string {
	return FaceColorName[f.Index]
}
func (f *Face) GetIndex() FaceIndex {
	return f.Index
}
func (f *Face) GetCubeCoordinate(row, col int) CubeCoordinate {
	faceCoordinate := FaceToCoordinate(f.Index)

	// Adjust coordinates based on the face's axis alignment
	switch {
	case faceCoordinate.X != 0: // Face is aligned along the X-axis
		if faceCoordinate.X == 1 { // Front face YZ
			return CubeCoordinate{
				X: faceCoordinate.X, // X remains constant
				Y: row - 1,          // Adjust Y based on column
				Z: col - 1,          // Adjust Z based on row
			}
		} else {
			return CubeCoordinate{
				X: faceCoordinate.X, // X remains constant
				Y: col - 1,          // Adjust Y based on column
				Z: row - 1,          // Adjust Z based on row
			}
		}
	case faceCoordinate.Y != 0: // Face is aligned along the Y-axis
		return CubeCoordinate{
			X: col - 1,          // Adjust X based on column
			Y: faceCoordinate.Y, // Y remains constant
			Z: row - 1,          // Adjust Z based on row
		}
	case faceCoordinate.Z != 0: // Face is aligned along the Z-axis
		return CubeCoordinate{
			X: col - 1,          // Adjust X based on column
			Y: row - 1,          // Adjust Y based on row
			Z: faceCoordinate.Z, // Z remains constant
		}
	default:
		log.Println("Invalid face alignment")
		return CubeCoordinate{}
	}
}
func (f *Face) GetFaceCoordinate(coord CubeCoordinate) (row, col int) {
	faceCoordinate := FaceToCoordinate(f.Index)

	// Adjust coordinates based on the face's axis alignment
	switch {
	case faceCoordinate.X != 0 && coord.X == faceCoordinate.X: // Face is aligned along the X-axis
		if faceCoordinate.X == 1 { // Front face YZ
			row = coord.Y + 1 // Adjust row based on Y
			col = coord.Z + 1 // Adjust column based on Z
		} else {
			row = 1 + coord.Z // Adjust row based on Z
			col = 1 + coord.Y // Adjust column based on Y
		}
	case faceCoordinate.Y != 0 && coord.Y == faceCoordinate.Y: // Face is aligned along the Y-axis
		row = 1 + coord.Z // Adjust row based on Z
		col = coord.X + 1 // Adjust column based on X
	case faceCoordinate.Z != 0 && coord.Z == faceCoordinate.Z: // Face is aligned along the Z-axis
		row = 1 + coord.Y // Adjust row based on Y
		col = coord.X + 1 // Adjust column based on X
	default:
		log.Println("Invalid face alignment")
	}
	return row, col
}

// -------------------------------------------
// Cube represents a Rubik's cube with 6 faces
// -------------------------------------------
type Cube struct {
	State [6]Face // 6 faces
}

// NewCube creates and initializes a new Rubik's cube
func NewCube() *Cube {
	c := &Cube{}

	// Initialize each face
	for i := range c.State {
		face := &c.State[i]
		face.Index = FaceIndex(i)
		// Initialize stickers for each face
		for row := range 3 {
			for col := range 3 {
				// Determine the sticker index based on its position
				var stickerIdx StickerIndex

				// For corners and edges, use predefined indices that match across faces
				switch face.Index {
				case Front:
					switch {
					case row == 0 && col == 0:
						stickerIdx = Front_2_0_0
					case row == 0 && col == 1:
						stickerIdx = Front_2_0_1
					case row == 0 && col == 2:
						stickerIdx = Front_2_0_2
					case row == 1 && col == 0:
						stickerIdx = Front_2_1_0
					case row == 1 && col == 1:
						stickerIdx = Front_2_1_1
					case row == 1 && col == 2:
						stickerIdx = Front_2_1_2
					case row == 2 && col == 0:
						stickerIdx = Front_2_2_0
					case row == 2 && col == 1:
						stickerIdx = Front_2_2_1
					case row == 2 && col == 2:
						stickerIdx = Front_2_2_2
					}
				case Back:
					switch {
					case row == 0 && col == 0:
						stickerIdx = Back_0_0_0
					case row == 0 && col == 1:
						stickerIdx = Back_0_1_0
					case row == 0 && col == 2:
						stickerIdx = Back_0_2_0
					case row == 1 && col == 0:
						stickerIdx = Back_0_0_1
					case row == 1 && col == 1:
						stickerIdx = Back_0_1_1
					case row == 1 && col == 2:
						stickerIdx = Back_0_2_1
					case row == 2 && col == 0:
						stickerIdx = Back_0_0_2
					case row == 2 && col == 1:
						stickerIdx = Back_0_1_2
					case row == 2 && col == 2:
						stickerIdx = Back_0_2_2
					}
				case Up:
					switch {
					case row == 0 && col == 0:
						stickerIdx = Up_0_0_2
					case row == 0 && col == 1:
						stickerIdx = Up_1_0_2
					case row == 0 && col == 2:
						stickerIdx = Up_2_0_2
					case row == 1 && col == 0:
						stickerIdx = Up_0_1_2
					case row == 1 && col == 1:
						stickerIdx = Up_1_1_2
					case row == 1 && col == 2:
						stickerIdx = Up_1_1_2
					case row == 2 && col == 0:
						stickerIdx = Up_0_2_2
					case row == 2 && col == 1:
						stickerIdx = Up_1_2_2
					case row == 2 && col == 2:
						stickerIdx = Up_2_2_2
					}
				case Down:
					switch {
					case row == 0 && col == 0:
						stickerIdx = Down_0_0_0
					case row == 0 && col == 1:
						stickerIdx = Down_1_0_0
					case row == 0 && col == 2:
						stickerIdx = Down_2_0_0
					case row == 1 && col == 0:
						stickerIdx = Down_0_1_0
					case row == 1 && col == 1:
						stickerIdx = Down_1_1_0
					case row == 1 && col == 2:
						stickerIdx = Down_2_1_0
					case row == 2 && col == 0:
						stickerIdx = Down_0_2_0
					case row == 2 && col == 1:
						stickerIdx = Down_1_2_0
					case row == 2 && col == 2:
						stickerIdx = Down_2_2_0
					}
				case Left:
					switch {
					case row == 0 && col == 0:
						stickerIdx = Left_0_0_0
					case row == 0 && col == 1:
						stickerIdx = Left_0_0_1
					case row == 0 && col == 2:
						stickerIdx = Left_0_0_1
					case row == 1 && col == 0:
						stickerIdx = Left_1_0_1
					case row == 1 && col == 1:
						stickerIdx = Left_1_0_1
					case row == 1 && col == 2:
						stickerIdx = Left_1_0_1
					case row == 2 && col == 0:
						stickerIdx = Left_2_0_0
					case row == 2 && col == 1:
						stickerIdx = Left_2_0_1
					case row == 2 && col == 2:
						stickerIdx = Left_2_0_2
					}
				case Right:
					switch {
					case row == 0 && col == 0:
						stickerIdx = Right_0_2_0
					case row == 0 && col == 1:
						stickerIdx = Right_1_2_0
					case row == 0 && col == 2:
						stickerIdx = Right_2_2_0
					case row == 1 && col == 0:
						stickerIdx = Right_0_2_1
					case row == 1 && col == 1:
						stickerIdx = Right_1_2_1
					case row == 1 && col == 2:
						stickerIdx = Right_2_2_1
					case row == 2 && col == 0:
						stickerIdx = Right_0_2_2
					case row == 2 && col == 1:
						stickerIdx = Right_1_2_2
					case row == 2 && col == 2:
						stickerIdx = Right_2_2_2
					}
				}

				// If no specific index was assigned (shouldn't happen), fall back to a calculated one
				if stickerIdx == 0 && face.Index != 0 {
					// Use face index * 9 + position in face as fallback
					stickerIdx = StickerIndex(int(face.Index)*9 + row*3 + col)
				}

				// Create the sticker with the appropriate color and index
				face.Stickers[row][col] = Sticker{
					Color: Color(face.Index), // Color matches the face
					Index: stickerIdx,        // Use the calculated sticker index
				}
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
