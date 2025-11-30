package part2

import (
	"testing"
)

type differenceTest struct {
	input    string
	expected int
}

var differenceTests = []differenceTest{
	{input: `""`, expected: 4},
	{input: `"abc"`, expected: 4},
	{input: `"aaa\"aaa"`, expected: 6},
	{input: `"\x27"`, expected: 5},
}

func TestCalculateDifference(t *testing.T) {
	for _, test := range differenceTests {
		received := CalculateDifference(test.input)

		if received != test.expected {
			t.Errorf(`%v should have been %v, was %v`, string(test.input), test.expected, received)
		}
	}
}

type conversionTest struct {
	input    string
	expected string
}

var conversionTests = []conversionTest{
	{input: `""`, expected: `"\"\""`},
	{input: `"abc"`, expected: `"\"abc\""`},
	{input: `"aaa\"aaa"`, expected: `"\"aaa\\\"aaa\""`},
	{input: `"\x27"`, expected: `"\"\\x27\""`},
}

func TestConvertString(t *testing.T) {
	for _, test := range conversionTests {
		received := ConvertString(test.input)

		if received != test.expected {
			t.Errorf(`%v should have been %v, was %v`, test.input, test.expected, received)
		}
	}
}
