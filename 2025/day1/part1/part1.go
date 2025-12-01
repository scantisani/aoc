package part1

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	return TotalZeroes(lines)
}

type direction int

const (
	left  direction = iota
	right direction = iota
)

type movement struct {
	dir    direction
	amount int
}

func TotalZeroes(lines []string) int {
	current := 50
	zeroes := 0

	for _, line := range lines {
		movement := parseMovement(line)
		current = performMovement(current, movement)

		if current == 0 {
			zeroes++
		}
	}

	return zeroes
}

func parseMovement(line string) movement {
	amount, _ := strconv.Atoi(line[1:])

	if line[0] == 'R' {
		return movement{right, amount}
	} else {
		return movement{left, amount}
	}
}

func performMovement(current int, movement movement) int {
	var newValue int

	if movement.dir == right {
		newValue = current + movement.amount
	} else {
		newValue = current - movement.amount
	}

	if newValue < 0 {
		newValue = 100 + (newValue % 100)
	}
	if newValue > 99 {
		return newValue % 100
	}

	return newValue
}
