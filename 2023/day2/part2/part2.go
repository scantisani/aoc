package part2

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func CubePowerSumFromInput() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	return CubePowerSum(lines[:100]) // don't include the last (blank) split line
}

func CubePowerSum(games []string) int {
	sum := 0

	for _, game := range games {
		matches := gameRegex.FindStringSubmatch(game)
		gameDesc := matches[2]

		sum += PowerForGame(gameDesc)
	}

	return sum
}

var gameRegex = regexp.MustCompile(`Game (\d+): (.*)`)

func PowerForGame(gameDesc string) int {
	mins := MinsForGame(gameDesc)

	return mins["red"] * mins["green"] * mins["blue"]
}

var cubeRegex = regexp.MustCompile(`(\d+) (red|green|blue)`)

func MinsForGame(gameDesc string) map[string]int {
	minCubes := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	rounds := strings.Split(gameDesc, ";")
	for _, round := range rounds {
		cubeCounts := strings.Split(round, ", ")

		for _, cubeCount := range cubeCounts {
			matches := cubeRegex.FindStringSubmatch(cubeCount)
			count, _ := strconv.Atoi(matches[1])
			colour := matches[2]

			if count > minCubes[colour] {
				minCubes[colour] = count
			}
		}
	}

	return minCubes
}
