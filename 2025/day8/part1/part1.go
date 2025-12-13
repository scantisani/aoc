package part1

import (
	"cmp"
	"log"
	"math"
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
	return MultiplyCircuitLengths(lines, 1000)
}

type JunctionBox struct {
	x, y, z int
}

type Connection struct {
	box1, box2 JunctionBox
	distance   float64
}

type Circuit []JunctionBox

func MultiplyCircuitLengths(lines []string, numCircuits int) int {
	multiplied := 1
	var boxes []JunctionBox

	for _, line := range lines {
		box := parseBox(line)
		boxes = append(boxes, box)
	}

	boxesByDistance := ClosestBoxes(boxes)
	circuits := connectBoxes(boxesByDistance, numCircuits)
	largest := circuits.threeLargest()

	for _, length := range largest {
		multiplied *= length
	}
	return multiplied
}

func parseBox(line string) JunctionBox {
	splits := strings.Split(line, ",")

	x, _ := strconv.Atoi(splits[0])
	y, _ := strconv.Atoi(splits[1])
	z, _ := strconv.Atoi(splits[2])

	return JunctionBox{
		x, y, z,
	}
}

func ClosestBoxes(boxes []JunctionBox) []Connection {
	var connections []Connection

	for i, box1 := range boxes {
		for _, box2 := range boxes[i+1:] {
			distance := EuclideanDistance(box1, box2)
			connection := Connection{box1, box2, distance}
			connections = append(connections, connection)
		}
	}

	slices.SortFunc(connections, func(connection1, connection2 Connection) int {
		return cmp.Compare(connection1.distance, connection2.distance)
	})

	return connections
}

func EuclideanDistance(box1 JunctionBox, box2 JunctionBox) float64 {
	return math.Sqrt(
		math.Pow(float64(box1.x)-float64(box2.x), 2) +
			math.Pow(float64(box1.y)-float64(box2.y), 2) +
			math.Pow(float64(box1.z)-float64(box2.z), 2),
	)
}

type CircuitMap map[int]Circuit

func connectBoxes(sortedConnections []Connection, numCircuits int) CircuitMap {
	maxCircuitNum := 1
	circuits := CircuitMap{}

	for i := range numCircuits {
		connection := sortedConnections[i]

		existingBox1Circuit := circuits.existingCircuit(connection.box1)
		existingBox2Circuit := circuits.existingCircuit(connection.box2)

		if existingBox1Circuit > 0 && existingBox2Circuit > 0 {
			if existingBox1Circuit == existingBox2Circuit {
				continue
			}
			circuits.combineCircuits(existingBox1Circuit, existingBox2Circuit)
		} else if existingBox1Circuit > 0 {
			circuits.addToCircuit(existingBox1Circuit, connection.box2)
		} else if existingBox2Circuit > 0 {
			circuits.addToCircuit(existingBox2Circuit, connection.box1)
		} else {
			circuits[maxCircuitNum] = Circuit{connection.box1, connection.box2}
			maxCircuitNum++
		}
	}

	return circuits
}

func (cm *CircuitMap) threeLargest() []int {
	var lengths []int
	for _, circuit := range *cm {
		lengths = append(lengths, len(circuit))
	}
	slices.SortFunc(lengths, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	return []int{lengths[0], lengths[1], lengths[2]}
}

func (cm *CircuitMap) existingCircuit(box JunctionBox) int {
	for i, circuit := range *cm {
		if slices.Contains(circuit, box) {
			return i
		}
	}
	return 0
}

func (cm *CircuitMap) combineCircuits(circuitAIndex int, circuitBIndex int) {
	circuitA := (*cm)[circuitAIndex]
	circuitB := (*cm)[circuitBIndex]

	(*cm)[circuitAIndex] = slices.Concat(circuitA, circuitB)
	delete(*cm, circuitBIndex)
}

func (cm *CircuitMap) addToCircuit(circuitIndex int, box JunctionBox) {
	(*cm)[circuitIndex] = append((*cm)[circuitIndex], box)
}
