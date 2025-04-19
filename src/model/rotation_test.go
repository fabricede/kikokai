package model

import (
	"reflect"
	"testing"
)

func TestMatrix5x5_RotateClockwise(t *testing.T) {
	tests := []struct {
		name string
		m    Matrix5x5
		want Matrix5x5
	}{
		// TODO: Add test cases.
		{
			name: "Test 1",
			m: *SetMatrix([5][5]StickerIndex{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25}}),
			want: *SetMatrix([5][5]StickerIndex{
				{21, 16, 11, 6, 1},
				{22, 17, 12, 7, 2},
				{23, 18, 13, 8, 3},
				{24, 19, 14, 9, 4},
				{25, 20, 15, 10, 5}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.RotateClockwise(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix5x5.RotateClockwise() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix5x5_RotateCounterClockwise(t *testing.T) {
	tests := []struct {
		name string
		m    Matrix5x5
		want Matrix5x5
	}{
		// TODO: Add test cases.
		{name: "Test 1",
			m: *SetMatrix([5][5]StickerIndex{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25}}),
			want: *SetMatrix([5][5]StickerIndex{
				{5, 10, 15, 20, 25},
				{4, 9, 14, 19, 24},
				{3, 8, 13, 18, 23},
				{2, 7, 12, 17, 22},
				{1, 6, 11, 16, 21}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.RotateCounterClockwise(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix5x5.RotateCounterClockwise() = %v, want %v", got, tt.want)
			}
		})
	}
}
