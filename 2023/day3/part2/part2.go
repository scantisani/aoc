package part2

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GearRatioSumFromInput() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	return GearRatioSum(lines[:140]) // don't include the last (blank) split line
}

func GearRatioSum(rows []string) int {
	sum := 0

	for i, row := range rows {
		var topRow string
		var bottomRow string

		if i == 0 {
			topRow = "............................................................................................................................................"
		} else {
			topRow = rows[i-1]
		}

		if i == len(row)-1 {
			bottomRow = "............................................................................................................................................"
		} else {
			bottomRow = rows[i+1]
		}

		gears := Gears(topRow, row, bottomRow)
		for _, gear := range gears {
			ratio := gear.firstPart * gear.secondPart
			sum += ratio
		}
	}

	return sum
}

type gear struct {
	firstPart  int
	secondPart int
}

func Gears(topRow string, row string, bottomRow string) []gear {
	gears := make([]gear, 0)

	gearSymbolPositions := regexp.MustCompile(`\*`).FindAllStringIndex(row, -1)

	for _, symbolPosition := range gearSymbolPositions {
		partNumbers := make([]int, 0)
		symbolIndex := symbolPosition[0]

		topPartNumbers := adjacentParts(topRow, symbolIndex)
		partNumbers = append(partNumbers, topPartNumbers...)

		sameRowPartNumbers := adjacentParts(row, symbolIndex)
		partNumbers = append(partNumbers, sameRowPartNumbers...)

		bottomPartNumbers := adjacentParts(bottomRow, symbolIndex)
		partNumbers = append(partNumbers, bottomPartNumbers...)

		if len(partNumbers) == 2 {
			gears = append(gears, gear{partNumbers[0], partNumbers[1]})
		}
	}

	return gears
}

func adjacentParts(row string, symbolIndex int) []int {
	partNumbers := make([]int, 0)

	positions := NumberPositions(row)
	for _, position := range positions {
		start := position.start
		finish := position.finish

		if symbolIndex >= start-1 && symbolIndex <= finish+1 {
			partNumber, _ := strconv.Atoi(row[start : finish+1])
			partNumbers = append(partNumbers, partNumber)
		}
	}

	return partNumbers
}

var numberRegex = regexp.MustCompile(`\d+`)

type position struct {
	start  int
	finish int
}

func NumberPositions(row string) []position {
	indices := numberRegex.FindAllStringIndex(row, -1)

	var positions = make([]position, 0)
	for _, index := range indices {
		positions = append(positions, position{index[0], index[1] - 1})
	}

	return positions
}
