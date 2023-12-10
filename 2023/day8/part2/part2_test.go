package part2

import (
	"reflect"
	"testing"
)

func TestTotalSteps(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  int
	}{
		{
			"AOC example",
			[]string{
				"LR",
				"",
				"11A = (11B, XXX)",
				"11B = (XXX, 11Z)",
				"11Z = (11B, XXX)",
				"22A = (22B, XXX)",
				"22B = (22C, 22C)",
				"22C = (22Z, 22Z)",
				"22Z = (22B, 22B)",
				"XXX = (XXX, XXX)",
			},
			6,
		},
		{
			"Longer example",
			[]string{
				"LR",
				"",
				"11A = (11B, XXX)",
				"11B = (XXX, 11Z)",
				"11Z = (11B, XXX)",
				"22A = (22B, XXX)",
				"22B = (22C, 22C)",
				"22C = (22Z, 22Z)",
				"22Z = (22B, 22B)",
				"33A = (33B, 33B)",
				"33B = (33C, 33C)",
				"33C = (33D, 33D)",
				"33D = (33Z, 33Z)",
				"33Z = (33A, 33A)",
				"XXX = (XXX, XXX)",
			},
			12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TotalSteps(tt.lines); got != tt.want {
				t.Errorf("TotalSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseNodes(t *testing.T) {
	tests := []struct {
		name        string
		nodeStrings []string
		want        Network
	}{
		{
			"AOC Example 1",
			[]string{
				"AAA = (BBB, CCC)",
				"BBB = (DDD, EEE)",
				"CCC = (ZZZ, GGG)",
				"DDD = (DDD, DDD)",
				"EEE = (EEE, EEE)",
				"GGG = (GGG, GGG)",
				"ZZZ = (ZZZ, ZZZ)",
			},
			Network{
				"AAA": Node{"BBB", "CCC"},
				"BBB": Node{"DDD", "EEE"},
				"CCC": Node{"ZZZ", "GGG"},
				"DDD": Node{"DDD", "DDD"},
				"EEE": Node{"EEE", "EEE"},
				"GGG": Node{"GGG", "GGG"},
				"ZZZ": Node{"ZZZ", "ZZZ"},
			},
		},
		{
			"AOC Example 2",
			[]string{
				"AAA = (BBB, BBB)",
				"BBB = (AAA, ZZZ)",
				"ZZZ = (ZZZ, ZZZ)",
			},
			Network{
				"AAA": Node{"BBB", "BBB"},
				"BBB": Node{"AAA", "ZZZ"},
				"ZZZ": Node{"ZZZ", "ZZZ"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseNetwork(tt.nodeStrings); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseNetwork() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGcd(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"Small",
			args{14, 21},
			7,
		},
		{
			"Bigger",
			args{144, 80},
			16,
		},
		{
			"Primes",
			args{13, 23},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gcd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLcm(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    int
	}{
		{
			"Small",
			[]int{6, 8, 12},
			24,
		},
		{
			"Bigger",
			[]int{2, 3, 4, 5, 7},
			420,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lcm(tt.numbers); got != tt.want {
				t.Errorf("Lcm() = %v, want %v", got, tt.want)
			}
		})
	}
}
