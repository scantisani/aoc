package part1

import (
	"reflect"
	"testing"
)

func TestPointTotalForCards(t *testing.T) {
	cards :=
		[]string{
			"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
		}

	t.Run("AOC input", func(t *testing.T) {
		if got := PointTotalForCards(cards); got != 13 {
			t.Errorf("PointTotalForCards() = %v, want %v", got, 13)
		}
	})
}

func TestPointTotalForCard(t *testing.T) {
	type args struct {
		card string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Card 1", args{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"}, 8},
		{"Card 2", args{"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19"}, 2},
		{"Card 3", args{"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1"}, 2},
		{"Card 4", args{"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83"}, 1},
		{"Card 5", args{"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36"}, 0},
		{"Card 6", args{"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PointTotalForCard(tt.args.card); got != tt.want {
				t.Errorf("PointTotalForCard(%v) = %v, want %v", tt.args.card, got, tt.want)
			}
		})
	}
}

func TestNumbersOnCard(t *testing.T) {
	tests := []struct {
		name           string
		card           string
		winningNumbers []int
		had            []int
	}{
		{"Card 1", "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", []int{41, 48, 83, 86, 17}, []int{83, 86, 6, 31, 17, 9, 48, 53}},
		{"Card 2", "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", []int{13, 32, 20, 16, 61}, []int{61, 30, 68, 82, 17, 32, 24, 19}},
		{"Card 3", "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", []int{1, 21, 53, 59, 44}, []int{69, 82, 63, 72, 16, 21, 14, 1}},
		{"Card 4", "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", []int{41, 92, 73, 84, 69}, []int{59, 84, 76, 51, 58, 5, 54, 83}},
		{"Card 5", "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", []int{87, 83, 26, 28, 32}, []int{88, 30, 70, 12, 93, 22, 82, 36}},
		{"Card 6", "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", []int{31, 18, 13, 56, 72}, []int{74, 77, 10, 23, 35, 67, 36, 11}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			winningNumbers, had := NumbersOnCard(tt.card)
			if !reflect.DeepEqual(winningNumbers, tt.winningNumbers) {
				t.Errorf("NumbersOnCard(%v).winningNumbers = %v, want %v", tt.card, winningNumbers, tt.winningNumbers)
			}

			if !reflect.DeepEqual(had, tt.had) {
				t.Errorf("NumbersOnCard(%v).had = %v, want %v", tt.card, had, tt.had)
			}
		})
	}
}
