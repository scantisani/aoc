package part1

import (
	"reflect"
	"testing"
)

func TestTotalWinnings(t *testing.T) {
	tests := []struct {
		name  string
		hands []string
		want  int
	}{
		{
			"AOC Example input",
			[]string{
				"32T3K 765",
				"T55J5 684",
				"KK677 28",
				"KTJJT 220",
				"QQQJA 483",
			},
			6440,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TotalWinnings(tt.hands); got != tt.want {
				t.Errorf("TotalWinnings(%#v) = %v, want %v", tt.hands, got, tt.want)
			}
		})
	}
}

func TestTypeOfHand(t *testing.T) {
	tests := []struct {
		name  string
		cards []rune
		want  int
	}{
		{
			"High",
			[]rune("A2345"),
			High,
		},
		{
			"One pair",
			[]rune("32T3K"),
			OnePair,
		},
		{
			"Two pair",
			[]rune("KTJJT"),
			TwoPair,
		},
		{
			"Two pair with high first card",
			[]rune("KK677"),
			TwoPair,
		},
		{
			"Three of a kind",
			[]rune("T55J5"),
			ThreeOfAKind,
		},
		{
			"Three of a kind with high first card",
			[]rune("QQQJA"),
			ThreeOfAKind,
		},
		{
			"Four of a kind",
			[]rune("4444A"),
			FourOfAKind,
		},
		{
			"Full house",
			[]rune("444JJ"),
			FullHouse,
		},
		{
			"Five of a kind",
			[]rune("66666"),
			FiveOfAKind,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TypeOfHand(tt.cards); got != tt.want {
				t.Errorf("TypeOfHand(%v) = %v, want %v", tt.cards, got, tt.want)
			}
		})
	}
}

func TestDuplicates(t *testing.T) {
	tests := []struct {
		name  string
		cards []rune
		want  []int
	}{
		{
			"High",
			[]rune("A5432"),
			[]int{},
		},
		{
			"One pair",
			[]rune("KT332"),
			[]int{2},
		},
		{
			"Two pair",
			[]rune("KJJTT"),
			[]int{2, 2},
		},
		{
			"Two pair with high first card",
			[]rune("KK776"),
			[]int{2, 2},
		},
		{
			"Three of a kind",
			[]rune("JT555"),
			[]int{3},
		},
		{
			"Three of a kind with high first card",
			[]rune("QQQJA"),
			[]int{3},
		},
		{
			"Four of a kind",
			[]rune("4444A"),
			[]int{4},
		},
		{
			"Full house",
			[]rune("JJ444"),
			[]int{2, 3},
		},
		{
			"Five of a kind",
			[]rune("66666"),
			[]int{5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Duplicates(tt.cards); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Duplicates(%v) = %v, want %v", tt.cards, got, tt.want)
			}
		})
	}
}

func TestCompareHands(t *testing.T) {
	type args struct {
		a Hand
		b Hand
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Five of a kind versus full house",
			args{
				Hand{[]rune("55555"), 0},
				Hand{[]rune("35535"), 0},
			},
			1,
		},
		{
			"Full house versus five of a kind",
			args{
				Hand{[]rune("35535"), 0},
				Hand{[]rune("55555"), 0},
			},
			-1,
		},
		{
			"Full house versus three of a kind",
			args{
				Hand{[]rune("35535"), 0},
				Hand{[]rune("AATAJ"), 0},
			},
			1,
		},
		{
			"Four of a kind versus four of a kind",
			args{
				Hand{[]rune("33332"), 0},
				Hand{[]rune("2AAAA"), 0},
			},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareHands(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("CompareHands(%v, %v) = %v, want %v", tt.args.a, tt.args.b, got, tt.want)
			}
		})
	}
}
