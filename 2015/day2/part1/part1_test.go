package part1

import (
	"testing"
)

type parsePresentTest struct {
	input    string
	expected Present
}

var parsePresentTests = []parsePresentTest{
	{"1x2x3", Present{length: 1, width: 2, height: 3}},
	{"1x1x10", Present{length: 1, width: 1, height: 10}},
}

func TestParsePresent(t *testing.T) {
	for _, test := range parsePresentTests {
		calibration := ParsePresent(test.input)
		if calibration != test.expected {
			t.Errorf(`Should have been %d, was %d`, test.expected, calibration)
		}
	}
}

type wrapPresentTest struct {
	input    Present
	expected int
}

var wrapPresentTests = []wrapPresentTest{
	{Present{2, 3, 4}, 58},
	{Present{1, 1, 10}, 43},
}

func TestWrapPresent(t *testing.T) {
	for _, test := range wrapPresentTests {
		paper := WrapPresent(test.input)
		if paper != test.expected {
			t.Errorf(`Should have been %d, was %d`, test.expected, paper)
		}
	}
}
