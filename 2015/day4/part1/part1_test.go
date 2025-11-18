package part1

import (
	"testing"
)

type mineCoinTest struct {
	input    string
	expected int
}

var mineCoinTests = []mineCoinTest{
	{"abcdef", 609043},
	{"pqrstuv", 1048970},
}

func TestMineAdventCoin(t *testing.T) {
	for _, test := range mineCoinTests {
		received := MineAdventCoin(test.input)
		if received != test.expected {
			t.Errorf(`Should have been %d, was %d`, test.expected, received)
		}
	}
}
