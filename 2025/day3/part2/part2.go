package part2

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

	lines := strings.Split(string(content), "\n")
	return SumOfJoltages(lines)
}

func SumOfJoltages(lines []string) int {
	sum := 0

	for _, line := range lines {
		sum += HighestJoltage(line)
	}

	return sum
}

func HighestJoltage(line string) int {
	runes := []rune(line)

	var maxes []rune
	start := 0

	for i := 11; i >= 0; i-- {
		maximum, newIndex := maxValue(runes[start : len(runes)-i])
		maxes = append(maxes, maximum)
		start += newIndex + 1
	}

	highestJoltage, _ := strconv.Atoi(string(maxes))
	return highestJoltage
}

func maxValue(runes []rune) (rune, int) {
	maximum := -1
	index := -1

	for i, char := range runes {
		asInt := int(char) - '0'
		if asInt > maximum {
			maximum = asInt
			index = i
		}
	}

	return rune(maximum + '0'), index
}
