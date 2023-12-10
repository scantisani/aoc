package part2

import (
	"log"
	"math/big"
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

type Node struct {
	left  string
	right string
}

type Network = map[string]Node

func TotalSteps(lines []string) int {
	steps := parseSteps(lines)
	network := ParseNetwork(lines[2:])

	startingPositions := startingPositions(network)

	stepCounts := make([]int, 0)
	for _, position := range startingPositions {
		cycleSteps := StepsToCycle(position, steps, network)
		stepCounts = append(stepCounts, cycleSteps)
	}

	return Lcm(stepCounts)
}

func StepsToCycle(startingPosition string, steps []string, network Network) int {
	currentPosition := startingPosition
	stepCount := 0

	for !endsInZ(currentPosition) || stepCount == 0 {
		for _, step := range steps {
			currentPosition = takeStep(step, currentPosition, network)
			stepCount++

			if endsInZ(currentPosition) {
				return stepCount
			}
		}
	}

	return 0
}

func takeStep(step string, currentPosition string, network Network) string {
	currentNode := network[currentPosition]

	if step == "L" {
		return currentNode.left
	} else {
		return currentNode.right
	}
}

func parseSteps(lines []string) []string {
	return regexp.MustCompile(`[LR]`).FindAllString(lines[0], -1)
}

func startingPositions(network Network) []string {
	positions := make([]string, 0)

	for position := range network {
		if strings.HasSuffix(position, "A") {
			positions = append(positions, position)
		}
	}

	return positions
}

func endsInZ(position string) bool {
	return strings.HasSuffix(position, "Z")
}

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

func Gcd(a, b int64) int64 {
	bigA := big.NewInt(a)
	bigB := big.NewInt(b)

	gcd := big.NewInt(0)
	gcd.GCD(nil, nil, bigA, bigB)

	return gcd.Int64()
}

func Lcm(numbers []int) int {
	lcm := numbers[0]

	for i := 0; i < len(numbers); i++ {
		gcd := Gcd(int64(numbers[i]), int64(lcm))
		lcm = (numbers[i] * lcm) / int(gcd)
	}

	return lcm
}
