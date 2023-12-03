package part2

import (
	"reflect"
	"testing"
)

func TestGearRatioSum(t *testing.T) {
	aocInput := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{"Should correctly handle AOC input", aocInput, 467835},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := GearRatioSum(test.input)
			if got != test.expected {
				t.Errorf("GearRatioSum() = %v, expected %v", got, test.expected)
			}
		})
	}
}

func TestGears(t *testing.T) {
	type args struct {
		topRow    string
		row       string
		bottomRow string
	}

	tests := []struct {
		args     args
		expected []gear
	}{
		{
			args{
				"467..114..",
				"...*......",
				"..35..633.",
			},
			[]gear{{467, 35}},
		},
		{
			args{
				".....114..",
				"..5*......",
				"..35..633.",
			},
			[]gear{{5, 35}},
		},
		{
			args{
				".....114..",
				"..5*10....",
				"..35..633.",
			},
			[]gear{},
		},
		{
			args{
				".....114..",
				"..5*....*.",
				"..35..633.",
			},
			[]gear{{5, 35}, {114, 633}},
		},
	}

	for _, test := range tests {
		got := Gears(test.args.topRow, test.args.row, test.args.bottomRow)
		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("GearRatios(%v) = %v, expected %v", test.args, got, test.expected)
		}
	}
}

func TestNumberPositions(t *testing.T) {
	tests := []struct {
		row      string
		expected []position
	}{
		{
			"467..114..",
			[]position{{0, 2}, {5, 7}},
		},
		{
			"...*......",
			[]position{},
		},
		{
			"..35..633.",
			[]position{{2, 3}, {6, 8}},
		},
	}

	for _, test := range tests {
		got := NumberPositions(test.row)
		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("NumberPositions(%v) = %v, expected %v", test.row, got, test.expected)
		}
	}
}
