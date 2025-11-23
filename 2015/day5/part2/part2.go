package part2

import (
	"log"
	"os"
	"strings"
)

func Solve() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	return CountNiceStrings(lines)
}

func CountNiceStrings(candidates []string) int {
	sum := 0

	for _, candidate := range candidates {
		if IsNiceString(candidate) {
			sum++
		}
	}

	return sum
}

func IsNiceString(candidate string) bool {
	if len(candidate) < 4 {
		return false
	}

	hasAba := false

	pairs := map[string]int{}
	hasTwoPairs := false

	heldPair := candidate[0:2]

	for i := 2; i < len(candidate); i++ {
		prevprev := candidate[i-2]
		current := candidate[i]

		if prevprev == current {
			hasAba = true
		}

		pairs[heldPair]++
		if pairs[heldPair] > 1 {
			hasTwoPairs = true
		}

		pair := candidate[i-1 : i+1]
		if pair == heldPair {
			heldPair = ""
		} else {
			heldPair = pair
		}
	}
	pairs[heldPair]++
	if pairs[heldPair] > 1 {
		hasTwoPairs = true
	}

	return hasAba && hasTwoPairs
}
