package part1

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
	}, 3},
	{[]string{
		"L50",
		"R1",
		"L1",
		"R200",
		"L201",
	}, 3},
	{[]string{
		"L550",
		"R224",
		"L24",
		"R10",
	}, 2},
}

func TestTotalZeroes(t *testing.T) {
	for _, test := range calibrationSumTests {
		calibration := TotalZeroes(test.input)
		if calibration != test.expected {
			t.Errorf(`%v should have been %v, was %v`, test.input, test.expected, calibration)
		}
	}
}
