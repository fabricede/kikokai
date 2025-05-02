package model

import (
	"reflect"
	"strings"
	"testing"
)

// Helper function to extract colors from a StickerColorName value
func extractColors(stickerColor string) []string {
	return strings.Split(stickerColor, "_")
}

func Test_CubeCornerColorsConsistency(t *testing.T) {
	// Map of cube corners to their corresponding StickerIndex values
	cubeCorners := map[string][3]StickerIndex{
		"Front-Left-Up":    {Front_2_0_0, Left_0_0_2, Up_2_0_2},
		"Front-Right-Up":   {Front_2_0_2, Right_0_2_0, Up_2_2_2},
		"Front-Left-Down":  {Front_2_2_0, Left_2_0_0, Down_0_0_0},
		"Front-Right-Down": {Front_2_2_2, Right_2_2_0, Down_0_2_0},
		"Back-Left-Up":     {Back_0_0_0, Left_0_0_0, Up_0_0_2},
		"Back-Right-Up":    {Back_0_0_2, Right_0_2_2, Up_0_2_2},
		"Back-Left-Down":   {Back_0_2_0, Left_2_0_2, Down_2_0_0},
		"Back-Right-Down":  {Back_0_2_2, Right_2_2_2, Down_2_2_0},
	}

	for cornerName, indices := range cubeCorners {
		t.Run(cornerName, func(t *testing.T) {
			// Extract colors for the three StickerIndex values
			colors1 := extractColors(StickerColorName[indices[0]])
			colors2 := extractColors(StickerColorName[indices[1]])
			colors3 := extractColors(StickerColorName[indices[2]])

			// Verify that the three sets of colors are the same (ignoring order)
			if !reflect.DeepEqual(toSet(colors1), toSet(colors2)) || !reflect.DeepEqual(toSet(colors1), toSet(colors3)) {
				t.Errorf("Colors for corner %s do not match: %v, %v, %v", cornerName, colors1, colors2, colors3)
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
