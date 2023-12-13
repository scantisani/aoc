package part1

import (
	"log"
	"os"
	"slices"
	"strings"
)

func SumOfReflectionsFromInput() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	return SumOfReflections(lines[:1349]) // skip last (blank) line
}

func SumOfReflections(inputRows []string) int {
	sumOfReflections := 0

	patterns := parsePatterns(inputRows)
	for _, pattern := range patterns {
		sumOfReflections += ReflectionSize(pattern)
	}

	return sumOfReflections
}

type Pattern [][]rune

func parsePatterns(inputRows []string) []Pattern {
	patterns := make([]Pattern, 0)
	currentPattern := make(Pattern, 0)

	for _, row := range inputRows {
		if row == "" {
			patterns = append(patterns, currentPattern)
			currentPattern = make(Pattern, 0)
		} else {
			currentPattern = append(currentPattern, []rune(row))
		}
	}
	patterns = append(patterns, currentPattern)

	return patterns
}

func ReflectionSize(pattern Pattern) int {
	for i := range pattern {
		if IsReflection(pattern[:i], pattern[i:]) {
			return i * 100
		}
	}

	transposedPattern := transpose(pattern)
	for i := range transposedPattern {
		if IsReflection(transposedPattern[:i], transposedPattern[i:]) {
			return i
		}
	}

	return 0
}

func IsReflection(patternPartA Pattern, patternPartB Pattern) bool {
	maxReflectedLines := slices.Min([]int{len(patternPartA), len(patternPartB)})
	if maxReflectedLines == 0 {
		return false
	}

	for i := 0; i < maxReflectedLines; i++ {
		rowA := patternPartA[len(patternPartA)-i-1]
		rowB := patternPartB[i]

		if !slices.Equal(rowA, rowB) {
			return false
		}
	}

	return true
}

func transpose(pattern Pattern) Pattern {
	rows := make(Pattern, 0)

	for i := range pattern[0] {
		newRow := make([]rune, 0)

		for _, row := range pattern {
			newRow = append(newRow, row[i])
		}

		rows = append(rows, newRow)
	}

	return rows
}
