package part2

import (
	"testing"
)

type test struct {
	input    string
	expected bool
}

var tests = []test{
	{"qjhvhtzxzqqjkmpb", true},
	{"xxyxx", true},
	{"uurcxstgmygtbstg", false},
	{"ieodomkazucvgmuy", false},
	{"aaa", false},
	{"aaaa", true},
	{"abab", true},
	{"abcdefef", true},
	{"aaabcdef", false},
}

func TestIsNiceString(t *testing.T) {
	for _, test := range tests {
		received := IsNiceString(test.input)
		if received != test.expected {
			t.Errorf(`Input "%s" should have been %t, was %t`, test.input, test.expected, received)
		}
	}
}
