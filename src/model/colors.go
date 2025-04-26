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
	// Front Face (x=2)
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

	// Back Face (x=0)
	// row = 2 - position.Z // Down /South (z=-1) is row 2, Up   /North (z=1) is row 0
	// col = 2 - position.Y // Left/West  (y=-1) is col 2, Right/East  (y=1) is col 0
	Back_0_2_0: "yellow_red_blue",     // row 0, col 0
	Back_0_2_1: "yellow_blue",         // row 0, col 1
	Back_0_2_2: "yellow_blue_orange",  // row 0, col 2
	Back_0_1_0: "yellow_red",          // row 1, col 0
	Back_0_1_1: "yellow",              // row 1, col 1
	Back_0_1_2: "yellow_orange",       // row 1, col 2
	Back_0_0_0: "yellow_red_green",    // row 2, col 0
	Back_0_0_1: "yellow_green",        // row 2, col 1
	Back_0_0_2: "yellow_green_orange", // row 2, col 2

	// Up Face (z=2)
	// row = 2 - position.X // Back  (x=-1) is row 2, Front (x=1) is row 0
	// col = 1 + position.Y // Left  (y=-1) is col 0, Right (y=1) is col 2
	Up_0_2_2: "blue_red_blue",     // row 0, col 2
	Up_1_2_2: "blue_blue",         // row 1, col 2
	Up_2_2_2: "blue_blue_orange",  // row 2, col 2
	Up_0_1_2: "blue_red",          // row 0, col 1
	Up_1_1_2: "blue",              // row 1, col 1
	Up_2_1_2: "blue_orange",       // row 2, col 1
	Up_0_0_2: "blue_red_green",    // row 0, col 0
	Up_1_0_2: "blue_green",        // row 1, col 0
	Up_2_0_2: "blue_green_orange", // row 2, col 0

	// Down Face (z=0)
	// row = position.X + 1 // Back  (x=-1) is row 0, Front (x=1) is row 2
	// col = position.Y + 1 // Left  (y=-1) is col 0, Right (y=1) is col 2
	Down_2_2_0: "green_red_blue",     // row 2, col 2
	Down_1_2_0: "green_blue",         // row 1, col 2
	Down_0_2_0: "green_blue_orange",  // row 0, col 2
	Down_0_1_0: "green_red",          // row 0, col 1
	Down_1_1_0: "green",              // row 1, col 1
	Down_2_1_0: "green_orange",       // row 2, col 1
	Down_0_0_0: "green_red_green",    // row 0, col 0
	Down_1_0_0: "green_green",        // row 1, col 0
	Down_2_0_0: "green_green_orange", // row 2, col 0

	// Left Face (y=0)
	// row = position.X + 1 // Back  (x=-1) is row 0, Front (x=1) is row 2
	// col = position.Z + 1 // Down  (z=-1) is col 0, Up    (z=1) is col 2
	Left_2_0_2: "red_red_blue",     // row 2, col 2
	Left_2_0_1: "red_blue",         // row 2, col 1
	Left_2_0_0: "red_blue_orange",  // row 2, col 0
	Left_1_0_0: "red_red",          // row 1, col 0
	Left_1_0_1: "red",              // row 1, col 1
	Left_1_0_2: "red_orange",       // row 1, col 2
	Left_0_0_0: "red_red_green",    // row 0, col 0
	Left_0_0_1: "red_green",        // row 0, col 1
	Left_0_0_2: "red_green_orange", // row 0, col 2

	// Right Face (y=2)
	// row = position.X + 1 // Back  (x=-1) is row 0, Front (x=1) is row 2
	// col = 2 - position.Z // Down  (z=-1) is col 2, Up    (z=1) is col 0
	Right_2_2_2: "orange_red_blue",     // row 2, col 2
	Right_2_2_1: "orange_blue",         // row 2, col 1
	Right_2_2_0: "orange_blue_orange",  // row 2, col 0
	Right_1_2_0: "orange_red",          // row 1, col 0
	Right_1_2_1: "orange",              // row 1, col 1
	Right_1_2_2: "orange_orange",       // row 1, col 2
	Right_0_2_0: "orange_red_green",    // row 0, col 0
	Right_0_2_1: "orange_green",        // row 0, col 1
	Right_0_2_2: "orange_green_orange", // row 0, col 2
}
