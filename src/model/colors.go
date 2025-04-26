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
	// Front Face (White)
	Front_2_0_0: "white_green_red",    // top-left (row 0, col 0)
	Front_2_0_1: "white_green",        // top-center (row 0, col 1)
	Front_2_0_2: "white_green_orange", // top-right (row 0, col 2)
	Front_2_1_0: "white_red",          // middle-left (row 1, col 0)
	Front_2_1_1: "white",              // center (row 1, col 1)
	Front_2_1_2: "white_orange",       // middle-right (row 1, col 2)
	Front_2_2_0: "white_blue_red",     // bottom-left (row 2, col 0)
	Front_2_2_1: "white_blue",         // bottom-center (row 2, col 1)
	Front_2_2_2: "white_blue_orange",  // bottom-right (row 2, col 2)

	// Back Face (Yellow)
	Back_0_0_0: "yellow_green_red",    // top-left (row 0, col 0)
	Back_0_0_1: "yellow_green",        // top-center (row 0, col 1)
	Back_0_0_2: "yellow_green_orange", // top-right (row 0, col 2)
	Back_0_1_0: "yellow_red",          // middle-left (row 1, col 0)
	Back_0_1_1: "yellow",              // center (row 1, col 1)
	Back_0_1_2: "yellow_orange",       // middle-right (row 1, col 2)
	Back_0_2_0: "yellow_blue_red",     // bottom-left (row 2, col 0)
	Back_0_2_1: "yellow_blue",         // bottom-center (row 2, col 1)
	Back_0_2_2: "yellow_blue_orange",  // bottom-right (row 2, col 2)

	// Up Face (Blue)
	Up_0_0_2: "blue_yellow_red",    // top-left (row 0, col 0)
	Up_0_1_2: "blue_yellow",        // top-center (row 0, col 1)
	Up_0_2_2: "blue_yellow_orange", // top-right (row 0, col 2)
	Up_1_0_2: "blue_red",           // middle-left (row 1, col 0)
	Up_1_1_2: "blue",               // center (row 1, col 1)
	Up_1_2_2: "blue_orange",        // middle-right (row 1, col 2)
	Up_2_0_2: "blue_white_red",     // bottom-left (row 2, col 0)
	Up_2_1_2: "blue_white",         // bottom-center (row 2, col 1)
	Up_2_2_2: "blue_white_orange",  // bottom-right (row 2, col 2)

	// Down Face (Green)
	Down_0_0_0: "green_yellow_red",    // top-left (row 0, col 0)
	Down_0_1_0: "green_yellow",        // top-center (row 0, col 1)
	Down_0_2_0: "green_yellow_orange", // top-right (row 0, col 2)
	Down_1_0_0: "green_red",           // middle-left (row 1, col 0)
	Down_1_1_0: "green",               // center (row 1, col 1)
	Down_1_2_0: "green_orange",        // middle-right (row 1, col 2)
	Down_2_0_0: "green_white_red",     // bottom-left (row 2, col 0)
	Down_2_1_0: "green_white",         // bottom-center (row 2, col 1)
	Down_2_2_0: "green_white_orange",  // bottom-right (row 2, col 2)

	// Left Face (Red)
	Left_0_0_0: "red_yellow_green", // top-left (row 0, col 0)
	Left_0_0_1: "red_yellow",       // top-center (row 0, col 1)
	Left_0_0_2: "red_yellow_blue",  // top-right (row 0, col 2)
	Left_1_0_0: "red_green",        // middle-left (row 1, col 0)
	Left_1_0_1: "red",              // center (row 1, col 1)
	Left_1_0_2: "red_blue",         // middle-right (row 1, col 2)
	Left_2_0_0: "red_white_green",  // bottom-left (row 2, col 0)
	Left_2_0_1: "red_white",        // bottom-center (row 2, col 1)
	Left_2_0_2: "red_white_blue",   // bottom-right (row 2, col 2)

	// Right Face (Orange)
	Right_0_2_0: "orange_yellow_green", // top-left (row 0, col 0)
	Right_0_2_1: "orange_yellow",       // top-center (row 0, col 1)
	Right_0_2_2: "orange_yellow_blue",  // top-right (row 0, col 2)
	Right_1_2_0: "orange_green",        // middle-left (row 1, col 0)
	Right_1_2_1: "orange",              // center (row 1, col 1)
	Right_1_2_2: "orange_blue",         // middle-right (row 1, col 2)
	Right_2_2_0: "orange_white_green",  // bottom-left (row 2, col 0)
	Right_2_2_1: "orange_white",        // bottom-center (row 2, col 1)
	Right_2_2_2: "orange_white_blue",   // bottom-right (row 2, col 2)
}
