package part2

import (
	"testing"
)

type calibrationSumTest struct {
	input    []string
	expected int
}

var calibrationSumTests = []calibrationSumTest{
	{[]string{"1abc2", "pqr3two8vwx", "a1b2c3d4e5six", "treb7uchet"}, 143},
	{[]string{"2a3", "one9aaa"}, 42},
	{[]string{"nodigits", "lots4and8lots9of4digits7", "trebu78et"}, 125},
	{
		[]string{
			"two1nine",
			"eightwothree",
			"abcone2threexyz",
			"xtwone3four",
			"4nineeightseven2",
			"zoneight234",
			"7pqrstsixteen",
		},
		281,
	},
}

func TestCalibrationSum(t *testing.T) {
	for _, test := range calibrationSumTests {
		calibration := CalibrationSum(test.input)
		if calibration != test.expected {
			t.Errorf(`Should have been %d, was %d`, test.expected, calibration)
		}
	}
}

type calibrationTest struct {
	input    string
	expected int
}

var calibrationTests = []calibrationTest{
	{"two1nine", 29},
	{"eightwothree", 83},
	{"abcone2threexyz", 13},
	{"xtwone3four", 24},
	{"4nineeightseven2", 42},
	{"zoneight234", 14},
	{"7pqrstsixteen", 76},
	{"twoeightfcrnmvvbrd93threetkhklbcxdptfq", 23},
	{"four7five9nd31", 41},
	{"three1frglrrm435dsqbxxtrj7", 37},
	{"twothree78fpghbvq7jfsjsqnd", 27},
	{"hvmbmqnxk4onesix29kdhrdqtcfx1znmjhfjx", 41},
	{"3twoeighteightfivepztpjsbcrfour", 34},
	{"16stctmrmj3threeninepdsxb", 19},
	{"seven7nhrtgnltntgfzb", 77},
	{"one1onermlsevenseven", 17},
	{"17fkg", 17},
	{"53ldplzx", 53},
	{"1abc2", 12},
	{"pqr3stu8vwx", 38},
	{"a1b2c3d4e5f", 15},
	{"treb7uchet", 77},
	{"2a3", 23},
	{"aaa9aaa", 99},
	{"nodigits", 0},
	{"lots4and8lots9of4digits7", 47},
}

func TestCalibration(t *testing.T) {
	for _, test := range calibrationTests {
		calibration := Calibration(test.input)
		if calibration != test.expected {
			t.Errorf(`Input %q should have been %d, was %d`, test.input, test.expected, calibration)
		}
	}
}
