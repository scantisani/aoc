package part1

import (
	"log"
	"os"
	"regexp"
	"strings"
)

func TotalStepsFromInput() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	return TotalSteps(lines[:788]) // skip last (blank) line
}

func TotalSteps(lines []string) int {
	steps := regexp.MustCompile(`[LR]`).FindAllString(lines[0], -1)
	network := ParseNetwork(lines[2:])

	stepsTaken := 0
	currentPosition := "AAA"
	for currentPosition != "ZZZ" {
		for _, step := range steps {
			currentNode := network[currentPosition]

			if step == "L" {
				currentPosition = currentNode.left
			} else if step == "R" {
				currentPosition = currentNode.right
			}

			stepsTaken++
			if currentPosition == "ZZZ" {
				break
			}
		}
	}

	return stepsTaken
}

type Node struct {
	left  string
	right string
}

type Network = map[string]Node

func ParseNetwork(nodeStrings []string) Network {
	network := Network{}

	for _, nodeString := range nodeStrings {
		title, node := parseNode(nodeString)
		network[title] = node
	}

	return network
}

var nodeRegex = regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)

func parseNode(nodeString string) (string, Node) {
	matches := nodeRegex.FindAllStringSubmatch(nodeString, -1)

	title := matches[0][1]
	left := matches[0][2]
	right := matches[0][3]

	return title, Node{left, right}
}
