package part1

import (
	"log"
	"maps"
	"os"
	"regexp"
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
	return ShortestDistance(lines)
}

type Destination string

type Pair struct {
	first  string
	second string
}

type Graph map[string]map[string]int

func ShortestDistance(lines []string) int {
	graph := ParseDistances(lines)

	return shortestInGraph(graph)
}

func ParseDistances(lines []string) Graph {
	graph := Graph{}

	for _, line := range lines {
		first, second, dist := parseDistance(line)

		if len(graph[first]) == 0 {
			graph[first] = map[string]int{}
		}

		if len(graph[second]) == 0 {
			graph[second] = map[string]int{}
		}

		graph[first][second] = dist
		graph[second][first] = dist
	}

	return graph
}

func parseDistance(line string) (string, string, int) {
	regex := regexp.MustCompile("(.*) to (.*) = (.*)")

	from := regex.FindStringSubmatch(line)[1]
	to := regex.FindStringSubmatch(line)[2]
	dist, _ := strconv.Atoi(regex.FindStringSubmatch(line)[3])

	return from, to, dist
}

func shortestInGraph(graph Graph) int {
	destinations := slices.Sorted(maps.Keys(graph))

	minimum := 10_000_000

	permutations := Permutations(destinations)
	for _, permutation := range permutations {
		distance := graph.calculateDistance(permutation)
		minimum = min(minimum, distance)
	}

	return minimum
}

func Permutations(destinations []string) [][]string {
	if len(destinations) == 1 {
		return [][]string{}
	}
	if len(destinations) == 2 {
		return [][]string{{destinations[0], destinations[1]}, {destinations[1], destinations[0]}}
	}

	first := destinations[0]
	rest := destinations[1:]

	permutations := [][]string{}

	permutationsOfRest := Permutations(rest)
	for _, permutation := range permutationsOfRest {
		for i := 0; i <= len(permutation); i++ {
			newPermutation := slices.Clone(permutation)
			inserted := slices.Insert(newPermutation, i, first)
			permutations = append(permutations, inserted)
		}
	}

	return permutations
}

func (graph *Graph) calculateDistance(permutation []string) int {
	distance := 0

	for i := 0; i < len(permutation)-1; i++ {
		distance += (*graph)[permutation[i]][permutation[i+1]]
	}

	return distance
}
