package part1

import (
	"reflect"
	"testing"
)

func TestNumberOfSteps(t *testing.T) {
	tests := []struct {
		name      string
		gridLines []string
		want      int
	}{
		{
			"Simple loop",
			[]string{
				"-L|F7",
				"7S-7|",
				"L|7||",
				"-L-J|",
				"L|-JF",
			},
			4,
		},
		{
			"Longer loop",
			[]string{
				"...F7.",
				"..FJ|.",
				".SJ.L7",
				".|F--J",
				".LJ...",
			},
			8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberOfSteps(tt.gridLines); got != tt.want {
				t.Errorf("NumberOfSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindStart(t *testing.T) {
	tests := []struct {
		name string
		grid []string
		want Coordinates
	}{
		{
			"Simple loop",
			[]string{
				"-L|F7",
				"7S-7|",
				"L|7||",
				"-L-J|",
				"L|-JF",
			},
			Coordinates{1, 1},
		},
		{
			"Longer loop",
			[]string{
				"..F7.",
				".FJ|.",
				"SJ.L7",
				"|F--J",
				"LJ...",
			},
			Coordinates{2, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindStart(tt.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindNextDirection(t *testing.T) {
	type args struct {
		direction   Direction
		currentTile Tile
	}
	tests := []struct {
		name string
		args args
		want Direction
	}{
		{"Vertical moving south",
			args{
				direction:   South,
				currentTile: Tile{true, North, South},
			},
			South,
		},
		{"Vertical moving north",
			args{
				direction:   North,
				currentTile: Tile{true, North, South},
			},
			North,
		},
		{"Horizontal moving east",
			args{
				direction:   East,
				currentTile: Tile{true, East, West},
			},
			East,
		},
		{"Horizontal moving west",
			args{
				direction:   West,
				currentTile: Tile{true, East, West},
			},
			West,
		},
		{"90 degrees south-west moving north",
			args{
				direction:   North,
				currentTile: Tile{true, South, West},
			},
			West,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := FindNextDirection(tt.args.direction, tt.args.currentTile)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindNextMovement() = %v, want %v", got, tt.want)
			}
		})
	}
}
