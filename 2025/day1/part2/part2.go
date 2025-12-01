package part2

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

type Movement struct {
	dir    direction
	amount int
}

func TotalZeroes(lines []string) int {
	position := 50
	zeroes := 0

	for _, line := range lines {
		movement := parseMovement(line)
		movementResult := PerformMovement(position, movement)

		position = movementResult.position

		zeroTouchCount := movementResult.zeroTouchCount
		zeroes += zeroTouchCount
	}

	return zeroes
}

func parseMovement(line string) Movement {
	amount, _ := strconv.Atoi(line[1:])

	if line[0] == 'R' {
		return Movement{right, amount}
	} else {
		return Movement{left, amount}
	}
}

type MovementResult struct {
	position       int
	zeroTouchCount int
}

func PerformMovement(position int, movement Movement) MovementResult {
	var newPosition int
	touchedZeroTimes := 0

	if movement.dir == right {
		newPosition = position + movement.amount

		touchedZeroTimes = newPosition / 100
		newPosition = newPosition % 100

		return MovementResult{newPosition, touchedZeroTimes}
	} else {
		newPosition = position - movement.amount

		if newPosition > 0 {
			return MovementResult{newPosition, 0}
		}

		touchedZeroTimes = -(newPosition / 100)
		if position != 0 {
			touchedZeroTimes++
		}

		newPosition = (100 + (newPosition % 100)) % 100

		return MovementResult{newPosition, touchedZeroTimes}
	}
}
