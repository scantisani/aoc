package part2

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

	lines := strings.Split(string(content), ",")
	return SumOfInvalidIds(lines)
}

type idRange struct {
	from int
	to   int
}

func SumOfInvalidIds(lines []string) int {
	sum := 0

	for _, line := range lines {
		idRange := parseRange(line)
		invalidIds := InvalidIdsInRange(idRange)

		for _, invalidId := range invalidIds {
			sum += invalidId
		}
	}

	return sum
}

func parseRange(line string) idRange {
	splits := strings.Split(line, "-")
	from, _ := strconv.Atoi(splits[0])
	to, _ := strconv.Atoi(splits[1])

	return idRange{from, to}
}

func InvalidIdsInRange(idRange idRange) []int {
	var invalidIds []int

	for id := idRange.from; id <= idRange.to; id++ {
		if isInvalid(id) {
			invalidIds = append(invalidIds, id)
		}
	}

	return invalidIds
}

var factors = map[int][]int{
	1:  {},
	2:  {2},
	3:  {3},
	4:  {4, 2},
	5:  {5},
	6:  {6, 3, 2},
	7:  {7},
	8:  {8, 4, 2},
	9:  {9, 3},
	10: {10, 5, 2},
}

func isInvalid(id int) bool {
	idAsString := strconv.Itoa(id)

	idLength := len(idAsString)
	possibleSplits := factors[idLength]
	for _, possibleSplit := range possibleSplits {
		parts := SplitIntoParts(idAsString, possibleSplit)
		if allEqual(parts) {
			return true
		}
	}

	return false
}

func SplitIntoParts(id string, numParts int) []string {
	var parts []string

	partLength := len(id) / numParts

	for i := 0; i <= len(id)-partLength; i += partLength {
		part := id[i : i+partLength]
		parts = append(parts, part)
	}

	return parts
}

func allEqual(parts []string) bool {
	compacted := slices.Compact(parts)
	return len(compacted) == 1
}
