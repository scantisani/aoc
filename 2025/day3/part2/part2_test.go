package part2

import (
	"testing"
)

type joltageTest struct {
	input    string
	expected int
}

var joltageTests = []joltageTest{
	{"987654321111111", 987654321111},
	{"811111111111119", 811111111119},
	{"234234234234278", 434234234278},
	{"818181911112111", 888911112111},
	{"2139999998999", 239999998999},
	{"11123456789111", 123456789111},
}

func TestHighestJoltage(t *testing.T) {
	for _, test := range joltageTests {
		received := HighestJoltage(test.input)
		if received != test.expected {
			t.Errorf(`%v should have been %v, was %v`, test.input, test.expected, received)
		}
	}
}
