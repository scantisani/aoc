package part1

import (
	"cmp"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func TotalWinningsFromInput() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	return TotalWinnings(lines[:1000]) // skip last (blank) line
}

type Hand struct {
	cards []rune
	bid   int
}

func TotalWinnings(handStrings []string) int {
	totalWinnings := 0

	hands := parseHands(handStrings)
	slices.SortFunc(hands, func(a, b Hand) int {
		return CompareHands(a, b)
	})

	for i, hand := range hands {
		totalWinnings += hand.bid * (i + 1)
	}

	return totalWinnings
}

func parseHands(handStrings []string) []Hand {
	hands := make([]Hand, 0)
	for _, handString := range handStrings {
		hands = append(hands, parseHand(handString))
	}
	return hands
}

func parseHand(handString string) Hand {
	matches := regexp.MustCompile(`(\w{5}) (\d+)`).FindStringSubmatch(handString)

	cards := parseCards(matches[1])
	bid, _ := strconv.Atoi(matches[2])

	return Hand{cards, bid}
}

func CompareHands(a, b Hand) int {
	typeA := TypeOfHand(a.cards)
	typeB := TypeOfHand(b.cards)

	if typeA > typeB {
		return 1
	}
	if typeA < typeB {
		return -1
	}

	for i := range a.cards {
		if cardRanks[a.cards[i]] > cardRanks[b.cards[i]] {
			return 1
		}

		if cardRanks[a.cards[i]] < cardRanks[b.cards[i]] {
			return -1
		}
	}

	return 0
}

var cardRanks = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

// cards ranks
const (
	High         = iota
	OnePair      = iota
	TwoPair      = iota
	ThreeOfAKind = iota
	FullHouse    = iota
	FourOfAKind  = iota
	FiveOfAKind  = iota
)

func TypeOfHand(cards []rune) int {
	sortedCards := append([]rune{}, cards...)
	slices.SortFunc(sortedCards, func(a, b rune) int {
		return cmp.Compare(cardRanks[b], cardRanks[a])
	})

	duplicates := Duplicates(sortedCards)

	if len(duplicates) == 1 {
		switch duplicates[0] {
		case 2:
			return OnePair
		case 3:
			return ThreeOfAKind
		case 4:
			return FourOfAKind
		case 5:
			return FiveOfAKind
		}
	}

	if len(duplicates) == 2 {
		slices.Sort(duplicates)

		if duplicates[1] == 2 {
			return TwoPair
		}
		if duplicates[1] == 3 {
			return FullHouse
		}
	}

	return High
}

func Duplicates(cards []rune) []int {
	duplicates := make([]int, 0)

	currentDuplicates := 1
	previousCard := cards[0]
	for _, card := range cards[1:] {
		if card == previousCard {
			currentDuplicates++
		} else {
			if currentDuplicates > 1 {
				duplicates = append(duplicates, currentDuplicates)
				currentDuplicates = 1
			}
		}

		previousCard = card
	}

	if currentDuplicates > 1 {
		duplicates = append(duplicates, currentDuplicates)
	}

	return duplicates
}

func parseCards(cardString string) []rune {
	matches := regexp.MustCompile(`\w`).FindAllString(cardString, -1)
	return stringsToRunes(matches)
}

func stringsToRunes(strings []string) []rune {
	runes := make([]rune, 0)

	for _, s := range strings {
		runes = append(runes, rune(s[0]))
	}

	return runes
}
