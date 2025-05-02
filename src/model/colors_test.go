package model

import (
	"reflect"
	"strings"
	"testing"

	"log"
)

// Helper function to extract colors from a StickerColorName value
func extractColors(stickerColor string) []string {
	return strings.Split(stickerColor, "_")
}

func Test_CubeCornerColorsConsistency(t *testing.T) {
	// Map of cube corners to their corresponding faces
	cubeCorners := map[CubeCoordinate][3]FaceIndex{
		{-1, -1, -1}: {Back, Left, Down},
		{-1, -1, 1}:  {Back, Left, Up},
		//{-1, 1, -1}: {Back, Right, Down},
		//{-1, 1, 1}: {Back, Right, Up},
		//{1, -1, -1}: {Front, Left, Down},
		//{1, -1, 1}: {Front, Left, Up},
		//{1, 1, -1}: {Front, Right, Down},
		{1, 1, 1}: {Front, Right, Up},
	}

	cube := NewCube()

	for cornerCoord, face := range cubeCorners {
		// Generate a name for the corner based on its coordinates
		cornerName := strings.Join([]string{FaceColorName[face[0]], FaceColorName[face[1]], FaceColorName[face[2]]}, "-")
		t.Run(cornerName, func(t *testing.T) {
			face0row, face0col := cube.State[face[0]].GetFaceCoordinate(cornerCoord)
			log.Printf("face0row: %d, face0col: %d", face0row, face0col)
			face1row, face1col := cube.State[face[1]].GetFaceCoordinate(cornerCoord)
			log.Printf("face1row: %d, face1col: %d", face1row, face1col)
			face2row, face2col := cube.State[face[2]].GetFaceCoordinate(cornerCoord)
			log.Printf("face2row: %d, face2col: %d", face2row, face2col)
			indices := [3]StickerIndex{
				cube.State[face[0]].Stickers[face0row][face0col].GetIndex(),
				cube.State[face[1]].Stickers[face1row][face1col].GetIndex(),
				cube.State[face[2]].Stickers[face2row][face2col].GetIndex(),
			}
			log.Printf("Indices for corner %v : (%v)%v, (%v)%v, (%v)%v", cornerCoord, StickerColorName[indices[0]], indices[0], StickerColorName[indices[1]], indices[1], StickerColorName[indices[2]], indices[2])
			// Extract colors for the three StickerIndex values
			colors1 := extractColors(StickerColorName[indices[0]])
			colors2 := extractColors(StickerColorName[indices[1]])
			colors3 := extractColors(StickerColorName[indices[2]])

			log.Printf("Colors for corner %v : (%v)%v, (%v)%v, (%v)%v", cornerCoord, FaceColorName[face[0]], colors1, FaceColorName[face[1]], colors2, FaceColorName[face[2]], colors3)

			// Verify that the three sets of colors are the same (ignoring order)
			if !reflect.DeepEqual(toSet(colors1), toSet(colors2)) || !reflect.DeepEqual(toSet(colors1), toSet(colors3)) {
				t.Errorf("Colors for corner %v do not match: (%v)%v, (%v)%v, (%v)%v", cornerCoord, FaceColorName[face[0]], colors1, FaceColorName[face[1]], colors2, FaceColorName[face[2]], colors3)
			}
		})
	}
}

func Test_CubeEdgeColorsConsistency(t *testing.T) {
	// Map of cube edges to their corresponding StickerIndex values
	cubeEdges := map[string][2]StickerIndex{
		"Front-Up":    {Front_2_0_1, Up_2_1_2},
		"Front-Right": {Front_2_1_2, Right_1_2_1},
		"Front-Down":  {Front_2_2_1, Down_0_1_0},
		"Front-Left":  {Front_2_1_0, Left_1_0_1},
		"Back-Up":     {Back_0_0_1, Up_0_1_2},
		"Back-Right":  {Back_0_1_2, Right_0_2_1},
		"Back-Down":   {Back_0_2_1, Down_2_1_0},
		"Back-Left":   {Back_0_1_0, Left_0_0_1},
		"Left-Up":     {Left_0_0_1, Up_0_1_2},
		"Left-Down":   {Left_2_0_1, Down_2_1_0},
		"Right-Up":    {Right_0_2_1, Up_0_2_2},
		"Right-Down":  {Right_2_2_1, Down_0_2_0},
	}

	for edgeName, indices := range cubeEdges {
		t.Run(edgeName, func(t *testing.T) {
			// Extract colors for the three StickerIndex values
			colors1 := extractColors(StickerColorName[indices[0]])
			colors2 := extractColors(StickerColorName[indices[1]])

			// Verify that the three sets of colors are the same (ignoring order)
			if !reflect.DeepEqual(toSet(colors1), toSet(colors2)) {
				t.Errorf("Colors for edges %s do not match: %v, %v", edgeName, colors1, colors2)
			}
		})
	}
}

// Helper function to convert a slice of strings to a set (map with empty struct values)
func toSet(colors []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, color := range colors {
		set[color] = struct{}{}
	}
	return set
}
