package part1

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
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

func Calibration(input string) int {
	inputRunes := []rune(input)

	firstDigit := 0
	lastDigit := 0

	for i := 0; i < len(inputRunes); i++ {
		char := inputRunes[i]
		if unicode.IsDigit(char) {
			firstDigit = int(char) - '0'
			break
		}
	}

	for i := len(inputRunes) - 1; i >= 0; i-- {
		char := inputRunes[i]
		if unicode.IsDigit(char) {
			lastDigit = int(char) - '0'
			break
		}
	}

	final, _ := strconv.Atoi(fmt.Sprintf(`%d%d`, firstDigit, lastDigit))
	return final
}
