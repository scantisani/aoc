package part2

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
	return TotalRibbon(lines)
}

type Present struct {
	length int
	width  int
	height int
}

func TotalRibbon(lines []string) int {
	sum := 0

	for _, line := range lines {
		present := ParsePresent(line)
		ribbonForPresent := RibbonPresent(present)
		sum += ribbonForPresent
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

func RibbonPresent(present Present) int {
	front := (present.width * 2) + (present.height * 2)
	side := (present.length * 2) + (present.height * 2)
	top := (present.length * 2) + (present.width * 2)

	smallestSide := min(front, side, top)

	cubicVolume := present.width * present.height * present.length

	return smallestSide + cubicVolume
}
