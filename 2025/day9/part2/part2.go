package part2

import (
	"log"
	"os"
	"slices"
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

type Position struct {
	x, y int
}

type ColourGrid struct {
	reds []Position
	xs   map[int][]int
	ys   map[int][]int
}

func LargestRectangle(lines []string) int {
	grid := ParseGrid(lines)
	reds := grid.reds

	maximum := 0

	for i, red1 := range reds {
		if i+2 > len(reds) {
			continue
		}

		for _, red2 := range reds[i+2:] {
			if grid.IsValidRectangle(red1, red2) {
				area := rectangleArea(red1, red2)
				maximum = max(maximum, area)
			}
		}
	}

	return maximum
}

func ParseGrid(lines []string) ColourGrid {
	grid := ColourGrid{xs: map[int][]int{}, ys: map[int][]int{}}

	firstRed := parsePosition(lines[0])
	currentRed := firstRed

	for _, line := range lines[1:] {
		nextRed := parsePosition(line)

		grid.addRed(currentRed, nextRed)
		currentRed = nextRed
	}
	grid.addRed(currentRed, firstRed)

	return grid
}

func parsePosition(line string) Position {
	splits := strings.Split(line, ",")
	x, _ := strconv.Atoi(splits[0])
	y, _ := strconv.Atoi(splits[1])
	return Position{x, y}
}

func rectangleArea(position1, position2 Position) int {
	vertical := absoluteValue(position1.x-position2.x) + 1
	horizontal := absoluteValue(position1.y-position2.y) + 1

	return vertical * horizontal
}

func absoluteValue(number int) int {
	if number < 0 {
		return -number
	}

	return number
}

func (cg *ColourGrid) addRed(previousRed, nextRed Position) {
	cg.reds = append(cg.reds, nextRed)
	cg.setPosition(Position{nextRed.x, nextRed.y})
	cg.drawGreenLine(nextRed, previousRed)
}

func (cg *ColourGrid) drawGreenLine(nextRed Position, previousRed Position) {
	if previousRed.x == nextRed.x {
		var greaterY, lesserY int
		if previousRed.y > nextRed.y {
			greaterY, lesserY = previousRed.y, nextRed.y
		} else {
			greaterY, lesserY = nextRed.y, previousRed.y
		}

		for y := lesserY + 1; y < greaterY; y++ {
			cg.setPosition(Position{nextRed.x, y})
		}
	} else if previousRed.y == nextRed.y {
		var greaterX, lesserX int
		if previousRed.x > nextRed.x {
			greaterX, lesserX = previousRed.x, nextRed.x
		} else {
			greaterX, lesserX = nextRed.x, previousRed.x
		}

		for x := lesserX + 1; x < greaterX; x++ {
			cg.setPosition(Position{x, nextRed.y})
		}
	}
}

func (cg *ColourGrid) IsValidRectangle(position1 Position, position2 Position) bool {
	var lesserX, greaterX int
	var lesserY, greaterY int

	if position1.x < position2.x {
		lesserX, greaterX = position1.x+1, position2.x-1
	} else {
		lesserX, greaterX = position2.x+1, position1.x-1
	}

	if position1.y < position2.y {
		lesserY, greaterY = position1.y+1, position2.y-1
	} else {
		lesserY, greaterY = position2.y+1, position1.y-1
	}

	// any walls just inside left edge?
	if cg.elementInRange(cg.ys[lesserX], lesserY, greaterY) {
		return false
	}

	// any walls just inside right edge?
	if cg.elementInRange(cg.ys[greaterX], lesserY, greaterY) {
		return false
	}

	// any walls just inside bottom edge?
	if cg.elementInRange(cg.xs[lesserY], lesserX, greaterX) {
		return false
	}

	// any walls just inside top edge?
	if cg.elementInRange(cg.xs[greaterY], lesserX, greaterX) {
		return false
	}

	return true
}

func (cg *ColourGrid) elementInRange(elements []int, lowerBound, upperBound int) bool {
	return slices.ContainsFunc(elements, func(i int) bool {
		return i >= lowerBound && i <= upperBound
	})
}

func (cg *ColourGrid) setPosition(position Position) {
	cg.xs[position.y] = append(cg.xs[position.y], position.x)
	cg.ys[position.x] = append(cg.ys[position.x], position.y)
}
