package part1

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
	return CountFreshIngredients(lines)
}

func CountFreshIngredients(lines []string) int {
	separator := slices.Index(lines, "")

	freshRanges := buildFreshRanges(lines[:separator])
	ingredientIds := buildIds(lines[separator+1:])

	sum := 0
	for _, ingredientId := range ingredientIds {
		if isFresh(freshRanges, ingredientId) {
			sum++
		}
	}
	return sum
}

func buildIds(stringIds []string) []int {
	var ids []int
	for _, stringId := range stringIds {
		id, _ := strconv.Atoi(stringId)
		ids = append(ids, id)
	}
	return ids
}

type freshRange struct {
	from, to int
}

func buildFreshRanges(rangeLines []string) []freshRange {
	var freshRanges []freshRange

	for _, line := range rangeLines {
		splits := strings.Split(line, "-")
		from, _ := strconv.Atoi(splits[0])
		to, _ := strconv.Atoi(splits[1])

		freshRanges = append(freshRanges, freshRange{from, to})
	}

	return freshRanges
}

func isFresh(ranges []freshRange, id int) bool {
	for _, freshRange := range ranges {
		if id >= freshRange.from && id <= freshRange.to {
			return true
		}
	}

	return false
}
