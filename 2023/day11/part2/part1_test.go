package part2

import (
	"reflect"
	"strings"
	"testing"
)

func TestSumOfShortestPaths(t *testing.T) {
	gridLines := []string{
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
	}

	tests := []struct {
		name            string
		gridLines       []string
		expansionFactor int
		want            int
	}{
		{
			"AOC example 1",
			gridLines,
			10,
			1030,
		},
		{
			"AOC example 2",
			gridLines,
			100,
			8410,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfShortestPaths(tt.gridLines, tt.expansionFactor); got != tt.want {
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
				[]rune("...x#..x...x."),
				[]rune("...x...x.#.x."),
				[]rune("#..x...x...x."),
				[]rune("...x...x...x."),
				[]rune("xxxxxxxxxxxxx"),
				[]rune("...x...x#..x."),
				[]rune(".#.x...x...x."),
				[]rune("...x...x...x#"),
				[]rune("...x...x...x."),
				[]rune("xxxxxxxxxxxxx"),
				[]rune("...x...x.#.x."),
				[]rune("#..x.#.x...x."),
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
	universe := Universe{
		[]rune("...x#..x...x."),
		[]rune("...x...x.#.x."),
		[]rune("#..x...x...x."),
		[]rune("...x...x...x."),
		[]rune("xxxxxxxxxxxxx"),
		[]rune("...x...x#..x."),
		[]rune(".#.x...x...x."),
		[]rune("...x...x...x#"),
		[]rune("...x...x...x."),
		[]rune("xxxxxxxxxxxxx"),
		[]rune("...x...x.#.x."),
		[]rune("#..x.#.x...x."),
	}

	tests := []struct {
		name          string
		coordinatesA  Coordinates
		coordinatesB  Coordinates
		universe      Universe
		expansionSize int
		want          int
	}{
		{
			"AOC Example 1",
			Coordinates{6, 1},
			Coordinates{11, 5},
			universe,
			1_000_000,
			2_000_005,
		},
		{
			"AOC Example 2",
			Coordinates{0, 4},
			Coordinates{10, 9},
			universe,
			1_000_000,
			3_000_009,
		},
		{
			"AOC Example 3",
			Coordinates{11, 0},
			Coordinates{11, 5},
			universe,
			1_000_000,
			1_000_003,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindShortestPath(tt.coordinatesA, tt.coordinatesB, tt.universe, tt.expansionSize); got != tt.want {
				t.Errorf("FindShortestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
