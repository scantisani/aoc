package part1

import (
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func PointTotalForInput() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	return PointTotalForCards(lines[:189])
}

func PointTotalForCards(cards []string) int {
	sum := 0

	for _, card := range cards {
		sum += PointTotalForCard(card)
	}

	return sum
}

func PointTotalForCard(card string) int {
	winningNumbers, had := NumbersOnCard(card)

	var winningNumbersHad []int
	for _, numberHad := range had {
		if slices.Contains(winningNumbers, numberHad) {
			winningNumbersHad = append(winningNumbersHad, numberHad)
		}
	}

	return int(math.Pow(2, float64(len(winningNumbersHad)-1)))
}

var cardRegex = regexp.MustCompile(`Card\s+\d+:(.*)\|(.*)`)

func NumbersOnCard(card string) (winningNumbers []int, numbersYouHave []int) {
	matches := cardRegex.FindAllStringSubmatch(card, -1)

	winningNumberString := matches[0][1]
	numbersYouHaveString := matches[0][2]

	return digitsFrom(winningNumberString), digitsFrom(numbersYouHaveString)
}

var numberRegex = regexp.MustCompile(`\d+`)

func digitsFrom(winningNumberString string) []int {
	digits := make([]int, 0)

	digitMatches := numberRegex.FindAllString(winningNumberString, -1)
	for _, digitMatch := range digitMatches {
		digit, _ := strconv.Atoi(digitMatch)
		digits = append(digits, digit)
	}

	return digits
}
