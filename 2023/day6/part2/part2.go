package part2

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func TotalWaysFromInput() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	return TotalWays(lines)
}

func TotalWays(lines []string) int {
	race := ParseRace(lines[0], lines[1])
	return WaysToBeatRecord(race)
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

func ParseRace(timeLine string, distanceLine string) Race {
	timeMatches := numRegex.FindAllString(timeLine, -1)
	distanceMatches := numRegex.FindAllString(distanceLine, -1)

	time, _ := strconv.Atoi(strings.Join(timeMatches, ""))
	distance, _ := strconv.Atoi(strings.Join(distanceMatches, ""))

	return Race{time, distance}
}
