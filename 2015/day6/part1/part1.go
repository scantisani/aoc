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
	return CountLights(lines)
}

var lightMap [1000][1000]bool

func CountLights(instructions []string) int {
	lightMap = [1000][1000]bool{}

	for _, instruction := range instructions {
		displayLights(instruction)
	}

	sum := 0
	for _, row := range lightMap {
		for _, light := range row {
			if light {
				sum++
			}
		}
	}
	return sum
}

type lightOp int

const (
	on lightOp = iota
	off
	toggle
)

type position struct {
	x, y int
}

func displayLights(instruction string) {
	var op lightOp
	var start position
	var end position

	words := strings.Split(instruction, " ")

	if words[1] == "on" {
		op = on
		start = parsePosition(words[2])
		end = parsePosition(words[4])
	} else if words[1] == "off" {
		op = off
		start = parsePosition(words[2])
		end = parsePosition(words[4])
	} else {
		op = toggle
		start = parsePosition(words[1])
		end = parsePosition(words[3])
	}

	for i := start.x; i <= end.x; i++ {
		for j := start.y; j <= end.y; j++ {
			switch op {
			case on:
				lightMap[i][j] = true
			case off:
				lightMap[i][j] = false
			default:
				lightMap[i][j] = !lightMap[i][j]
			}
		}
	}
}

func parsePosition(coordinates string) position {
	splits := strings.Split(coordinates, ",")
	x, _ := strconv.Atoi(splits[0])
	y, _ := strconv.Atoi(splits[1])

	return position{x, y}
}
