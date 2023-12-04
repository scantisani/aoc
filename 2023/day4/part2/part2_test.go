package part2

import (
	"reflect"
	"testing"
)

func TestTotalCards(t *testing.T) {
	cards :=
		[]string{
			"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
		}

	want := 30

	t.Run("AOC input", func(t *testing.T) {
		if got := TotalCards(cards); got != want {
			t.Errorf("TotalCards() = %v, want %v", got, want)
		}
	})
}

func TestTotalCardsForCard(t *testing.T) {
	cardMap := CardMap{
		1: {1, []int{41, 48, 83, 86, 17}, []int{83, 86, 6, 31, 17, 9, 48, 53}},
		2: {2, []int{13, 32, 20, 16, 61}, []int{61, 30, 68, 82, 17, 32, 24, 19}},
		3: {3, []int{1, 21, 53, 59, 44}, []int{69, 82, 63, 72, 16, 21, 14, 1}},
		4: {4, []int{41, 92, 73, 84, 69}, []int{59, 84, 76, 51, 58, 5, 54, 83}},
		5: {5, []int{87, 83, 26, 28, 32}, []int{88, 30, 70, 12, 93, 22, 82, 36}},
		6: {6, []int{31, 18, 13, 56, 72}, []int{74, 77, 10, 23, 35, 67, 36, 11}},
	}

	tests := []struct {
		name string
		card Card
		want int
	}{
		{"Card 1", cardMap[1], 15},
		{"Card 2", cardMap[2], 7},
		{"Card 3", cardMap[3], 4},
		{"Card 4", cardMap[4], 2},
		{"Card 5", cardMap[5], 1},
		{"Card 6", cardMap[6], 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TotalCardsForCard(tt.card, cardMap); got != tt.want {
				t.Errorf("TotalCardsForCard(%v) = %v, want %v", tt.card, got, tt.want)
			}
		})
	}
}

func TestMatchesForCard(t *testing.T) {
	tests := []struct {
		name            string
		card            Card
		expectedMatches int
	}{
		{"Card 1", Card{1, []int{41, 48, 83, 86, 17}, []int{83, 86, 6, 31, 17, 9, 48, 53}}, 4},
		{"Card 2", Card{2, []int{13, 32, 20, 16, 61}, []int{61, 30, 68, 82, 17, 32, 24, 19}}, 2},
		{"Card 3", Card{3, []int{1, 21, 53, 59, 44}, []int{69, 82, 63, 72, 16, 21, 14, 1}}, 2},
		{"Card 4", Card{4, []int{41, 92, 73, 84, 69}, []int{59, 84, 76, 51, 58, 5, 54, 83}}, 1},
		{"Card 5", Card{5, []int{87, 83, 26, 28, 32}, []int{88, 30, 70, 12, 93, 22, 82, 36}}, 0},
		{"Card 6", Card{6, []int{31, 18, 13, 56, 72}, []int{74, 77, 10, 23, 35, 67, 36, 11}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := MatchesForCard(tt.card)
			if !reflect.DeepEqual(card, tt.expectedMatches) {
				t.Errorf("MatchesForCard(%v) = %v, want %v", tt.card, card, tt.expectedMatches)
			}
		})
	}
}

func TestParseCards(t *testing.T) {
	cardLines := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}

	cardMap := CardMap{
		1: {1, []int{41, 48, 83, 86, 17}, []int{83, 86, 6, 31, 17, 9, 48, 53}},
		2: {2, []int{13, 32, 20, 16, 61}, []int{61, 30, 68, 82, 17, 32, 24, 19}},
		3: {3, []int{1, 21, 53, 59, 44}, []int{69, 82, 63, 72, 16, 21, 14, 1}},
		4: {4, []int{41, 92, 73, 84, 69}, []int{59, 84, 76, 51, 58, 5, 54, 83}},
		5: {5, []int{87, 83, 26, 28, 32}, []int{88, 30, 70, 12, 93, 22, 82, 36}},
		6: {6, []int{31, 18, 13, 56, 72}, []int{74, 77, 10, 23, 35, 67, 36, 11}},
	}

	card := ParseCards(cardLines)
	if !reflect.DeepEqual(card, cardMap) {
		t.Errorf("ParseCards(%v).winningNumbers = %v, want %v", cardLines, card, cardMap)
	}
}

func TestParseCard(t *testing.T) {
	tests := []struct {
		name         string
		cardString   string
		expectedCard Card
	}{
		{"Card 1", "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", Card{1, []int{41, 48, 83, 86, 17}, []int{83, 86, 6, 31, 17, 9, 48, 53}}},
		{"Card 2", "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", Card{2, []int{13, 32, 20, 16, 61}, []int{61, 30, 68, 82, 17, 32, 24, 19}}},
		{"Card 3", "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", Card{3, []int{1, 21, 53, 59, 44}, []int{69, 82, 63, 72, 16, 21, 14, 1}}},
		{"Card 4", "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", Card{4, []int{41, 92, 73, 84, 69}, []int{59, 84, 76, 51, 58, 5, 54, 83}}},
		{"Card 5", "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", Card{5, []int{87, 83, 26, 28, 32}, []int{88, 30, 70, 12, 93, 22, 82, 36}}},
		{"Card 6", "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", Card{6, []int{31, 18, 13, 56, 72}, []int{74, 77, 10, 23, 35, 67, 36, 11}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := ParseCard(tt.cardString)
			if !reflect.DeepEqual(card, tt.expectedCard) {
				t.Errorf("ParseCard(%v).winningNumbers = %v, want %v", tt.cardString, card, tt.expectedCard)
			}
		})
	}
}
