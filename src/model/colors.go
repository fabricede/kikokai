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
)
const (
	// Sticker Colors
	// Front Face (x=1) => 2
	Front_2_0_0 StickerIndex = iota
	Front_2_0_1
	Front_2_0_2
	Front_2_1_0
	Front_2_1_1
	Front_2_1_2
	Front_2_2_0
	Front_2_2_1
	Front_2_2_2
	// Back Face (x=-1) => 0
	Back_0_0_0
	Back_0_0_1
	Back_0_0_2
	Back_0_1_0
	Back_0_1_1
	Back_0_1_2
	Back_0_2_0
	Back_0_2_1
	Back_0_2_2
	// Up Face (z=1) => 2
	Up_0_0_2
	Up_1_0_2
	Up_2_0_2
	Up_0_1_2
	Up_1_1_2
	Up_2_1_2
	Up_0_2_2
	Up_1_2_2
	Up_2_2_2
	// Down Face (z=-1) => 0
	Down_0_0_0
	Down_1_0_0
	Down_2_0_0
	Down_0_1_0
	Down_1_1_0
	Down_2_1_0
	Down_0_2_0
	Down_1_2_0
	Down_2_2_0
	// Left Face (y=-1) => 0
	Left_0_0_0
	Left_0_0_1
	Left_0_0_2
	Left_1_0_0
	Left_1_0_1
	Left_1_0_2
	Left_2_0_0
	Left_2_0_1
	Left_2_0_2
	// Right Face (y=1) => 2
	Right_0_2_0
	Right_0_2_1
	Right_0_2_2
	Right_1_2_0
	Right_1_2_1
	Right_1_2_2
	Right_2_2_0
	Right_2_2_1
	Right_2_2_2
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
	// Front Face
	// row = 1 + position.Z // Down /South (z=-1) is row 0, Up   /North (z=1) is row 2
	// col = 1 + position.Y // Left/West  (y=-1)  is col 0, Right/East  (y=1) is col 2
	Front_2_0_0: "white_red_green",    // row 0, col 0
	Front_2_0_1: "white_green",        // row 0, col 1
	Front_2_0_2: "white_green_orange", // row 0, col 2
	Front_2_1_0: "white_red",          // row 1, col 0
	Front_2_1_1: "white",              // row 1, col 1
	Front_2_1_2: "white_orange",       // row 1, col 2
	Front_2_2_0: "white_red_blue",     // row 2, col 0
	Front_2_2_1: "white_blue",         // row 2, col 1
	Front_2_2_2: "white_blue_orange",  // row 2, col 2
	// Back Face
	Back_0_2_0: "yellow_red_blue",
	Back_0_2_1: "yellow_blue",
	Back_0_2_2: "yellow_blue_orange",
	Back_0_1_0: "yellow_red",
	Back_0_1_1: "yellow",
	Back_0_1_2: "yellow_orange",
	Back_0_0_0: "yellow_red_green",
	Back_0_0_1: "yellow_green",
	Back_0_0_2: "yellow_green_orange",
	// Up Face
	Up_0_2_2: "blue_red_blue",
	Up_1_2_2: "blue_blue",
	Up_2_2_2: "blue_blue_orange",
	Up_0_1_2: "blue_red",
	Up_1_1_2: "blue",
	Up_2_1_2: "blue_orange",
	Up_0_0_2: "blue_red_green",
	Up_1_0_2: "blue_green",
	Up_2_0_2: "blue_green_orange",
	// Down Face
	Down_2_2_0: "green_red_blue",
	Down_1_2_0: "green_blue",
	Down_0_2_0: "green_blue_orange",
	Down_0_1_0: "green_red",
	Down_1_1_0: "green",
	Down_2_1_0: "green_orange",
	Down_0_0_0: "green_red_green",
	Down_1_0_0: "green_green",
	Down_2_0_0: "green_green_orange",
	// Left Face
	Left_2_0_2: "red_red_blue",
	Left_2_0_1: "red_blue",
	Left_2_0_0: "red_blue_orange",
	Left_1_0_0: "red_red",
	Left_1_0_1: "red",
	Left_1_0_2: "red_orange",
	Left_0_0_0: "red_red_green",
	Left_0_0_1: "red_green",
	Left_0_0_2: "red_green_orange",
	// Right Face
	Right_2_2_2: "orange_red_blue",
	Right_2_2_1: "orange_blue",
	Right_2_2_0: "orange_blue_orange",
	Right_1_2_0: "orange_red",
	Right_1_2_1: "orange",
	Right_1_2_2: "orange_orange",
	Right_0_2_0: "orange_red_green",
	Right_0_2_1: "orange_green",
	Right_0_2_2: "orange_green_orange",
}

// ColorNames maps Color constants to their string representations
