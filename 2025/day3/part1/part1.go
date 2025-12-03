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
	highestJoltage := -1

	for i, first := range runes {
		for _, second := range runes[i+1:] {
			joltage, _ := strconv.Atoi(string(first) + string(second))
			highestJoltage = max(highestJoltage, joltage)
		}
	}

	return highestJoltage
}
