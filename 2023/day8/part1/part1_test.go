package part1

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
			"AOC Example 1",
			[]string{
				"RL",
				"",
				"AAA = (BBB, CCC)",
				"BBB = (DDD, EEE)",
				"CCC = (ZZZ, GGG)",
				"DDD = (DDD, DDD)",
				"EEE = (EEE, EEE)",
				"GGG = (GGG, GGG)",
				"ZZZ = (ZZZ, ZZZ)",
			},
			2,
		},
		{
			"AOC Example 2",
			[]string{
				"LLR",
				"",
				"AAA = (BBB, BBB)",
				"BBB = (AAA, ZZZ)",
				"ZZZ = (ZZZ, ZZZ)",
			},
			6,
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
