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
	return TotalWrappingPaper(lines)
}

type Present struct {
	length int
	width  int
	height int
}

func TotalWrappingPaper(lines []string) int {
	sum := 0

	for _, line := range lines {
		present := ParsePresent(line)
		paperForPresent := WrapPresent(present)
		sum += paperForPresent
	}

	return sum
}

func ParsePresent(line string) Present {
	dimensions := strings.Split(line, "x")

	length, err := strconv.Atoi(dimensions[0])
	if err != nil {
		log.Fatal(err)
	}

	width, err := strconv.Atoi(dimensions[1])
	if err != nil {
		log.Fatal(err)
	}

	height, err := strconv.Atoi(dimensions[2])
	if err != nil {
		log.Fatal(err)
	}

	return Present{
		length: length,
		width:  width,
		height: height,
	}
}

func WrapPresent(present Present) int {
	front := present.width * present.height
	side := present.length * present.height
	top := present.length * present.width

	smallestSide := min(front, side, top)

	return (2 * front) + (2 * side) + (2 * top) + smallestSide
}
