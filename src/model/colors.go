package model

// Color represents a color on the Rubik's cube
type Color int // Color constants
const (
	White Color = iota
	Yellow
	Blue
	Green
	Red
	Orange
)

// ColorNames maps Color constants to their string representations
var ColorNames = map[Color]string{
	White:  "white",
	Yellow: "yellow",
	Blue:   "blue",
	Green:  "green",
	Red:    "red",
	Orange: "orange",
}
