package part2

import (
	"reflect"
	"testing"
)

type cubePowerSumTest struct {
	games    []string
	expected int
}

var cubePowerSumTests = []cubePowerSumTest{
	{
		[]string{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
		},
		2286,
	},
}

func TestCubePowerSum(t *testing.T) {
	for _, test := range cubePowerSumTests {
		power := CubePowerSum(test.games)
		if !reflect.DeepEqual(power, test.expected) {
			t.Errorf(`Output for %q should have been %v, was %v`, test.games, test.expected, power)
		}
	}
}

type powerForGameTest struct {
	gameDesc string
	expected int
}

var powerForGameTests = []powerForGameTest{
	{"3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 48},
	{"1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 12},
	{"8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 1560},
	{"1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 630},
	{"6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 36},
}

func TestPowerForGame(t *testing.T) {
	for _, test := range powerForGameTests {
		power := PowerForGame(test.gameDesc)
		if !reflect.DeepEqual(power, test.expected) {
			t.Errorf(`Output for %q should have been %v, was %v`, test.gameDesc, test.expected, power)
		}
	}
}

type minsForGameTest struct {
	gameDesc string
	expected map[string]int
}

var minsForGameTests = []minsForGameTest{
	{"3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", map[string]int{"red": 4, "green": 2, "blue": 6}},
	{"1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", map[string]int{"red": 1, "green": 3, "blue": 4}},
	{"8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", map[string]int{"red": 20, "green": 13, "blue": 6}},
	{"1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", map[string]int{"red": 14, "green": 3, "blue": 15}},
	{"6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", map[string]int{"red": 6, "green": 3, "blue": 2}},
}

func TestMinsForGame(t *testing.T) {
	for _, test := range minsForGameTests {
		mins := MinsForGame(test.gameDesc)
		if !reflect.DeepEqual(mins, test.expected) {
			t.Errorf(`Output for %q should have been %v, was %v`, test.gameDesc, test.expected, mins)
		}
	}
}
