package part2

import (
	"log"
	"os"
	"regexp"
	"strings"
)

func Solve() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	return CharacterDifference(lines)
}

func CharacterDifference(lines []string) int {
	sum := 0

	for _, line := range lines {
		sum += CalculateDifference(line)
	}

	return sum
}

func CalculateDifference(input string) int {
	codeLength := len(input)
	convertedInput := len(ConvertString(input))

	return convertedInput - codeLength
}

func ConvertString(input string) string {
	slashesEscaped := regexp.MustCompile(`\\`).ReplaceAllString(input, `\\`)
	quotesEscaped := regexp.MustCompile(`"`).ReplaceAllString(slashesEscaped, `\"`)

	return `"` + quotesEscaped + `"`
}
