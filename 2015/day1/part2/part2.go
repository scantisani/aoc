package part2

import (
	"log"
	"os"
)

func Solve() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	return TraverseFloors(string(content))
}

func TraverseFloors(input string) int {
	sum := 0

	inputRunes := []rune(input)
	for i, char := range inputRunes {
		if char == '(' {
			sum += 1
		}
		if char == ')' {
			sum -= 1
		}

		if sum == -1 {
			return i + 1
		}
	}

	return sum
}
