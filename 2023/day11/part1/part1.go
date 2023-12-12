package part1

import (
	"log"
	"os"
	"slices"
	"strings"
)

func SumOfShortestPathsFromInput() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	return SumOfShortestPaths(lines[:140]) // skip last (blank) line
}

type Coordinates struct {
	row    int
	column int
}

func SumOfShortestPaths(gridLines []string) int {
	sumOfShortestPaths := 0

	universe := parseUniverse(gridLines)
	universe = ExpandUniverse(universe)
	galaxies := findGalaxies(universe)

	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			sumOfShortestPaths += FindShortestPath(galaxies[i], galaxies[j])
		}
	}

	return sumOfShortestPaths
}

func FindShortestPath(coordinatesA Coordinates, coordinatesB Coordinates) int {
	return diff(coordinatesA.row, coordinatesB.row) + diff(coordinatesA.column, coordinatesB.column)
}

func diff(a, b int) int {
	diff := a - b

	if diff < 0 {
		return diff * -1
	} else {
		return diff
	}
}

func findGalaxies(universe Universe) []Coordinates {
	coordinates := make([]Coordinates, 0)

	for i := range universe {
		for j := range universe[i] {
			if universe[i][j] == '#' {
				coordinates = append(coordinates, Coordinates{i, j})
			}
		}
	}

	return coordinates
}

func parseUniverse(gridLines []string) Universe {
	universe := make(Universe, 0)
	for _, line := range gridLines {
		universe = append(universe, []rune(line))
	}
	return universe
}

type Universe [][]rune

func ExpandUniverse(universe Universe) Universe {
	emptyRowLocations := make([]int, 0)
	emptyColumnLocations := make([]int, 0)

	for i, row := range universe {
		if isEmpty(row) {
			emptyRowLocations = append(emptyRowLocations, i)
		}
	}
	for i := 0; i < len(universe[0]); i++ {
		column := extractColumn(i, universe)
		if isEmpty(column) {
			emptyColumnLocations = append(emptyColumnLocations, i)
		}
	}

	newUniverse := insertColumns(emptyColumnLocations, universe)
	newUniverse = insertRows(emptyRowLocations, newUniverse)

	return newUniverse
}

func insertRows(locations []int, universe Universe) Universe {
	newUniverse := slices.Clone(universe)

	emptyRow := make([]rune, len(universe[0]))
	for i := range emptyRow {
		emptyRow[i] = '.'
	}

	for i, location := range locations {
		newUniverse = slices.Insert(newUniverse, location+i+1, emptyRow)
	}

	return newUniverse
}

func insertColumns(locations []int, universe Universe) Universe {
	newUniverse := make(Universe, 0)

	for _, row := range universe {
		newRow := slices.Clone(row)
		for i, location := range locations {
			newRow = slices.Insert(newRow, location+i+1, '.')
		}

		newUniverse = append(newUniverse, newRow)
	}

	return newUniverse
}

func extractColumn(columnNumber int, universe Universe) []rune {
	column := make([]rune, 0)

	for _, row := range universe {
		column = append(column, row[columnNumber])
	}

	return column
}

func isEmpty(row []rune) bool {
	for _, element := range row {
		if element != '.' {
			return false
		}
	}

	return true
}
