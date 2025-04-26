package model

import (
	"reflect"
	"testing"
)

func TestGetNorthFace(t *testing.T) {
	tests := []struct {
		name  string
		face  FaceIndex
		want  FaceIndex
		want1 Orientation
	}{
		// TODO: Add test cases.
		{name: "Front", face: Front, want: Up, want1: West},
		{name: "Up", face: Up, want: Left, want1: East},
		{name: "Back", face: Back, want: Down, want1: East},
		{name: "Down", face: Down, want: Right, want1: West},
		{name: "Left", face: Left, want: Back, want1: West},
		{name: "Right", face: Right, want: Front, want1: East},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetNorthFace(tt.face)
			if got != tt.want {
				t.Errorf("GetNorthFace() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetNorthFace() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGetEastFace(t *testing.T) {
	tests := []struct {
		name  string
		face  FaceIndex
		want  FaceIndex
		want1 Orientation
	}{
		{name: "Front", face: Front, want: Right, want1: North},
		{name: "Up", face: Up, want: Back, want1: South},
		{name: "Back", face: Back, want: Right, want1: South},
		{name: "Down", face: Down, want: Back, want1: North},
		{name: "Left", face: Left, want: Up, want1: North},
		{name: "Right", face: Right, want: Up, want1: South},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetEastFace(tt.face)
			if got != tt.want {
				t.Errorf("GetEastFace() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetEastFace() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGetWestFace(t *testing.T) {
	tests := []struct {
		name  string
		face  FaceIndex
		want  FaceIndex
		want1 Orientation
	}{
		{name: "Front", face: Front, want: Left, want1: South},
		{name: "Up", face: Up, want: Front, want1: North},
		{name: "Back", face: Back, want: Left, want1: North},
		{name: "Down", face: Down, want: Front, want1: South},
		{name: "Left", face: Left, want: Down, want1: South},
		{name: "Right", face: Right, want: Down, want1: North},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetWestFace(tt.face)
			if got != tt.want {
				t.Errorf("GetWestFace() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetWestFace() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGetSouthFace(t *testing.T) {
	tests := []struct {
		name  string
		face  FaceIndex
		want  FaceIndex
		want1 Orientation
	}{
		{name: "Front", face: Front, want: Down, want1: West},
		{name: "Up", face: Up, want: Right, want1: East},
		{name: "Back", face: Back, want: Up, want1: East},
		{name: "Down", face: Down, want: Left, want1: West},
		{name: "Left", face: Left, want: Front, want1: West},
		{name: "Right", face: Right, want: Back, want1: East},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetSouthFace(tt.face)
			if got != tt.want {
				t.Errorf("GetSouthFace() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetSouthFace() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGetNorthEdge(t *testing.T) {
	type args struct {
		cube *Cube
		face FaceIndex
	}
	// Define the colors for the cube
	cube := NewCube()
	tests := []struct {
		name     string
		args     args
		wantEdge [3]Sticker
	}{
		// TODO: Add test cases.
		{
			name: "front North edge",
			args: args{
				cube: cube,
				face: Front,
			},
			wantEdge: [3]Sticker{{Index: Up_NW, Color: Blue}, {Index: Up_W, Color: Blue}, {Index: Up_SW, Color: Blue}},
		},
		{
			name: "back North edge",
			args: args{
				cube: cube,
				face: Back,
			},
			wantEdge: [3]Sticker{{Index: Down_NE, Color: Green}, {Index: Down_E, Color: Green}, {Index: Down_SE, Color: Green}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotEdge := GetNorthEdge(tt.args.cube, tt.args.face); !reflect.DeepEqual(gotEdge, tt.wantEdge) {
				t.Errorf("GetNorthEdge() = %v, want %v", gotEdge, tt.wantEdge)
			}
		})
	}
}

func TestGetSticker(t *testing.T) {
	tests := []struct {
		name     string
		position CubeCoordinate
		want     Sticker
	}{
		// TODO: Add test cases.
		{
			name:     "Front NE sticker",
			position: CubeCoordinate{X: 1, Y: 1, Z: 1},
			want:     Sticker{Index: Front_2_2_2, Color: White},
		},
		{
			name:     "Front SE sticker",
			position: CubeCoordinate{X: 1, Y: 1, Z: -1},
			want:     Sticker{Index: Front_2_0_2, Color: White},
		},
		{
			name:     "Front SW sticker",
			position: CubeCoordinate{X: 1, Y: -1, Z: -1},
			want:     Sticker{Index: Front_2_0_0, Color: White},
		},
		{
			name:     "Front NW sticker",
			position: CubeCoordinate{X: 1, Y: -1, Z: 1},
			want:     Sticker{Index: Front_2_2_0, Color: White},
		},
		{
			name:     "Back NE sticker",
			position: CubeCoordinate{X: -1, Y: 1, Z: 1},
			want:     Sticker{Index: Back_NE, Color: Yellow},
		},
		{
			name:     "Back NW sticker",
			position: CubeCoordinate{X: -1, Y: 1, Z: -1},
			want:     Sticker{Index: Back_NW, Color: Yellow},
		},
		{
			name:     "Back SW sticker",
			position: CubeCoordinate{X: -1, Y: -1, Z: -1},
			want:     Sticker{Index: Back_SW, Color: Yellow},
		},
		{
			name:     "Back SE sticker",
			position: CubeCoordinate{X: -1, Y: -1, Z: 1},
			want:     Sticker{Index: Back_SE, Color: Yellow},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Define the colors for the cube
			cube := NewCube()
			if got := GetSticker(cube, tt.position); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSticker() = %v, want %v", got, tt.want)
			}
		})
	}
}
