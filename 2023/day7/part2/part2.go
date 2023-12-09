package part2

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
	'A': 13,
	'K': 12,
	'Q': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
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
	nonJokers, numJokers := removeJokers(cards)
	typeWithoutJokers := typeWithoutJokers(nonJokers)

	if numJokers == 0 {
		return typeWithoutJokers
	}

	switch typeWithoutJokers {
	case High:
		switch numJokers {
		case 1:
			return OnePair
		case 2:
			return ThreeOfAKind
		case 3:
			return FourOfAKind
		case 4, 5:
			return FiveOfAKind
		}
	case OnePair:
		switch numJokers {
		case 1:
			return ThreeOfAKind
		case 2:
			return FourOfAKind
		case 3:
			return FiveOfAKind
		}
	case TwoPair:
		return FullHouse
	case ThreeOfAKind:
		switch numJokers {
		case 1:
			return FourOfAKind
		case 2:
			return FiveOfAKind
		}
	case FourOfAKind:
		return FiveOfAKind
	default:
		return typeWithoutJokers
	}

	return typeWithoutJokers
}

func typeWithoutJokers(nonJokers []rune) int {
	sortedCards := append([]rune{}, nonJokers...)
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

func removeJokers(cards []rune) ([]rune, int) {
	numJokers := 0
	nonJokers := make([]rune, 0)

	for _, card := range cards {
		if card == 'J' {
			numJokers++
		} else {
			nonJokers = append(nonJokers, card)
		}
	}

	return nonJokers, numJokers
}

func Duplicates(cards []rune) []int {
	duplicates := make([]int, 0)
	if len(cards) == 0 {
		return duplicates
	}

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
