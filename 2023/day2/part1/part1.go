package part1

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var cubeMaxes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func PossibleGamesFromInput() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	return PossibleGamesSum(lines[:100]) // don't include the last (blank) split line
}

func PossibleGamesSum(games []string) int {
	sum := 0

	for _, id := range PossibleGames(games) {
		sum += id
	}

	return sum
}

var gameRegex = regexp.MustCompile(`Game (\d+): (.*)`)

func PossibleGames(games []string) []int {
	var possibleGames []int

	for _, game := range games {
		matches := gameRegex.FindStringSubmatch(game)

		id, _ := strconv.Atoi(matches[1])
		cubeCounts := matches[2]

		if GamePossible(cubeCounts) {
			possibleGames = append(possibleGames, id)
		}
	}

	return possibleGames
}

var cubeRegex = regexp.MustCompile(`(\d+) (red|green|blue)`)

func GamePossible(gameDesc string) bool {
	rounds := strings.Split(gameDesc, ";")
	for _, round := range rounds {
		cubeCounts := strings.Split(round, ", ")

		for _, cubeCount := range cubeCounts {
			matches := cubeRegex.FindStringSubmatch(cubeCount)
			count, _ := strconv.Atoi(matches[1])
			colour := matches[2]

			if count > cubeMaxes[colour] {
				return false
			}
		}
	}

	return true
}
