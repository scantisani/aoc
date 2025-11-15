package part2

import (
	"log"
	"os"
)

func Solve() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	return CountHouses(string(content))
}

type Position struct {
	x int
	y int
}

func CountHouses(path string) int {
	houseVisits := VisitHouses(path)

	return len(houseVisits)
}

func VisitHouses(path string) map[Position]int {
	santaPosition := Position{0, 0}
	roboSantaPosition := Position{0, 0}

	houseMap := map[Position]int{
		Position{0, 0}: 1,
	}

	for i, move := range []rune(path) {
		var currentPosition *Position

		if i%2 == 0 {
			currentPosition = &santaPosition
		} else {
			currentPosition = &roboSantaPosition
		}

		switch move {
		case '>':
			currentPosition.x++
		case '<':
			currentPosition.x--
		case '^':
			currentPosition.y++
		case 'v':
			currentPosition.y--
		}

		if houseMap[*currentPosition] == 0 {
			houseMap[*currentPosition] = 1
		} else {
			houseMap[*currentPosition]++
		}
	}

	return houseMap
}
