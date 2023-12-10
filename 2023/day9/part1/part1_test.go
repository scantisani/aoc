package part1

import (
	"reflect"
	"testing"
)

func TestSumOfPredictions(t *testing.T) {
	tests := []struct {
		name      string
		histories []string
		want      int
	}{
		{
			"AOC examples",
			[]string{
				"0 3 6 9 12 15",
				"1 3 6 10 15 21",
				"10 13 16 21 30 45",
			},
			114,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfPredictions(tt.histories); got != tt.want {
				t.Errorf("SumOfPredictions(%#v) = %v, want %v", tt.histories, got, tt.want)
			}
		})
	}
}

func TestPredictionFor(t *testing.T) {
	tests := []struct {
		name    string
		history string
		want    int
	}{
		{
			"Single extrapolation",
			"0 3 6 9 12 15",
			18,
		},
		{
			"Two extrapolations",
			"1 3 6 10 15 21",
			28,
		},
		{
			"Three extrapolations",
			"10 13 16 21 30 45",
			68,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PredictionFor(tt.history); got != tt.want {
				t.Errorf("PredictionFor(%s) = %v, want %v", tt.history, got, tt.want)
			}
		})
	}
}

func TestExtrapolationFor(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   []int
	}{
		{
			"Single extrapolation",
			[]int{0, 3, 6, 9, 12, 15},
			[]int{3, 3, 3, 3, 3},
		},
		{
			"Two extrapolations",
			[]int{1, 3, 6, 10, 15, 21},
			[]int{2, 3, 4, 5, 6},
		},
		{
			"Three extrapolations",
			[]int{10, 13, 16, 21, 30, 45},
			[]int{3, 3, 5, 9, 15},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtrapolationFor(tt.values)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PredictionFor(%v) = %v, want %v", tt.values, got, tt.want)
			}
		})
	}
}

func TestExtrapolationsFor(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   [][]int
	}{
		{
			"Single extrapolation",
			[]int{0, 3, 6, 9, 12, 15},
			[][]int{{3, 3, 3, 3, 3}, {0, 0, 0, 0}},
		},
		{
			"Two extrapolations",
			[]int{1, 3, 6, 10, 15, 21},
			[][]int{{2, 3, 4, 5, 6}, {1, 1, 1, 1}, {0, 0, 0}},
		},
		{
			"Three extrapolations",
			[]int{10, 13, 16, 21, 30, 45},
			[][]int{{3, 3, 5, 9, 15}, {0, 2, 4, 6}, {2, 2, 2}, {0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtrapolationsFor(tt.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtrapolationsFor() = %v, want %v", got, tt.want)
			}
		})
	}
}
