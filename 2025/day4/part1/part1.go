package part1

import (
	"log"
	"os"
	"strings"
)

func Solve() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	return CountAccessibleRolls(lines)
}

type Position struct {
	x, y int
}

type Diagram map[Position]bool

func (d Diagram) addRoll(position Position) {
	d[position] = true
}

func (d Diagram) hasRoll(position Position) bool {
	return d[position]
}

func (d Diagram) AccessibleRoll(position Position) bool {
	surroundingPositions := []Position{
		{x: position.x + 1, y: position.y + 1},
		{x: position.x + 1, y: position.y - 1},
		{x: position.x + 1, y: position.y},
		{x: position.x - 1, y: position.y + 1},
		{x: position.x - 1, y: position.y - 1},
		{x: position.x - 1, y: position.y},
		{x: position.x, y: position.y + 1},
		{x: position.x, y: position.y - 1},
	}

	sum := 0

	for _, surroundingPosition := range surroundingPositions {
		if d.hasRoll(surroundingPosition) {
			sum++
		}
	}

	return sum < 4
}

func CountAccessibleRolls(lines []string) int {
	diagram := buildDiagram(lines)

	sum := 0
	for position := range diagram {
		if diagram.AccessibleRoll(position) {
			sum++
		}
	}
	return sum
}

func buildDiagram(lines []string) Diagram {
	diagram := Diagram{}

	for i, line := range lines {
		for j, char := range []rune(line) {
			if char == '@' {
				diagram.addRoll(Position{i, j})
			}
		}
	}

	return diagram
}
