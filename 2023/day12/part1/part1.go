package part1

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func SumOfArrangementsFromInput() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	return SumOfArrangements(lines[:1000]) // skip last (blank) line
}

func SumOfArrangements(inputRows []string) int {
	sumOfArrangements := 0
	for _, inputRow := range inputRows {
		row := ParseRow(inputRow)
		sumOfArrangements += PossibleArrangements(row)
	}
	return sumOfArrangements
}

type Row struct {
	springs []rune
	groups  []int
}

func ParseRow(inputRow string) Row {
	fields := strings.Fields(inputRow)

	springs := []rune(fields[0])
	groups := parseGroups(fields[1])

	return Row{springs, groups}
}

func parseGroups(groupsString string) []int {
	groups := make([]int, 0)

	groupStrings := strings.Split(groupsString, ",")
	for _, groupString := range groupStrings {
		group, _ := strconv.Atoi(groupString)
		groups = append(groups, group)
	}

	return groups
}

func PossibleArrangements(row Row) int {
	possibleArrangements := 0

	permutations := GeneratePermutations(row.springs)
	for _, permutation := range permutations {
		if IsValidRow(Row{permutation, row.groups}) {
			possibleArrangements++
		}
	}

	return possibleArrangements
}

func GeneratePermutations(springs []rune) [][]rune {
	currentRune := springs[0]
	if len(springs) == 1 {
		if currentRune == '?' {
			return [][]rune{{'#'}, {'.'}}
		} else {
			return [][]rune{{currentRune}}
		}
	}

	permutations := make([][]rune, 0)
	tailPermutations := GeneratePermutations(springs[1:])

	if currentRune != '?' {
		for _, tailPermutation := range tailPermutations {
			unchanged := append([]rune{currentRune}, tailPermutation...)
			permutations = append(permutations, unchanged)
		}
	} else {
		for _, tailPermutation := range tailPermutations {
			springPerm := append([]rune{'#'}, tailPermutation...)
			permutations = append(permutations, springPerm)

			emptyPerm := append([]rune{'.'}, tailPermutation...)
			permutations = append(permutations, emptyPerm)
		}
	}

	return permutations
}

func IsValidRow(row Row) bool {
	groups := countGroups(row.springs)
	return slices.Equal(groups, row.groups)
}

func countGroups(springs []rune) []int {
	groups := make([]int, 0)
	currentCount := 0

	for _, spring := range springs {
		if spring == '#' {
			currentCount++
		}

		if spring == '.' && currentCount > 0 {
			groups = append(groups, currentCount)
			currentCount = 0
		}
	}

	if currentCount > 0 {
		groups = append(groups, currentCount)
	}

	return groups
}
