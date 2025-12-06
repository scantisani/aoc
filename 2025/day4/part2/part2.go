package part2

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
	return CountRemovableRolls(lines)
}

type Position struct {
	x, y int
}

type Diagram map[Position]bool

func (d *Diagram) addRoll(position Position) {
	(*d)[position] = true
}

func (d *Diagram) hasRoll(position Position) bool {
	return (*d)[position]
}

func (d *Diagram) accessibleRoll(position Position) bool {
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

func (d *Diagram) removeRolls(positions []Position) {
	for _, position := range positions {
		delete(*d, position)
	}
}

func (d *Diagram) removableRolls() []Position {
	var positions []Position

	for position := range *d {
		if d.accessibleRoll(position) {
			positions = append(positions, position)
		}
	}

	return positions
}

func CountRemovableRolls(lines []string) int {
	diagram := buildDiagram(lines)

	sum := 0
	removableRolls := diagram.removableRolls()

	for len(removableRolls) > 0 {
		sum += len(removableRolls)

		diagram.removeRolls(removableRolls)
		removableRolls = diagram.removableRolls()
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
