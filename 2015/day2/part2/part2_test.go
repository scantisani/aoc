package part2

import (
	"testing"
)

type ribbonTest struct {
	input    Present
	expected int
}

var solveTests = []ribbonTest{
	{Present{2, 3, 4}, 34},
	{Present{1, 1, 10}, 14},
}

func TestRibbon(t *testing.T) {
	for _, test := range solveTests {
		ribbon := RibbonPresent(test.input)
		if ribbon != test.expected {
			t.Errorf(`Should have been %d, was %d`, test.expected, ribbon)
		}
	}
}
