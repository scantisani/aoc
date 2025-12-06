package part2

import (
	"testing"
)

type removableTest struct {
	input    []string
	expected int
}

var removableTests = []removableTest{
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
	}, 43},
	{[]string{
		".@.@.",
		"@@@.@",
		".@.@.",
		"@.@@@",
	}, 12},
}

func TestCountRemovableRolls(t *testing.T) {
	for _, test := range removableTests {
		received := CountRemovableRolls(test.input)
		if received != test.expected {
			t.Errorf(`%v should have been %v, was %v`, test.input, test.expected, received)
		}
	}
}
