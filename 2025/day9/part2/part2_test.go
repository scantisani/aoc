package part2

import (
	"cmp"
	"maps"
	"slices"
	"testing"
)

type test struct {
	input    []string
	expected int
}

var tests = []test{
	{[]string{
		"7,1",
		"11,1",
		"11,7",
		"9,7",
		"9,5",
		"2,5",
		"2,3",
		"7,3",
	}, 24},
	{[]string{
		"7,1",
		"11,1",
		"11,7",
		"9,7",
		"9,5",
		"2,5",
		"2,3",
		"7,3",
	}, 24},
}

func TestLargestRectangle(t *testing.T) {
	for _, test := range tests {
		received := LargestRectangle(test.input)
		if received != test.expected {
			t.Errorf(`%v should have been %v, was %v`, test.input, test.expected, received)
		}
	}
}

type parseGridTest struct {
	input    []string
	expected ColourGrid
}

var exampleGrid = ColourGrid{
	reds: []Position{
		{11, 1},
		{11, 7},
		{9, 7},
		{9, 5},
		{2, 5},
		{2, 3},
		{7, 3},
		{7, 1},
	},
	xs: map[int][]int{
		1: {7, 8, 9, 10, 11},
		2: {7, 11},
		3: {2, 3, 4, 5, 6, 7, 11},
		4: {2, 11},
		5: {2, 3, 4, 5, 6, 7, 8, 9, 11},
		6: {9, 11},
		7: {10, 11, 9},
	},
	ys: map[int][]int{
		2:  {3, 4, 5},
		3:  {3, 5},
		4:  {3, 5},
		5:  {3, 5},
		6:  {3, 5},
		7:  {1, 2, 3, 5},
		8:  {1, 5},
		9:  {1, 5, 6, 7},
		10: {1, 7},
		11: {1, 2, 3, 4, 5, 6, 7},
	},
}

var gridTests = []parseGridTest{
	{[]string{
		"7,1",
		"11,1",
		"11,7",
		"9,7",
		"9,5",
		"2,5",
		"2,3",
		"7,3",
	}, exampleGrid},
	{[]string{
		"1,1",
		"3,1",
		"3,4",
		"1,4",
	}, ColourGrid{
		reds: []Position{
			{3, 1},
			{3, 4},
			{1, 4},
			{1, 1},
		},
		xs: map[int][]int{
			1: {1, 2, 3},
			2: {1, 3},
			3: {1, 3},
			4: {1, 2, 3},
		},
		ys: map[int][]int{
			1: {1, 2, 3, 4},
			2: {1, 4},
			3: {1, 2, 3, 4},
		},
	}},
}

func unorderedEqual[A cmp.Ordered](slice1, slice2 []A) bool {
	slices.Sort(slice1)
	slices.Sort(slice2)

	return slices.Equal(slice1, slice2)
}

func gridsEqual(grid1, grid2 ColourGrid) bool {
	return slices.Equal(grid1.reds, grid2.reds) &&
		maps.EqualFunc(grid1.xs, grid2.xs, unorderedEqual) &&
		maps.EqualFunc(grid1.ys, grid2.ys, unorderedEqual)
}

func TestParseGrid(t *testing.T) {
	for _, test := range gridTests {
		received := ParseGrid(test.input)
		if !gridsEqual(received, test.expected) {
			t.Errorf(`%v reds should have been: %v, was: %v`, test.input, test.expected, received)
		}
	}
}

type validRectangleTest struct {
	position1, position2 Position
	expected             bool
}

var validRectangleTests = []validRectangleTest{
	{Position{7, 3}, Position{11, 1}, true},
	{Position{9, 7}, Position{9, 5}, true},
	{Position{9, 5}, Position{2, 3}, true},
	{Position{2, 3}, Position{11, 1}, false},
	{Position{2, 3}, Position{9, 7}, false},
	{Position{2, 5}, Position{11, 7}, false},
}

func TestValidRectangle(t *testing.T) {
	for _, test := range validRectangleTests {
		received := exampleGrid.IsValidRectangle(test.position1, test.position2)
		if received != test.expected {
			t.Errorf(`%v, %v should have been %v, was %v`, test.position1, test.position2, test.expected, received)
		}
	}
}
