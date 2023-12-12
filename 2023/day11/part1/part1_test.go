package part1

import (
	"reflect"
	"strings"
	"testing"
)

func TestSumOfShortestPaths(t *testing.T) {
	tests := []struct {
		name      string
		gridLines []string
		want      int
	}{
		{
			"AOC example 1",
			[]string{
				"...#......",
				".......#..",
				"#.........",
				"..........",
				"......#...",
				".#........",
				".........#",
				"..........",
				".......#..",
				"#...#.....",
			},
			374,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfShortestPaths(tt.gridLines); got != tt.want {
				t.Errorf("SumOfShortestPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpandUniverse(t *testing.T) {
	tests := []struct {
		name     string
		universe Universe
		want     Universe
	}{
		{
			"AOC example 1",
			Universe{
				[]rune("...#......"),
				[]rune(".......#.."),
				[]rune("#........."),
				[]rune(".........."),
				[]rune("......#..."),
				[]rune(".#........"),
				[]rune(".........#"),
				[]rune(".........."),
				[]rune(".......#.."),
				[]rune("#...#....."),
			},
			Universe{
				[]rune("....#........"),
				[]rune(".........#..."),
				[]rune("#............"),
				[]rune("............."),
				[]rune("............."),
				[]rune("........#...."),
				[]rune(".#..........."),
				[]rune("............#"),
				[]rune("............."),
				[]rune("............."),
				[]rune(".........#..."),
				[]rune("#....#......."),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandUniverse(tt.universe); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandUniverse() = \n%vwant\n%v", printUniverse(got), printUniverse(tt.want))
			}
		})
	}
}

func printUniverse(universe Universe) string {
	builder := strings.Builder{}

	for _, row := range universe {
		builder.WriteString(string(row))
		builder.WriteRune('\n')
	}

	return builder.String()
}

func TestFindShortestPath(t *testing.T) {
	tests := []struct {
		name         string
		coordinatesA Coordinates
		coordinatesB Coordinates
		want         int
	}{
		{
			"AOC Example 1",
			Coordinates{6, 1},
			Coordinates{11, 5},
			9,
		},
		{
			"AOC Example 2",
			Coordinates{0, 4},
			Coordinates{10, 9},
			15,
		},
		{
			"AOC Example 3",
			Coordinates{11, 0},
			Coordinates{11, 5},
			5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindShortestPath(tt.coordinatesA, tt.coordinatesB); got != tt.want {
				t.Errorf("FindShortestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
