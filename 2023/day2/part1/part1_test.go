package part1

import (
	"reflect"
	"testing"
)

type possibleGamesSumTest struct {
	games    []string
	expected int
}

var possibleGamesSumTests = []possibleGamesSumTest{
	{
		[]string{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
		},
		8,
	},
	{
		[]string{
			"Game 11: 8 red, 4 blue, 1 green; 3 red; 1 green; 2 green, 3 blue",
			"Game 12: 10 red, 2 green, 4 blue; 4 red, 2 green; 1 blue, 1 red, 1 green; 10 red, 1 green, 5 blue",
			"Game 13: 20 blue, 9 green, 7 red; 13 red, 13 blue, 16 green; 17 blue, 6 red, 6 green; 1 red, 1 blue, 9 green; 9 blue, 18 green, 7 red",
			"Game 14: 6 blue, 14 red; 9 red, 8 blue; 2 red, 1 green, 8 blue; 3 blue, 1 green, 9 red; 8 blue, 2 green, 1 red",
		},
		23,
	},
}

func TestPossibleGamesSum(t *testing.T) {
	for _, test := range possibleGamesSumTests {
		sum := PossibleGamesSum(test.games)
		if sum != test.expected {
			t.Errorf(`Output for %v should have been %v, was %v`, test.games, test.expected, sum)
		}
	}
}

type possibleGamesTest struct {
	games    []string
	expected []int
}

var possibleGamesTests = []possibleGamesTest{
	{
		[]string{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
		},
		[]int{1, 2, 5},
	},
	{
		[]string{
			"Game 11: 8 red, 4 blue, 1 green; 3 red; 1 green; 2 green, 3 blue",
			"Game 12: 10 red, 2 green, 4 blue; 4 red, 2 green; 1 blue, 1 red, 1 green; 10 red, 1 green, 5 blue",
			"Game 13: 20 blue, 9 green, 7 red; 13 red, 13 blue, 16 green; 17 blue, 6 red, 6 green; 1 red, 1 blue, 9 green; 9 blue, 18 green, 7 red",
			"Game 14: 6 blue, 14 red; 9 red, 8 blue; 2 red, 1 green, 8 blue; 3 blue, 1 green, 9 red; 8 blue, 2 green, 1 red",
		},
		[]int{11, 12},
	},
}

func TestPossibleGames(t *testing.T) {
	for _, test := range possibleGamesTests {
		ids := PossibleGames(test.games)
		if !reflect.DeepEqual(ids, test.expected) {
			t.Errorf(`Output for %v should have been %v, was %v`, test.games, test.expected, ids)
		}
	}
}

type gamePossibleTest struct {
	cubeCounts string
	expected   bool
}

var gamePossibleTests = []gamePossibleTest{
	{"3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", true},
	{"1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", true},
	{"8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", false},
	{"1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", false},
	{"6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", true},
}

func TestGamePossible(t *testing.T) {
	for _, test := range gamePossibleTests {
		possible := GamePossible(test.cubeCounts)
		if possible != test.expected {
			t.Errorf(`Output for %q should have been %t, was %t`, test.cubeCounts, test.expected, possible)
		}
	}
}
