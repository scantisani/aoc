package part1

import (
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
	}, 50},
}

func TestLargestRectangle(t *testing.T) {
	for _, test := range tests {
		received := LargestRectangle(test.input)
		if received != test.expected {
			t.Errorf(`%v should have been %v, was %v`, test.input, test.expected, received)
		}
	}
}
