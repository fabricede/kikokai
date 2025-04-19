package model

import (
	"reflect"
	"testing"
)

func TestGetNorthFace(t *testing.T) {
	type args struct {
		face FaceIndex
	}
	tests := []struct {
		name  string
		args  args
		want  FaceIndex
		want1 Orientation
	}{
		// TODO: Add test cases.
		{
			name: "Test 1",
			args: args{
				face: Front,
			},
			want:  Up,
			want1: West,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetNorthFace(tt.args.face)
			if got != tt.want {
				t.Errorf("GetNorthFace() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetNorthFace() got1 = %v, want %v", got1, tt.want1)
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
			wantEdge: [3]Sticker{{Index: Up_NE}, {Index: Up_E}, {Index: Up_SE}},
		},
		{
			name: "back North edge",
			args: args{
				cube: cube,
				face: Back,
			},
			wantEdge: [3]Sticker{{Index: Down_NE}, {Index: Down_E}, {Index: Down_SE}},
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
