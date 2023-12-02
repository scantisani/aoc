package part1

import (
	"testing"
)

type calibrationSumTest struct {
	input    []string
	expected int
}

var calibrationSumTests = []calibrationSumTest{
	{[]string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}, 142},
	{[]string{"2a3", "aaa9aaa"}, 122},
	{[]string{"nodigits", "lots4and8lots9of4digits7", "trebu78et"}, 125},
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
	{"twoeightfcrnmvvbrd93threetkhklbcxdptfq", 93},
	{"four7five9nd31", 71},
	{"three1frglrrm435dsqbxxtrj7", 17},
	{"twothree78fpghbvq7jfsjsqnd", 77},
	{"hvmbmqnxk4onesix29kdhrdqtcfx1znmjhfjx", 41},
	{"3twoeighteightfivepztpjsbcrfour", 33},
	{"16stctmrmj3threeninepdsxb", 13},
	{"seven7nhrtgnltntgfzb", 77},
	{"one1onermlsevenseven", 11},
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
			t.Errorf(`Should have been %d, was %d`, test.expected, calibration)
		}
	}
}
