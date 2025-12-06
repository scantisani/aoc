package part1

import (
	"testing"
)

type sumTest struct {
	input    []string
	expected int
}

var sumTests = []sumTest{
	{[]string{
		"123 328  51 64",
		"45 64  387 23",
		"6 98  215 314",
		"*   +   *   +",
	}, 4277556},
	{[]string{
		"1  2  3  4",
		"6  9  10 1",
		"*  +  *  +",
	}, 52},
}

func TestSumSolutions(t *testing.T) {
	for _, test := range sumTests {
		received := SumSolutions(test.input)
		if received != test.expected {
			t.Errorf(`%v should have been %v, was %v`, test.input, test.expected, received)
		}
	}
}
