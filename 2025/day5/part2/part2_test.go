package part2

import (
	"slices"
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
	}, 14},
	{[]string{
		"1-10",
		"11-13",
		"8-20",
		"25-30",
		"",
	}, 26},
	{[]string{
		"2-4",
		"5-8",
		"12-18",
		"1-16",
		"",
	}, 18},
}

func TestCountFreshIngredients(t *testing.T) {
	for _, test := range freshTests {
		received := CountFreshIngredients(test.input)
		if received != test.expected {
			t.Errorf(`%v should have been %v, was %v`, test.input, test.expected, received)
		}
	}
}

type buildTest struct {
	input    []string
	expected FreshStore
}

var buildTests = []buildTest{
	{[]string{
		"3-5",
		"10-14",
		"16-20",
		"12-18",
	}, FreshStore{
		{3, 5},
		{10, 20},
	}},
	{[]string{
		"1-10",
		"11-13",
		"8-20",
		"25-30",
	}, FreshStore{
		{1, 20},
		{25, 30},
	}},
	{[]string{
		"2-4",
		"5-8",
		"12-18",
		"1-16",
	}, FreshStore{
		{1, 18},
	}},
}

func TestBuild(t *testing.T) {
	for _, test := range buildTests {
		var store FreshStore
		store.Build(test.input)

		if !slices.Equal(store, test.expected) {
			t.Errorf(`%v should have been %v, was %v`, test.input, test.expected, store)
		}
	}
}
