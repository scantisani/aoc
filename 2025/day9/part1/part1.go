package part1

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	return LargestRectangle(lines)
}

type position struct {
	x, y int
}

func LargestRectangle(lines []string) int {
	positions := parsePositions(lines)
	maximum := 0

	for i := range positions {
		for j := range positions[i+1:] {
			area := rectangleArea(positions[i], positions[j])
			maximum = max(maximum, area)
		}
	}

	return maximum
}

func parsePositions(lines []string) []position {
	var positions []position

	for _, line := range lines {
		splits := strings.Split(line, ",")
		x, _ := strconv.Atoi(splits[0])
		y, _ := strconv.Atoi(splits[1])

		positions = append(positions, position{x, y})
	}

	return positions
}

func absoluteValue(number int) int {
	if number >= 0 {
		return number
	} else {
		return -number
	}
}

func rectangleArea(position1, position2 position) int {
	vertical := absoluteValue(position1.x-position2.x) + 1
	horizontal := absoluteValue(position1.y-position2.y) + 1

	return vertical * horizontal
}
