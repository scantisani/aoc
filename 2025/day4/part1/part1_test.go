package part1

import (
	"testing"
)

type accessibleTest struct {
	input    []string
	expected int
}

var accessibleTests = []accessibleTest{
	{[]string{
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
		"@.@@@@..@.",
		"@@.@@@@.@@",
		".@@@@@@@.@",
		".@.@.@.@@@",
		"@.@@@.@@@@",
		".@@@@@@@@.",
		"@.@.@@@.@.",
	}, 13},
	{[]string{
		".@.@.",
		"@@@.@",
		".@.@.",
		"@.@@@",
	}, 8},
}

func TestCountAccessibleRolls(t *testing.T) {
	for _, test := range accessibleTests {
		received := CountAccessibleRolls(test.input)
		if received != test.expected {
			t.Errorf(`%v should have been %v, was %v`, test.input, test.expected, received)
		}
	}
}
