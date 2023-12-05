package part2

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
	seedRanges := parseSeedRanges(lines[0])
	almanac := ParseAlmanac(lines[2:])

	channel := make(chan int)
	for _, seedRange := range seedRanges {
		go LowestLocationForRange(seedRange, almanac, channel)
	}

	lowestLocation := 9999999999
	for range seedRanges {
		location := <-channel
		if location < lowestLocation {
			lowestLocation = location
		}
	}

	return lowestLocation
}

func LowestLocationForRange(seedRange SeedRange, almanac Almanac, channel chan int) {
	lowestLocation := LocationForSeed(seedRange.start, almanac)

	for seed := seedRange.start; seed <= seedRange.finish; seed++ {
		location := LocationForSeed(seed, almanac)
		if location < lowestLocation {
			lowestLocation = location
		}
	}

	channel <- lowestLocation
}

type SeedRange struct {
	start  int
	finish int
}

func parseSeedRanges(line string) []SeedRange {
	matches := regexp.MustCompile(`\d+`).FindAllString(line, -1)

	seedRanges := make([]SeedRange, 0)
	for i := 0; i < len(matches); i += 2 {
		start, _ := strconv.Atoi(matches[i])
		rangeLength, _ := strconv.Atoi(matches[i+1])
		seedRanges = append(seedRanges, SeedRange{start, start + rangeLength})
	}

	return seedRanges
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
