package model

import (
	"reflect"
	"testing"
)

func TestGetSticker(t *testing.T) {
	tests := []struct {
		name     string
		face     FaceIndex
		position CubeCoordinate
		want     Sticker
	}{
		// TODO: Add test cases.
		{
			name:     "Front NE sticker",
			face:     Front,
			position: CubeCoordinate{X: 1, Y: 1, Z: 1},
			want:     Sticker{Index: Front_2_2_2, Color: White},
		},
		{
			name:     "Front SE sticker",
			face:     Front,
			position: CubeCoordinate{X: 1, Y: 1, Z: -1},
			want:     Sticker{Index: Front_2_0_2, Color: White},
		},
		{
			name:     "Front SW sticker",
			face:     Front,
			position: CubeCoordinate{X: 1, Y: -1, Z: -1},
			want:     Sticker{Index: Front_2_0_0, Color: White},
		},
		{
			name:     "Front NW sticker",
			face:     Front,
			position: CubeCoordinate{X: 1, Y: -1, Z: 1},
			want:     Sticker{Index: Front_2_2_0, Color: White},
		},
		{
			name:     "Back NE sticker",
			face:     Back,
			position: CubeCoordinate{X: -1, Y: 1, Z: 1},
			want:     Sticker{Index: Back_0_2_2, Color: Yellow},
		},
		{
			name:     "Back SE sticker",
			face:     Back,
			position: CubeCoordinate{X: -1, Y: -1, Z: 1},
			want:     Sticker{Index: Back_0_2_0, Color: Yellow},
		},
		{
			name:     "Back SW sticker",
			face:     Back,
			position: CubeCoordinate{X: -1, Y: -1, Z: -1},
			want:     Sticker{Index: Back_0_0_0, Color: Yellow},
		},
		{
			name:     "Back NW sticker",
			face:     Back,
			position: CubeCoordinate{X: -1, Y: 1, Z: -1},
			want:     Sticker{Index: Back_0_0_2, Color: Yellow},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Define the colors for the cube
			cube := NewCube()
			if got, _ := cube.GetSticker(tt.face, tt.position); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSticker() = %v, want %v", got, tt.want)
			}
		})
	}
}
