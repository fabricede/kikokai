package model

// Color represents a color on the Rubik's cube
type Color int        // Color constants
type StickerIndex int // Sticker constants

const (
	// Common Colors
	White Color = iota
	Orange
	Yellow
	Red
	Blue
	Green
)

// ColorNames maps Color constants to their string representations
var FaceColorName = map[FaceIndex]string{
	Front: "white",
	Right: "orange",
	Back:  "yellow",
	Left:  "red",
	Up:    "blue",
	Down:  "green",
}

const (
	// Sticker Colors
	// Front Face (x=1) => 2 Y Z
	Front_2_0_0 StickerIndex = iota
	Front_2_0_1
	Front_2_0_2
	Front_2_1_0
	Front_2_1_1
	Front_2_1_2
	Front_2_2_0
	Front_2_2_1
	Front_2_2_2
	// Back Face (x=-1) => 0 Z Y
	Back_0_0_0
	Back_0_0_1
	Back_0_0_2
	Back_0_1_0
	Back_0_1_1
	Back_0_1_2
	Back_0_2_0
	Back_0_2_1
	Back_0_2_2
	// Up Face (z=1) => Y X 2
	Up_0_0_2
	Up_1_0_2
	Up_2_0_2
	Up_0_1_2
	Up_1_1_2
	Up_2_1_2
	Up_0_2_2
	Up_1_2_2
	Up_2_2_2
	// Down Face (z=-1) => Y X 0
	Down_0_0_0
	Down_1_0_0
	Down_2_0_0
	Down_0_1_0
	Down_1_1_0
	Down_2_1_0
	Down_0_2_0
	Down_1_2_0
	Down_2_2_0
	// Left Face (y=-1) => Z 0 X
	Left_0_0_0
	Left_0_0_1
	Left_0_0_2
	Left_1_0_0
	Left_1_0_1
	Left_1_0_2
	Left_2_0_0
	Left_2_0_1
	Left_2_0_2
	// Right Face (y=1) => Z 2 X
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

var StickerColorName = map[StickerIndex]string{
	// Front Face (White) YZ
	Front_2_0_0: "white_green_red",    // bottom-left (row 0, col 0)
	Front_2_0_1: "white_green",        // bottom-center (row 0, col 1)
	Front_2_0_2: "white_green_orange", // bottom-right (row 0, col 2)
	Front_2_1_0: "white_red",          // middle-left (row 1, col 0)
	Front_2_1_1: "white",              // center (row 1, col 1)
	Front_2_1_2: "white_orange",       // middle-right (row 1, col 2)
	Front_2_2_0: "white_blue_red",     // top-left (row 2, col 0)
	Front_2_2_1: "white_blue",         // top-center (row 2, col 1)
	Front_2_2_2: "white_blue_orange",  // top-right (row 2, col 2)

	// Back Face (Yellow) ZY
	Back_0_0_0: "yellow_green_red",    // bottom-left (row 0, col 0)
	Back_0_0_1: "yellow_red",          // bottom-center (row 0, col 1)
	Back_0_0_2: "yellow_red_blue",     // bottom-right (row 0, col 2)
	Back_0_1_0: "yellow_green",        // middle-left (row 1, col 0)
	Back_0_1_1: "yellow",              // center (row 1, col 1)
	Back_0_1_2: "yellow_blue",         // middle-right (row 1, col 2)
	Back_0_2_0: "yellow_green_orange", // top-left (row 2, col 0)
	Back_0_2_1: "yellow_orange",       // top-center (row 2, col 1)
	Back_0_2_2: "yellow_blue_orange",  // top-right (row 2, col 2)

	// Up Face (Blue) ZX
	Up_0_0_2: "blue_yellow_red",    // bottom-left (row 0, col 0)
	Up_0_1_2: "blue_yellow",        // bottom-center (row 0, col 1)
	Up_0_2_2: "blue_yellow_orange", // bottom-right (row 0, col 2)
	Up_1_0_2: "blue_red",           // middle-left (row 1, col 0)
	Up_1_1_2: "blue",               // center (row 1, col 1)
	Up_1_2_2: "blue_orange",        // middle-right (row 1, col 2)
	Up_2_0_2: "blue_white_red",     // top-left (row 2, col 0)
	Up_2_1_2: "blue_white",         // top-center (row 2, col 1)
	Up_2_2_2: "blue_white_orange",  // top-right (row 2, col 2)

	// Down Face (Green) ZX
	Down_0_0_0: "green_yellow_red",    // bottom-left (row 0, col 0)
	Down_0_1_0: "green_red",           // bottom-center (row 0, col 1)
	Down_0_2_0: "green_white_red",     // bottom-left (row 0, col 2)
	Down_1_0_0: "green_yellow",        // middle-left (row 1, col 0)
	Down_1_1_0: "green",               // center (row 1, col 1)
	Down_1_2_0: "green_white",         // middle-right (row 1, col 2)
	Down_2_0_0: "green_yellow_orange", // top-left (row 2, col 0)
	Down_2_1_0: "green_orange",        // top-center (row 2, col 1)
	Down_2_2_0: "green_white_orange",  // top-right (row 2, col 2)

	// Left Face (Red) YX
	Left_0_0_0: "red_yellow_green", // bottom-left (row 0, col 0)
	Left_0_0_1: "red_green",        // bottom-center (row 0, col 1)
	Left_0_0_2: "red_green_white",  // bottom-right (row 0, col 2)
	Left_1_0_0: "red_yellow",       // middle-left (row 1, col 0)
	Left_1_0_1: "red",              // center (row 1, col 1)
	Left_1_0_2: "red_white",        // middle-right (row 1, col 2)
	Left_2_0_0: "red_blue_yellow",  // top-left (row 2, col 0)
	Left_2_0_1: "red_blue",         // top-center (row 2, col 1)
	Left_2_0_2: "red_white_blue",   // top-right (row 2, col 2)

	// Right Face (Orange) YX
	Right_0_2_0: "orange_yellow_green", // bottom-left (row 0, col 0)
	Right_0_2_1: "orange_green",        // bottom-center (row 0, col 1)
	Right_0_2_2: "orange_white_green",  // bottom-right (row 0, col 2)
	Right_1_2_0: "orange_green",        // middle-left (row 1, col 0)
	Right_1_2_1: "orange",              // center (row 1, col 1)
	Right_1_2_2: "orange_blue",         // middle-right (row 1, col 2)
	Right_2_2_0: "orange_white_green",  // top-left (row 2, col 0)
	Right_2_2_1: "orange_white",        // top-center (row 2, col 1)
	Right_2_2_2: "orange_white_blue",   // top-right (row 2, col 2)
}
