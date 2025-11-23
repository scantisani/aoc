package part1

import (
	"log"
	"os"
	"strings"
)

func Solve() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	return CountNiceStrings(lines)
}

func CountNiceStrings(candidates []string) int {
	sum := 0

	for _, candidate := range candidates {
		if IsNiceString(candidate) {
			sum++
		}
	}

	return sum
}

var vowels = [...]string{"a", "e", "i", "o", "u"}
var doubleLetters = [...]string{"aa", "bb", "cc", "dd", "ee", "ff", "gg", "hh", "ii", "jj", "kk", "ll", "mm", "nn", "oo", "pp", "qq", "rr", "ss", "tt", "uu", "vv", "ww", "xx", "yy", "zz"}
var naughtySubStrings = [...]string{"ab", "cd", "pq", "xy"}

func IsNiceString(candidate string) bool {
	if containsNaughtySubstring(candidate) {
		return false
	}

	if !hasMinThreeVowels(candidate) {
		return false
	}

	if !hasDoubleLetter(candidate) {
		return false
	}

	return true
}

func containsNaughtySubstring(candidate string) bool {
	for _, naughtySubString := range naughtySubStrings {
		if strings.Contains(candidate, naughtySubString) {
			return true
		}
	}
	return false
}

func hasMinThreeVowels(candidate string) bool {
	numVowels := CountVowels(candidate)
	return numVowels >= 3
}

func CountVowels(candidate string) int {
	sum := 0

	for _, vowel := range vowels {
		sum += strings.Count(candidate, vowel)
	}

	return sum
}

func hasDoubleLetter(candidate string) bool {
	for _, doubles := range doubleLetters {
		if strings.Contains(candidate, doubles) {
			return true
		}
	}

	return false
}
