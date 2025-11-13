package part1

import (
	"testing"
)

type solveTest struct {
	input    string
	expected int
}

var solveTests = []solveTest{
	{"(())", 0},
	{"()()", 0},
	{"(((", 3},
	{"(()(()(", 3},
	{"))(((((", 3},
	{"())", -1},
	{"))(", -1},
	{")))", -3},
	{")())())", -3},
}

func TestSolve(t *testing.T) {
	for _, test := range solveTests {
		floor := TraverseFloors(test.input)
		if floor != test.expected {
			t.Errorf(`Should have been %d, was %d`, test.expected, floor)
		}
	}
}
