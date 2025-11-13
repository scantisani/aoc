package part2

import (
	"testing"
)

type solveTest struct {
	input    string
	expected int
}

var solveTests = []solveTest{
	{")", 1},
	{"()())", 5},
}

func TestSolve(t *testing.T) {
	for _, test := range solveTests {
		index := TraverseFloors(test.input)
		if index != test.expected {
			t.Errorf(`Should have been %d, was %d`, test.expected, index)
		}
	}
}
