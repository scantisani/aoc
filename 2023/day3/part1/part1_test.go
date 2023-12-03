package part1

import (
	"reflect"
	"testing"
)

func TestPartNumberSum(t *testing.T) {
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
		{"Should correctly handle AOC input", aocInput, 4361},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := PartNumberSum(test.input)
			if got != test.expected {
				t.Errorf("PartNumberSum() = %v, expected %v", got, test.expected)
			}
		})
	}
}

func TestPartNumbers(t *testing.T) {
	type args struct {
		topRow    []rune
		row       string
		bottomRow []rune
	}

	tests := []struct {
		args     args
		expected []int
	}{
		{
			args{
				[]rune("......#..."),
				"617*......",
				[]rune(".....+.58."),
			},
			[]int{617},
		},
		{
			args{
				[]rune(".....+.58."),
				"..592.....",
				[]rune("......755."),
			},
			[]int{592},
		},
		{
			args{
				[]rune("...$.*...."),
				".664.598..",
				[]rune("..35..633."),
			},
			[]int{664, 598},
		},
	}
	for _, test := range tests {
		got := PartNumbers(test.args.topRow, test.args.row, test.args.bottomRow)
		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("NumberPositions() = %v, expected %v", got, test.expected)
		}
	}
}

func TestContainsSymbol(t *testing.T) {
	tests := []struct {
		runes    []rune
		expected bool
	}{
		{
			[]rune("467..114.."),
			false,
		},
		{
			[]rune("...*......"),
			true,
		},
		{
			[]rune("...$.*...."),
			true,
		},
		{
			[]rune("../.."),
			true,
		},
	}
	for _, test := range tests {
		got := ContainsSymbol(test.runes)
		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("NumberPositions(%v) = %v, expected %v", string(test.runes), got, test.expected)
		}
	}
}

func TestNumberPositions(t *testing.T) {
	tests := []struct {
		row      string
		expected [][]int
	}{
		{
			"467..114..",
			[][]int{{0, 2}, {5, 7}},
		},
		{
			"...*......",
			[][]int{},
		},
		{
			"..35..633.",
			[][]int{{2, 3}, {6, 8}},
		},
	}
	for _, test := range tests {
		got := NumberPositions(test.row)
		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("NumberPositions() = %v, expected %v", got, test.expected)
		}
	}
}
