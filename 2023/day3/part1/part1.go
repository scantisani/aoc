package part1

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func PartNumberSumFromInput() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	return PartNumberSum(lines[:140]) // don't include the last (blank) split line
}

func PartNumberSum(rows []string) int {
	sum := 0

	for i, row := range rows {
		var topRow = make([]rune, len(row))
		var bottomRow = make([]rune, len(row))

		if i == 0 {
			for j := range topRow {
				topRow[j] = '.'
			}
		} else {
			topRow = []rune(rows[i-1])
		}

		if i == len(row)-1 {
			for j := range bottomRow {
				bottomRow[j] = '.'
			}
		} else {
			bottomRow = []rune(rows[i+1])
		}

		numbers := PartNumbers(topRow, row, bottomRow)
		for _, number := range numbers {
			sum += number
		}
	}

	return sum
}

func PartNumbers(topRow []rune, row string, bottomRow []rune) []int {
	partNumbers := make([]int, 0)

	positions := NumberPositions(row)
	for _, position := range positions {
		surroundingRunes := make([]rune, 0)

		start := position[0]

		if start > 0 {
			surroundingRunes = append(surroundingRunes, topRow[start-1])
			surroundingRunes = append(surroundingRunes, []rune(row)[start-1])
			surroundingRunes = append(surroundingRunes, bottomRow[start-1])
		}

		surroundingRunes = append(surroundingRunes, topRow[start])
		surroundingRunes = append(surroundingRunes, bottomRow[start])

		if start < len(row)-1 {
			surroundingRunes = append(surroundingRunes, topRow[start+1])
			surroundingRunes = append(surroundingRunes, []rune(row)[start+1])
			surroundingRunes = append(surroundingRunes, bottomRow[start+1])
		}

		finish := position[1]

		if finish > 0 {
			surroundingRunes = append(surroundingRunes, topRow[finish-1])
			surroundingRunes = append(surroundingRunes, []rune(row)[finish-1])
			surroundingRunes = append(surroundingRunes, bottomRow[finish-1])
		}

		surroundingRunes = append(surroundingRunes, topRow[finish])
		surroundingRunes = append(surroundingRunes, bottomRow[finish])

		if finish < len(row)-1 {
			surroundingRunes = append(surroundingRunes, topRow[finish+1])
			surroundingRunes = append(surroundingRunes, []rune(row)[finish+1])
			surroundingRunes = append(surroundingRunes, bottomRow[finish+1])
		}

		if ContainsSymbol(surroundingRunes) {
			number, _ := strconv.Atoi(row[start : finish+1])
			partNumbers = append(partNumbers, number)
		}
	}

	return partNumbers
}

func ContainsSymbol(runes []rune) bool {
	for _, r := range runes {
		if (unicode.IsSymbol(r) || unicode.IsPunct(r)) && r != '.' {
			return true
		}
	}

	return false
}

var numberRegex = regexp.MustCompile(`\d+`)

func NumberPositions(row string) [][]int {
	indices := numberRegex.FindAllStringIndex(row, -1)

	var positions = make([][]int, 0)
	for _, index := range indices {
		positions = append(positions, []int{index[0], index[1] - 1})
	}

	return positions
}
