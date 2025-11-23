package part1

import (
	"testing"
)

type test struct {
	input    string
	expected bool
}

var tests = []test{
	{"ugknbfddgicrmopn", true},
	{"aaa", true},
	{"jchzalrnumimnmhp", false},
	{"haegwjzuvuyypxyu", false},
	{"dvszwmarrgswjxmb", false},
}

func TestIsNiceString(t *testing.T) {
	for _, test := range tests {
		received := IsNiceString(test.input)
		if received != test.expected {
			t.Errorf(`Should have been %t, was %t`, test.expected, received)
		}
	}
}
