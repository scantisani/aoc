package part2

import (
	"testing"
)

type test struct {
	input    []string
	expected int
}

var tests = []test{
	{[]string{"toggle 0,0 through 999,999"}, 2_000_000},
	{[]string{"turn on 0,0 through 0,0", "toggle 0,0 through 0,0"}, 3},
	{[]string{"toggle 0,0 through 999,0", "turn off 500,0 through 999,0"}, 1500},
	{[]string{"turn off 500,0 through 999,0"}, 0},
}

func TestMeasureLights(t *testing.T) {
	for _, test := range tests {
		received := MeasureLights(test.input)
		if received != test.expected {
			t.Errorf(`Should have been %d, was %d`, test.expected, received)
		}
	}
}
