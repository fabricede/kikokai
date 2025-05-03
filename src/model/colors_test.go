package model

import (
	"strings"
)

// Helper function to extract colors from a StickerColorName value
func extractColors(stickerColor string) []string {
	return strings.Split(stickerColor, "_")
}
