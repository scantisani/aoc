package part2

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
		sumOfReflections += SmudgedReflectionSize(pattern)
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

func SmudgedReflectionSize(pattern Pattern) int {
	originalSize := ReflectionSize(pattern, -1)

	for i := range pattern {
		for j := range pattern[i] {
			smudgedPattern := smudge(i, j, pattern)
			smudgedSize := ReflectionSize(smudgedPattern, originalSize)

			if smudgedSize != 0 {
				return smudgedSize
			}
		}
	}

	return 0
}

func ReflectionSize(pattern Pattern, forbiddenSize int) int {
	for i := range pattern {
		if IsReflection(pattern[:i], pattern[i:]) {
			possibleNewSize := i * 100
			if possibleNewSize != forbiddenSize {
				return possibleNewSize
			}
		}
	}

	transposedPattern := transpose(pattern)
	for i := range transposedPattern {
		if IsReflection(transposedPattern[:i], transposedPattern[i:]) {
			possibleNewSize := i
			if possibleNewSize != forbiddenSize {
				return possibleNewSize
			}
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

func smudge(i int, j int, pattern Pattern) Pattern {
	smudgedPattern := deepClone(pattern)

	runeToSmudge := smudgedPattern[i][j]
	if runeToSmudge == '.' {
		smudgedPattern[i][j] = '#'
	} else {
		smudgedPattern[i][j] = '.'
	}

	return smudgedPattern
}

func deepClone(pattern Pattern) Pattern {
	newPattern := make(Pattern, 0)

	for _, row := range pattern {
		newPattern = append(newPattern, slices.Clone(row))
	}

	return newPattern
}
