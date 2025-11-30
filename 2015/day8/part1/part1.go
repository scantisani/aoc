package part1

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
		sum += CalculateDifference([]rune(line))
	}

	return sum
}

func CalculateDifference(runes []rune) int {
	codeLength := len(runes)

	withoutQuotes := runes[1 : len(runes)-1]
	withoutEscapes := regexp.MustCompile(`\\\\|\\"`).ReplaceAllString(string(withoutQuotes), "a")
	withoutHex := regexp.MustCompile(`\\x[0-9a-f]{2}`).ReplaceAllString(withoutEscapes, "a")

	memoryLength := len(withoutHex)

	return codeLength - memoryLength
}
