package part2

import (
	"maps"
	"slices"
	"testing"
)

type longestDistanceTest struct {
	input    []string
	expected int
}

var longestDistanceTests = []longestDistanceTest{
	{[]string{
		"London to Dublin = 464",
		"London to Belfast = 518",
		"Dublin to Belfast = 141",
	}, 605},
}

func TestlongestDistance(t *testing.T) {
	for _, test := range longestDistanceTests {
		received := LongestDistance(test.input)

		if received != test.expected {
			t.Errorf(`%v should have been %v, was %v`, test.input, test.expected, received)
		}
	}
}

type parseTest struct {
	input    []string
	expected Graph
}

var parseTests = []parseTest{
	{
		[]string{
			"London to Dublin = 464",
			"London to Belfast = 518",
			"Dublin to Belfast = 141",
		}, Graph{
			"London":  {"Belfast": 518, "Dublin": 464},
			"Dublin":  {"Belfast": 141, "London": 464},
			"Belfast": {"Dublin": 141, "London": 518},
		},
	},
	{
		[]string{
			"London to Dublin = 464",
			"London to Belfast = 518",
			"Dublin to Belfast = 141",
			"Dublin to Brussels = 314",
			"Brussels to London = 263",
			"Brussels to Belfast = 370",
		}, Graph{
			"Belfast":  {"Brussels": 370, "Dublin": 141, "London": 518},
			"Brussels": {"Belfast": 370, "Dublin": 314, "London": 263},
			"Dublin":   {"Belfast": 141, "Brussels": 314, "London": 464},
			"London":   {"Belfast": 517, "Brussels": 263, "Dublin": 464},
		},
	},
}

func TestParseDistance(t *testing.T) {
	for _, test := range parseTests {
		received := ParseDistances(test.input)

		for destination, distances := range test.expected {
			if !maps.Equal(distances, received[destination]) {
				t.Errorf(`%v should have been %v, was %v`, destination, distances, received[destination])
			}
		}
	}
}

type permutationTest struct {
	input    []string
	expected [][]string
}

var permutationTests = []permutationTest{
	{[]string{"London", "Dublin"}, [][]string{{"London", "Dublin"}, {"Dublin", "London"}}},
	{
		[]string{"Belfast", "Dublin", "London"},
		[][]string{
			{"Belfast", "Dublin", "London"},
			{"Dublin", "Belfast", "London"},
			{"Dublin", "London", "Belfast"},
			{"Belfast", "London", "Dublin"},
			{"London", "Belfast", "Dublin"},
			{"London", "Dublin", "Belfast"},
		},
	},
}

func TestPermutations(t *testing.T) {
	for _, test := range permutationTests {
		received := Permutations(test.input)

		if !slices.EqualFunc(received, test.expected, slices.Equal) {
			t.Errorf(`%v should have been %v, was %v`, test.input, test.expected, received)
		}
	}
}
