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
	return CountBeamSplits(lines)
}

func CountBeamSplits(lines []string) int {
	count := 0

	beamSourceX := strings.IndexRune(lines[0], 'S')
	beams := set{beamSourceX: true}

	for _, line := range lines[2:] {
		for x, char := range line {
			if char == '^' && beams[x] {
				count++

				delete(beams, x)
				beams[x-1] = true
				beams[x+1] = true
			}
		}
	}

	return count
}

type set map[int]bool
