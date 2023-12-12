package part1

import (
	"log"
	"os"
	"strings"
)

func NumberOfStepsFromInput() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	return NumberOfSteps(lines[:140]) // skip last (blank) line
}

type Direction int

const (
	North Direction = iota
	East  Direction = iota
	South Direction = iota
	West  Direction = iota
)

type Tile struct {
	isPipe bool
	endA   Direction
	endB   Direction
}

func NumberOfSteps(gridLines []string) int {
	steps := 0

	startingPosition := FindStart(gridLines)
	grid := ParseGrid(gridLines)

	surroundingTiles := surroundingTiles(startingPosition, grid)
	currentDirection := findStartingDirection(surroundingTiles)

	currentPosition := findNextPosition(startingPosition, currentDirection)
	steps++

	for currentPosition != startingPosition {
		currentTile := grid[currentPosition.row][currentPosition.column]

		currentDirection, _ = FindNextDirection(currentDirection, currentTile)
		currentPosition = findNextPosition(currentPosition, currentDirection)

		steps++
	}

	return steps / 2
}

func findNextPosition(currentPosition Coordinates, currentDirection Direction) Coordinates {
	rowChange := 0
	columnChange := 0

	switch currentDirection {
	case North:
		rowChange -= 1
	case East:
		columnChange += 1
	case South:
		rowChange += 1
	case West:
		columnChange -= 1
	}

	return Coordinates{currentPosition.row + rowChange, currentPosition.column + columnChange}
}

func surroundingTiles(currentPosition Coordinates, grid [][]Tile) []Tile {
	return []Tile{
		grid[currentPosition.row-1][currentPosition.column],
		grid[currentPosition.row][currentPosition.column+1],
		grid[currentPosition.row+1][currentPosition.column],
		grid[currentPosition.row][currentPosition.column-1],
	}
}

func findStartingDirection(surroundingTiles []Tile) Direction {
	directions := []Direction{North, East, South, West}

	for i, direction := range directions {
		_, ok := FindNextDirection(direction, surroundingTiles[i])
		if ok {
			return direction
		}
	}

	return 0
}

func ParseGrid(gridLines []string) [][]Tile {
	grid := make([][]Tile, 0)

	for _, row := range gridLines {
		gridRow := make([]Tile, 0)
		for _, element := range []rune(row) {
			tile := parseTile(element)
			gridRow = append(gridRow, tile)
		}

		grid = append(grid, gridRow)
	}

	return grid
}

func parseTile(element rune) Tile {

	switch element {
	case '|':
		return Tile{true, North, South}
	case '-':
		return Tile{true, East, West}
	case 'L':
		return Tile{true, North, East}
	case 'J':
		return Tile{true, North, West}
	case '7':
		return Tile{true, South, West}
	case 'F':
		return Tile{true, South, East}
	default:
		return Tile{false, -1, -1}
	}
}

type Coordinates struct {
	row    int
	column int
}

func FindStart(grid []string) Coordinates {
	coordinates := Coordinates{}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'S' {
				return Coordinates{i, j}
			}
		}
	}

	return coordinates
}

func FindNextDirection(direction Direction, currentTile Tile) (Direction, bool) {
	opposites := map[Direction]Direction{
		North: South,
		East:  West,
		South: North,
		West:  East,
	}

	if currentTile.endA == opposites[direction] {
		return currentTile.endB, true
	} else if currentTile.endB == opposites[direction] {
		return currentTile.endA, true
	} else {
		return -1, false
	}
}
