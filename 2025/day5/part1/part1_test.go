package part1

import (
	"testing"
)

type freshTest struct {
	input    []string
	expected int
}

var freshTests = []freshTest{
	{[]string{
		"3-5",
		"10-14",
		"16-20",
		"12-18",
		"",
		"1",
		"5",
		"8",
		"11",
		"17",
		"32",
	}, 3},
}

func TestCountFreshIngredients(t *testing.T) {
	for _, test := range freshTests {
		received := CountFreshIngredients(test.input)
		if received != test.expected {
			t.Errorf(`%v should have been %v, was %v`, test.input, test.expected, received)
		}
	}
}
