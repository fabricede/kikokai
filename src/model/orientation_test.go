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

func TestGetAdjacentFace(t *testing.T) {
	type args struct {
		face FaceIndex
		x    int
		y    int
		z    int
	}
	tests := []struct {
		name string
		args args
		want FaceIndex
	}{
		// TODO: Add test cases.
		{
			name: "Front face y",
			args: args{
				face: Front,
				x:    0,
				y:    1,
				z:    0,
			},
			want: Up,
		},
		{
			name: "Front face -y",
			args: args{
				face: Front,
				x:    0,
				y:    -1,
				z:    0,
			},
			want: Down,
		},
		{
			name: "Back face",
			args: args{
				face: Back,
				x:    0,
				y:    1,
				z:    0,
			},
			want: Up,
		},
		{
			name: "Left face",
			args: args{
				face: Left,
				x:    0,
				y:    1,
				z:    0,
			},
			want: Up,
		},
		{
			name: "Right face",
			args: args{
				face: Right,
				x:    0,
				y:    1,
				z:    0,
			},
			want: Up,
		},
		{
			name: "Up face x",
			args: args{
				face: Up,
				x:    1,
				y:    0,
				z:    0,
			},
			want: Front,
		},
		{
			name: "Up face -x",
			args: args{
				face: Up,
				x:    -1,
				y:    0,
				z:    0,
			},
			want: Back,
		},
		{
			name: "Down face",
			args: args{
				face: Down,
				x:    1,
				y:    0,
				z:    0,
			},
			want: Front,
		},
		{
			name: "Up face z",
			args: args{
				face: Up,
				x:    0,
				y:    0,
				z:    1,
			},
			want: Right,
		},
		{
			name: "Up face -z",
			args: args{
				face: Up,
				x:    0,
				y:    0,
				z:    -1,
			},
			want: Left,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAdjacentFace(tt.args.face, tt.args.x, tt.args.y, tt.args.z); got != tt.want {
				t.Errorf("GetAdjacentFace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAdjacentXEdge(t *testing.T) {
	type args struct {
		cube *Cube
		face FaceIndex
		x    int
	}
	tests := []struct {
		name             string
		args             args
		wantAdjacentFace FaceIndex
		wantEdge         [3]Sticker
		wantErr          bool
	}{
		// TODO: Add test cases.
		{
			name: "up face +x",
			args: args{
				cube: initialCube,
				face: Up,
				x:    1,
			},
			wantAdjacentFace: Front,
			wantEdge:         [3]Sticker{{Index: Front_2_0_2, Color: White}, {Index: Front_2_1_2, Color: White}, {Index: Front_2_2_2, Color: White}},
			wantErr:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAdjacentFace, gotEdge, err := GetAdjacentXEdge(tt.args.cube, tt.args.face, tt.args.x)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAdjacentXEdge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAdjacentFace != tt.wantAdjacentFace {
				t.Errorf("GetAdjacentXEdge() gotAdjacentFace = %v, want %v", gotAdjacentFace, tt.wantAdjacentFace)
			}
			if !reflect.DeepEqual(gotEdge, tt.wantEdge) {
				t.Errorf("GetAdjacentXEdge() gotEdge = %v, want %v", gotEdge, tt.wantEdge)
			}
		})
	}
}
