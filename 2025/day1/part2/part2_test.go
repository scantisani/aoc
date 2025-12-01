package part2

import (
	"testing"
)

type test struct {
	input    []string
	expected int
}

var calibrationSumTests = []test{
	{[]string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}, 6},
	{[]string{
		"L51",  // 99, 1
		"R1",   // 0, 1
		"L1",   // 99
		"R200", // 99, 2
		"L299", // 0, 3
	}, 7},
	{[]string{
		"L550", // 0, 6
		"R224", // 24, 2
		"L24",  // 0, 1
		"R10",  // 10, 0
	}, 9},
}

func TestTotalZeroes(t *testing.T) {
	for _, test := range calibrationSumTests {
		calibration := TotalZeroes(test.input)
		if calibration != test.expected {
			t.Errorf(`%v should have been %v, was %v`, test.input, test.expected, calibration)
		}
	}
}
