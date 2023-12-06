package part1

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func WaysMultipliedFromInput() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	return WaysMultiplied(lines)
}

func WaysMultiplied(lines []string) int {
	totalWays := 1

	races := ParseRaces(lines[0], lines[1])
	for _, race := range races {
		ways := WaysToBeatRecord(race)
		totalWays = totalWays * ways
	}

	return totalWays
}

type Race struct {
	time           int
	distanceRecord int
}

func WaysToBeatRecord(race Race) int {
	ways := 0

	for i := 0; i < race.time-1; i++ {
		totalDistance := totalDistance(i, race.time-i)
		if totalDistance > race.distanceRecord {
			ways++
		}
	}

	return ways
}

func totalDistance(speed int, timeRemaining int) int {
	return speed * timeRemaining
}

var numRegex = regexp.MustCompile(`\d+`)

func ParseRaces(timeLine string, distanceLine string) []Race {
	timeMatches := numRegex.FindAllString(timeLine, -1)
	distanceMatches := numRegex.FindAllString(distanceLine, -1)

	races := make([]Race, 0)
	for i := range timeMatches {
		time, _ := strconv.Atoi(timeMatches[i])
		distance, _ := strconv.Atoi(distanceMatches[i])

		race := Race{time, distance}
		races = append(races, race)
	}

	return races
}
