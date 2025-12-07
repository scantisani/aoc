package part2

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
	return CountTimelines(lines)
}

func CountTimelines(lines []string) int {
	beamSourceX := strings.IndexRune(lines[0], 'S')

	m := multiverse{lines: lines, timelines: map[position]int{}}
	return m.createTimelines(beamSourceX, 1)
}

type position struct {
	x, y int
}

type multiverse struct {
	lines     []string
	timelines map[position]int
}

func (m *multiverse) getTimelinesForPosition(x, y int) int {
	return m.timelines[position{x: x, y: y}]
}

func (m *multiverse) setTimelinesForPosition(x, y, numTimelines int) {
	m.timelines[position{x: x, y: y}] = numTimelines
}

func (m *multiverse) createTimelines(beamX int, lineIndex int) int {
	for y := lineIndex; y < len(m.lines); y++ {
		for x, char := range m.lines[y] {
			if char == '^' && x == beamX {
				precomputedTimelines := m.getTimelinesForPosition(x, y)
				if precomputedTimelines > 0 {
					return precomputedTimelines
				}

				leftTimelines := m.createTimelines(beamX-1, y+1)
				rightTimelines := m.createTimelines(beamX+1, y+1)

				numTimelines := leftTimelines + rightTimelines
				m.setTimelinesForPosition(x, y, numTimelines)
				return numTimelines
			}
		}
	}

	return 1
}
