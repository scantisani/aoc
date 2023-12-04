package part2

import (
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func TotalCardsFromInput() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	return TotalCards(lines[:189])
}

func TotalCards(cardStrings []string) int {
	sum := 0

	cardMap := ParseCards(cardStrings)
	for _, card := range cardMap {
		sum += TotalCardsForCard(card, cardMap)
	}

	return sum
}

type CardMap map[int]Card

func TotalCardsForCard(card Card, otherCards CardMap) int {
	totalCards := 1

	numMatches := MatchesForCard(card)

	cardCopies := make([]Card, 0)
	for i := card.id + 1; i <= card.id+numMatches; i++ {
		cardCopies = append(cardCopies, otherCards[i])
	}

	for _, cardCopy := range cardCopies {
		totalCards += TotalCardsForCard(cardCopy, otherCards)
	}

	return totalCards
}

func MatchesForCard(card Card) int {
	matches := 0
	for _, numberHad := range card.numbersYouHave {
		if slices.Contains(card.winningNumbers, numberHad) {
			matches += 1
		}
	}

	return matches
}

type Card struct {
	id             int
	winningNumbers []int
	numbersYouHave []int
}

var cardRegex = regexp.MustCompile(`Card\s+(\d+):(.*)\|(.*)`)

func ParseCards(cardStrings []string) CardMap {
	cardMap := map[int]Card{}

	for _, cardString := range cardStrings {
		card := ParseCard(cardString)
		cardMap[card.id] = card
	}

	return cardMap
}

func ParseCard(cardString string) Card {
	matches := cardRegex.FindAllStringSubmatch(cardString, -1)

	cardId, _ := strconv.Atoi(matches[0][1])
	winningNumbers := digitsFrom(matches[0][2])
	numbersYouHave := digitsFrom(matches[0][3])

	return Card{
		id:             cardId,
		winningNumbers: winningNumbers,
		numbersYouHave: numbersYouHave,
	}
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
