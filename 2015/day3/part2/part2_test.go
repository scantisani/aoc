package part2

import (
	"maps"
	"testing"
)

type visitHousesTest struct {
	input    string
	expected map[Position]int
}

var visitHousesTests = []visitHousesTest{
	{"^v", map[Position]int{
		Position{0, 0}:  1,
		Position{0, 1}:  1,
		Position{0, -1}: 1,
	}},
	{"^>v<", map[Position]int{
		Position{0, 0}: 3,
		Position{0, 1}: 1,
		Position{1, 0}: 1,
	}},
	{"^v^v^v^v^v", map[Position]int{
		Position{0, 0}:  1,
		Position{0, 1}:  1,
		Position{0, 2}:  1,
		Position{0, 3}:  1,
		Position{0, 4}:  1,
		Position{0, 5}:  1,
		Position{0, -1}: 1,
		Position{0, -2}: 1,
		Position{0, -3}: 1,
		Position{0, -4}: 1,
		Position{0, -5}: 1,
	}},
}

func TestVisitHouses(t *testing.T) {
	for _, test := range visitHousesTests {
		houseMap := VisitHouses(test.input)
		if !maps.Equal(houseMap, test.expected) {
			t.Errorf(`Should have been %d, was %d`, test.expected, houseMap)
		}
	}
}

type countHousesTest struct {
	input    string
	expected int
}

var countHousesTests = []countHousesTest{
	{"^v", 3},
	{"^>v<", 3},
	{"^v^v^v^v^v", 11},
}

func TestCountHouses(t *testing.T) {
	for _, test := range countHousesTests {
		countHouses := CountHouses(test.input)
		if countHouses != test.expected {
			t.Errorf(`Should have been %d, was %d`, test.expected, countHouses)
		}
	}
}
