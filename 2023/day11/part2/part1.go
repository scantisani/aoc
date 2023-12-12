package part2

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
	return SumOfShortestPaths(lines[:140], 1_000_000) // skip last (blank) line
}

type Coordinates struct {
	row    int
	column int
}

func SumOfShortestPaths(gridLines []string, expansionFactor int) int {
	sumOfShortestPaths := 0

	universe := parseUniverse(gridLines)
	universe = ExpandUniverse(universe)
	galaxies := findGalaxies(universe)

	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			sumOfShortestPaths += FindShortestPath(galaxies[i], galaxies[j], universe, expansionFactor)
		}
	}

	return sumOfShortestPaths
}

func FindShortestPath(coordinatesA Coordinates, coordinatesB Coordinates, universe Universe, expansionFactor int) int {
	distance := 0

	rows := []int{coordinatesA.row, coordinatesB.row}
	slices.Sort(rows)
	for i := rows[0]; i < rows[1]; i++ {
		if universe[i][coordinatesA.column] == 'x' {
			distance += expansionFactor - 1
		} else {
			distance++
		}
	}

	columns := []int{coordinatesA.column, coordinatesB.column}
	slices.Sort(columns)
	for i := columns[0]; i < columns[1]; i++ {
		if universe[coordinatesA.row][i] == 'x' {
			distance += expansionFactor - 1
		} else {
			distance++
		}
	}

	return distance
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

	expandedRow := make([]rune, len(universe[0]))
	for i := range expandedRow {
		expandedRow[i] = 'x'
	}

	for i, location := range locations {
		newUniverse = slices.Insert(newUniverse, location+i+1, expandedRow)
	}

	return newUniverse
}

func insertColumns(locations []int, universe Universe) Universe {
	newUniverse := make(Universe, 0)

	for _, row := range universe {
		newRow := slices.Clone(row)
		for i, location := range locations {
			newRow = slices.Insert(newRow, location+i+1, 'x')
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
