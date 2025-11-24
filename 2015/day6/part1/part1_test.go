package part1

import (
	"testing"
)

type test struct {
	input    []string
	expected int
}

var tests = []test{
	{[]string{"turn on 0,0 through 999,999"}, 1_000_000},
	{[]string{"turn on 0,0 through 1,1", "toggle 0,0 through 1,1"}, 0},
	{[]string{"toggle 0,0 through 999,0", "turn off 500,0 through 999,0"}, 500},
}

func TestCountLights(t *testing.T) {
	for _, test := range tests {
		received := CountLights(test.input)
		if received != test.expected {
			t.Errorf(`Should have been %d, was %d`, test.expected, received)
		}
	}
}
