package part2

import (
	"cmp"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Solve() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	return CountFreshIngredients(lines)
}

func CountFreshIngredients(lines []string) int {
	separator := slices.Index(lines, "")

	var store FreshStore
	store.Build(lines[:separator])
	return store.count()
}

type FreshRange struct {
	start, end int
}

type FreshStore []FreshRange

func (store *FreshStore) insert(newRange FreshRange) {
	index, _ := slices.BinarySearchFunc(*store, newRange, compareStores)
	*store = slices.Insert(*store, index, newRange)
}

func compareStores(f1 FreshRange, f2 FreshRange) int {
	return cmp.Compare(f1.start, f2.start)
}

func (store *FreshStore) Build(rangeLines []string) {
	for _, line := range rangeLines {
		parsedRange := parse(line)

		newRange, rangesToReplace := store.fitInNewRange(parsedRange)

		if len(rangesToReplace) > 0 {
			*store = slices.Delete(*store, rangesToReplace[0], rangesToReplace[len(rangesToReplace)-1]+1)
			store.insert(newRange)
		} else {
			store.insert(newRange)
		}
	}
}

func (store *FreshStore) fitInNewRange(newRange FreshRange) (FreshRange, []int) {
	var rangesToReplace []int

	for i, existingRange := range *store {
		// does it start and end within the range?
		if newRange.start >= existingRange.start && newRange.end <= existingRange.end {
			// already covered, just use existing range
			return existingRange, []int{i}
		}

		// does it start within the range?
		if newRange.start >= existingRange.start && newRange.start <= existingRange.end {
			// expand to the existing range's start
			newRange = FreshRange{existingRange.start, newRange.end}
			rangesToReplace = append(rangesToReplace, i)
		}

		// does it end within the range?
		if newRange.end >= existingRange.start && newRange.end <= existingRange.end {
			// expand to the existing range's end
			newRange = FreshRange{newRange.start, existingRange.end}
			rangesToReplace = append(rangesToReplace, i)
		}

		// does it completely cover the range?
		if newRange.start <= existingRange.start && newRange.end >= existingRange.end {
			// remove the existing range
			rangesToReplace = append(rangesToReplace, i)
		}
	}

	return newRange, rangesToReplace
}

func (store *FreshStore) count() int {
	sum := 0
	for _, freshRange := range *store {
		sum += freshRange.end - freshRange.start + 1
	}
	return sum
}

func parse(line string) FreshRange {
	splits := strings.Split(line, "-")
	start, _ := strconv.Atoi(splits[0])
	end, _ := strconv.Atoi(splits[1])

	return FreshRange{start, end}
}
