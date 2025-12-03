package part1

import (
	"testing"
)

type joltageTest struct {
	input    string
	expected int
}

var joltageTests = []joltageTest{
	{"987654321111111", 98},
	{"811111111111119", 89},
	{"234234234234278", 78},
	{"818181911112111", 92},
}

func TestHighestJoltage(t *testing.T) {
	for _, test := range joltageTests {
		received := HighestJoltage(test.input)
		if received != test.expected {
			t.Errorf(`%v should have been %v, was %v`, test.input, test.expected, received)
		}
	}
}
