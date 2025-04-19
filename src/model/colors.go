package model

// Color represents a color on the Rubik's cube
type Color int        // Color constants
type StickerIndex int // Sticker constants

const (
	// Common Colors
	White Color = iota
	Yellow
	Blue
	Green
	Red
	Orange

	// Sticker Colors
	// Front Face
	Front_NW StickerIndex = iota
	Front_N
	Front_NE
	Front_W
	Front_Center
	Front_E
	Front_SW
	Front_S
	Front_SE
	// Back Face
	Back_NW
	Back_N
	Back_NE
	Back_W
	Back_Center
	Back_E
	Back_SW
	Back_S
	Back_SE
	// Up Face
	Up_NW
	Up_N
	Up_NE
	Up_W
	Up_Center
	Up_E
	Up_SW
	Up_S
	Up_SE
	// Down Face
	Down_NW
	Down_N
	Down_NE
	Down_W
	Down_Center
	Down_E
	Down_SW
	Down_S
	Down_SE
	// Left Face
	Left_NW
	Left_N
	Left_NE
	Left_W
	Left_Center
	Left_E
	Left_SW
	Left_S
	Left_SE
	// Right Face
	Right_NW
	Right_N
	Right_NE
	Right_W
	Right_Center
	Right_E
	Right_SW
	Right_S
	Right_SE
)

// ColorNames maps Color constants to their string representations
var FaceColorName = map[FaceIndex]string{
	Front: "white",
	Back:  "yellow",
	Up:    "blue",
	Down:  "green",
	Left:  "red",
	Right: "orange",
}

var StickerColorName = map[StickerIndex]string{
	Front_NW:     "white_red_blue",
	Front_N:      "white_blue",
	Front_NE:     "white_blue_orange",
	Front_W:      "white_red",
	Front_Center: "white",
	Front_E:      "white_orange",
	Front_SW:     "white_red_green",
	Front_S:      "white_green",
	Front_SE:     "white_green_orange",
	Back_NW:      "yellow_red_blue",
	Back_N:       "yellow_blue",
	Back_NE:      "yellow_blue_orange",
	Back_W:       "yellow_red",
	Back_Center:  "yellow",
	Back_E:       "yellow_orange",
	Back_SW:      "yellow_red_green",
	Back_S:       "yellow_green",
	Back_SE:      "yellow_green_orange",
	Up_NW:        "blue_red_blue",
	Up_N:         "blue_blue",
	Up_NE:        "blue_blue_orange",
	Up_W:         "blue_red",
	Up_Center:    "blue",
	Up_E:         "blue_orange",
	Up_SW:        "blue_red_green",
	Up_S:         "blue_green",
	Up_SE:        "blue_green_orange",
	Down_NE:      "green_red_blue",
	Down_N:       "green_blue",
	Down_NW:      "green_blue_orange",
	Down_W:       "green_red",
	Down_Center:  "green",
	Down_E:       "green_orange",
	Down_SW:      "green_red_green",
	Down_S:       "green_green",
	Down_SE:      "green_green_orange",
	Left_NE:      "red_red_blue",
	Left_N:       "red_blue",
	Left_NW:      "red_blue_orange",
	Left_W:       "red_red",
	Left_Center:  "red",
	Left_E:       "red_orange",
	Left_SW:      "red_red_green",
	Left_S:       "red_green",
	Left_SE:      "red_green_orange",
	Right_NE:     "orange_red_blue",
	Right_N:      "orange_blue",
	Right_NW:     "orange_blue_orange",
	Right_W:      "orange_red",
	Right_Center: "orange",
	Right_E:      "orange_orange",
	Right_SW:     "orange_red_green",
	Right_S:      "orange_green",
	Right_SE:     "orange_green_orange",
}

// ColorNames maps Color constants to their string representations
