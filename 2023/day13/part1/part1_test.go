package part1

import "testing"

func TestSumOfReflections(t *testing.T) {
	tests := []struct {
		name      string
		inputRows []string
		want      int
	}{
		{
			"AOC example",
			[]string{
				"#.##..##.",
				"..#.##.#.",
				"##......#",
				"##......#",
				"..#.##.#.",
				"..##..##.",
				"#.#.##.#.",
				"",
				"#...##..#",
				"#....#..#",
				"..##..###",
				"#####.##.",
				"#####.##.",
				"..##..###",
				"#....#..#",
			},
			405,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfReflections(tt.inputRows); got != tt.want {
				t.Errorf("SumOfReflections() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReflectionSize(t *testing.T) {
	tests := []struct {
		name    string
		pattern Pattern
		want    int
	}{
		{
			"AOC example 1",
			Pattern{
				[]rune("#.##..##."),
				[]rune("..#.##.#."),
				[]rune("##......#"),
				[]rune("##......#"),
				[]rune("..#.##.#."),
				[]rune("..##..##."),
				[]rune("#.#.##.#."),
			},
			5,
		},
		{
			"AOC example 2",
			Pattern{
				[]rune("#...##..#"),
				[]rune("#....#..#"),
				[]rune("..##..###"),
				[]rune("#####.##."),
				[]rune("#####.##."),
				[]rune("..##..###"),
				[]rune("#....#..#"),
			},
			400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReflectionSize(tt.pattern); got != tt.want {
				t.Errorf("ReflectionSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsReflection(t *testing.T) {
	tests := []struct {
		name         string
		patternPartA Pattern
		patternPartB Pattern
		want         bool
	}{
		{
			"Single line",
			Pattern{
				[]rune("#####.##."),
			},
			Pattern{
				[]rune("#####.##."),
			},
			true,
		},
		{
			"Two lines",
			Pattern{
				[]rune("..##..###"),
				[]rune("#####.##."),
			},
			Pattern{
				[]rune("#####.##."),
				[]rune("..##..###"),
			},
			true,
		},
		{
			"A longer than B",
			Pattern{
				[]rune("........."),
				[]rune("#...##..#"),
				[]rune("#....#..#"),
				[]rune("..##..###"),
				[]rune("#####.##."),
			},
			Pattern{
				[]rune("#####.##."),
				[]rune("..##..###"),
				[]rune("#....#..#"),
			},
			true,
		},
		{
			"B longer than A",
			Pattern{
				[]rune("#....#..#"),
				[]rune("..##..###"),
				[]rune("#####.##."),
			},
			Pattern{
				[]rune("#####.##."),
				[]rune("..##..###"),
				[]rune("#....#..#"),
				[]rune("#...##..#"),
				[]rune("........."),
			},
			true,
		},
		{
			"Single line not a reflection",
			Pattern{
				[]rune("#....#..#"),
			},
			Pattern{
				[]rune("........."),
			},
			false,
		},
		{
			"Long pattern not a reflection",
			Pattern{
				[]rune("#####.##."),
				[]rune("..##..###"),
				[]rune("#....#..#"),
				[]rune("#...##..#"),
				[]rune("........."),
			},
			Pattern{
				[]rune("#....#..#"),
				[]rune("..##..###"),
				[]rune("#####.##."),
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsReflection(tt.patternPartA, tt.patternPartB); got != tt.want {
				t.Errorf("IsReflection() = %v, want %v", got, tt.want)
			}
		})
	}
}
