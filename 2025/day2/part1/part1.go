package part1

import (
	"log"
	"os"
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

func isInvalid(id int) bool {
	idAsString := strconv.Itoa(id)

	idLength := len(idAsString)
	if idLength%2 != 0 {
		return false
	}

	first, second := idAsString[0:idLength/2], idAsString[idLength/2:]
	return first == second
}
