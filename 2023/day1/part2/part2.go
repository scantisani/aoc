package part2

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func CalibrationFromInput() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	return CalibrationSum(lines)
}

func CalibrationSum(inputs []string) int {
	sum := 0

	for _, input := range inputs {
		sum += Calibration(input)
	}

	return sum
}

type digit struct {
	asString string
	asInt    int
}

func digitStrings() []digit {
	return []digit{
		{"one", 1},
		{"two", 2},
		{"three", 3},
		{"four", 4},
		{"five", 5},
		{"six", 6},
		{"seven", 7},
		{"eight", 8},
		{"nine", 9},
	}
}

func Calibration(input string) int {
	firstNumber := firstNumber(input)
	lastNumber := lastNumber(input)

	final, _ := strconv.Atoi(fmt.Sprintf(`%d%d`, firstNumber, lastNumber))
	return final
}

func firstNumber(input string) int {
	for i := 0; i < len(input); i++ {
		char := input[i : i+1]

		converted, err := strconv.Atoi(char)
		if err == nil {
			return converted
		}

		remainingString := input[i:]
		for _, digitString := range digitStrings() {
			if strings.HasPrefix(remainingString, digitString.asString) {
				return digitString.asInt
			}
		}
	}

	return 0
}

func lastNumber(input string) int {
	for i := len(input) - 1; i >= 0; i-- {
		char := input[i : i+1]

		converted, err := strconv.Atoi(char)
		if err == nil {
			return converted
		}

		remainingString := input[:i+1]
		for _, digitString := range digitStrings() {
			if strings.HasSuffix(remainingString, digitString.asString) {
				return digitString.asInt
			}
		}
	}

	return 0
}
