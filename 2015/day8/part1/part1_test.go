package part1

import (
	"testing"
)

type test struct {
	input    []rune
	expected int
}

var tests = []test{
	{input: []rune(`""`), expected: 2},
	{input: []rune(`"abc"`), expected: 2},
	{input: []rune(`"aaa\"aaa"`), expected: 3},
	{input: []rune(`"\x27"`), expected: 5},
	{input: []rune(`"ecn\x50ooprbstnq"`), expected: 5},
	{input: []rune(`"\xf2\"jdstiwqer\"h"`), expected: 7},
	{input: []rune(`"yd\\"`), expected: 3},
	{input: []rune(`"\"pa\\x\x18od\\emgje\\"`), expected: 9},
	{input: []rune(`"\\x45"`), expected: 3},
}

func TestCalculateDifference(t *testing.T) {
	for _, test := range tests {
		received := CalculateDifference(test.input)

		if received != test.expected {
			t.Errorf(`%v should have been %v, was %v`, string(test.input), test.expected, received)
		}
	}
}
