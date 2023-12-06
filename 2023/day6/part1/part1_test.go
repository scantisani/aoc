package part1

import (
	"reflect"
	"testing"
)

func TestWaysMultiplied(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  int
	}{
		{
			"AOC example input",
			[]string{
				"Time:      7  15   30",
				"Distance:  9  40  200",
			},
			288,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WaysMultiplied(tt.lines); got != tt.want {
				t.Errorf("WaysMultiplied() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWaysToBeatRecord(t *testing.T) {
	tests := []struct {
		name string
		race Race
		want int
	}{
		{
			"Short example race",
			Race{7, 9},
			4,
		},
		{
			"Medium example race",
			Race{15, 40},
			8,
		},
		{
			"Long example race",
			Race{30, 200},
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WaysToBeatRecord(tt.race); got != tt.want {
				t.Errorf("WaysToBeatRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseRaces(t *testing.T) {
	type args struct {
		timeLine     string
		distanceLine string
	}
	tests := []struct {
		name string
		args args
		want []Race
	}{
		{
			"AOC Example Input",
			args{
				"Time:      7  15   30",
				"Distance:  9  40  200",
			},
			[]Race{
				{7, 9},
				{15, 40},
				{30, 200},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseRaces(tt.args.timeLine, tt.args.distanceLine); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseRaces(%#v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
