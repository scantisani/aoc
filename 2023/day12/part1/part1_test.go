package part1

import (
	"reflect"
	"strings"
	"testing"
)

func TestSumOfArrangements(t *testing.T) {
	tests := []struct {
		name      string
		gridLines []string
		want      int
	}{
		{
			"AOC example",
			[]string{
				"???.### 1,1,3",
				".??..??...?##. 1,1,3",
				"?#?#?#?#?#?#?#? 1,3,1,6",
				"????.#...#... 4,1,1",
				"????.######..#####. 1,6,5",
				"?###???????? 3,2,1",
			},
			21,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfArrangements(tt.gridLines); got != tt.want {
				t.Errorf("SumOfArrangements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseRow(t *testing.T) {
	tests := []struct {
		name     string
		inputRow string
		want     Row
	}{
		{
			"AOC example 1",
			"???.### 1,1,3",
			Row{[]rune("???.###"), []int{1, 1, 3}},
		},
		{
			"AOC example 2",
			"?#?#?#?#?#?#?#? 1,3,1,6",
			Row{[]rune("?#?#?#?#?#?#?#?"), []int{1, 3, 1, 6}},
		},
		{
			"Input example 1",
			"????????##???#?.? 11,1,1",
			Row{[]rune("????????##???#?.?"), []int{11, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseRow(tt.inputRow); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPossibleArrangements(t *testing.T) {
	tests := []struct {
		name string
		row  Row
		want int
	}{
		{"AOC example row 1", Row{[]rune("???.###"), []int{1, 1, 3}}, 1},
		{"AOC example row 2", Row{[]rune(".??..??...?##."), []int{1, 1, 3}}, 4},
		{"AOC example row 3", Row{[]rune("?#?#?#?#?#?#?#?"), []int{1, 3, 1, 6}}, 1},
		{"AOC example row 4", Row{[]rune("????.#...#..."), []int{4, 1, 1}}, 1},
		{"AOC example row 5", Row{[]rune("????.######..#####."), []int{1, 6, 5}}, 4},
		{"AOC example row 6", Row{[]rune("?###????????"), []int{3, 2, 1}}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PossibleArrangements(tt.row); got != tt.want {
				t.Errorf("PossibleArrangements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGeneratePermutations(t *testing.T) {
	tests := []struct {
		name    string
		springs []rune
		want    [][]rune
	}{
		{
			"Two permutations",
			[]rune("?"),
			[][]rune{
				[]rune("#"),
				[]rune("."),
			},
		},
		{
			"Four permutations",
			[]rune("??"),
			[][]rune{
				[]rune("##"),
				[]rune(".#"),
				[]rune("#."),
				[]rune(".."),
			},
		},
		{
			"Eight permutations",
			[]rune("?#.??"),
			[][]rune{
				[]rune("##.##"),
				[]rune(".#.##"),
				[]rune("##..#"),
				[]rune(".#..#"),
				[]rune("##.#."),
				[]rune(".#.#."),
				[]rune("##..."),
				[]rune(".#..."),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GeneratePermutations(tt.springs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeneratePermutations() =\n%v\nwant\n%v", stringifyPermutations(got), stringifyPermutations(tt.want))
			}
		})
	}
}

func stringifyPermutations(permutations [][]rune) string {
	builder := strings.Builder{}
	for _, permutation := range permutations {
		builder.WriteString(string(permutation))
		builder.WriteRune('\n')
	}
	return builder.String()
}

func TestIsValidRow(t *testing.T) {
	tests := []struct {
		name string
		row  Row
		want bool
	}{
		{
			"Simple true",
			Row{[]rune("#.#.###"), []int{1, 1, 3}},
			true,
		},
		{
			"Simple false",
			Row{[]rune("#######"), []int{1, 1, 3}},
			false,
		},
		{
			"Trickier true",
			Row{[]rune(".#...######..#####."), []int{1, 6, 5}},
			true,
		},
		{
			"Trickier false",
			Row{[]rune(".##..######..#####."), []int{1, 4, 5}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidRow(tt.row); got != tt.want {
				t.Errorf("IsValidRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
