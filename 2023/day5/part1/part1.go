package part1

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func LowestLocationNumberFromInput() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	return LowestLocationNumber(lines)
}

func LowestLocationNumber(lines []string) int {
	seeds := parseSeeds(lines[0])
	almanac := ParseAlmanac(lines[2:])

	lowestLocation := LocationForSeed(seeds[0], almanac)
	for _, seed := range seeds[1:] {
		seedLocation := LocationForSeed(seed, almanac)
		if seedLocation < lowestLocation {
			lowestLocation = seedLocation
		}
	}

	return lowestLocation
}

func parseSeeds(line string) []int {
	matches := regexp.MustCompile(`\d+`).FindAllString(line, -1)

	seeds := make([]int, 0)
	for _, match := range matches {
		seed, _ := strconv.Atoi(match)
		seeds = append(seeds, seed)
	}

	return seeds
}

func LocationForSeed(seed int, almanac Almanac) int {
	currentNumber := seed
	source := "seed"

	for source != "location" {
		currentMap := almanac[source]
		currentNumber = convertNumber(currentMap, currentNumber)

		source = currentMap.destination
	}

	return currentNumber
}

func convertNumber(almanacMap AlmanacMap, number int) int {
	for _, conversion := range almanacMap.conversions {
		if number >= conversion.sourceStart && number <= conversion.sourceEnd {
			return number + conversion.difference
		}
	}

	return number
}

type AlmanacMap struct {
	source      string
	destination string
	conversions []Conversion
}

type Conversion struct {
	sourceStart int
	sourceEnd   int
	difference  int
}

type Almanac map[string]AlmanacMap

func ParseAlmanac(lines []string) Almanac {
	almanac := Almanac{}

	currentAlmanacLines := make([]string, 0)
	for _, line := range lines {
		if line == "" {
			almanacMap := ParseMap(currentAlmanacLines)
			almanac[almanacMap.source] = almanacMap

			currentAlmanacLines = make([]string, 0)
		} else {
			currentAlmanacLines = append(currentAlmanacLines, line)
		}
	}

	return almanac
}

func ParseMap(lines []string) AlmanacMap {
	titleRegex := regexp.MustCompile(`(\w+)-to-(\w+) map:`)
	matches := titleRegex.FindAllStringSubmatch(lines[0], -1)

	source := matches[0][1]
	destination := matches[0][2]

	conversions := make([]Conversion, 0)
	for _, line := range lines[1:] {
		conversion := ParseConversion(line)
		conversions = prepend(conversions, conversion)
	}

	return AlmanacMap{source, destination, conversions}
}

func ParseConversion(line string) Conversion {
	matches := regexp.MustCompile(`\d+`).FindAllString(line, -1)

	destRangeStart, _ := strconv.Atoi(matches[0])
	sourceRangeStart, _ := strconv.Atoi(matches[1])
	rangeLen, _ := strconv.Atoi(matches[2])

	conversion := Conversion{
		sourceStart: sourceRangeStart,
		sourceEnd:   sourceRangeStart + rangeLen - 1,
		difference:  destRangeStart - sourceRangeStart,
	}

	return conversion
}

func prepend[T any](slice []T, elems ...T) []T {
	return append(elems, slice...)
}
