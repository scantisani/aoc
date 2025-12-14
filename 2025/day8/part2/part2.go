package part2

import (
	"cmp"
	"log"
	"maps"
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
	return MultiplyLastTwoBoxes(lines)
}

type JunctionBox struct {
	x, y, z int
}

type Connection struct {
	box1, box2 JunctionBox
	distance   float64
}

type Circuit []JunctionBox

func MultiplyLastTwoBoxes(lines []string) int {
	var boxes []JunctionBox

	for _, line := range lines {
		box := parseBox(line)
		boxes = append(boxes, box)
	}

	boxesByDistance := ClosestBoxes(boxes)
	box1, box2 := connectUntilOne(boxesByDistance, len(boxes))

	return box1.x * box2.x
}

func parseBox(line string) JunctionBox {
	splits := strings.Split(line, ",")

	x, _ := strconv.Atoi(splits[0])
	y, _ := strconv.Atoi(splits[1])
	z, _ := strconv.Atoi(splits[2])

	return JunctionBox{x, y, z}
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

func connectUntilOne(sortedConnections []Connection, numToConnect int) (JunctionBox, JunctionBox) {
	maxCircuitNum := 1
	circuits := CircuitMap{}

	for _, connection := range sortedConnections {
		existingBox1Circuit := circuits.existingCircuit(connection.box1)
		existingBox2Circuit := circuits.existingCircuit(connection.box2)

		if existingBox1Circuit == 0 && existingBox2Circuit == 0 {
			circuits[maxCircuitNum] = Circuit{connection.box1, connection.box2}
			maxCircuitNum++
			continue
		}

		if existingBox1Circuit > 0 && existingBox2Circuit > 0 {
			if existingBox1Circuit == existingBox2Circuit {
				continue
			}
			circuits.combineCircuits(existingBox1Circuit, existingBox2Circuit)
		} else if existingBox1Circuit > 0 {
			circuits.addToCircuit(existingBox1Circuit, connection.box2)
		} else if existingBox2Circuit > 0 {
			circuits.addToCircuit(existingBox2Circuit, connection.box1)
		}

		if circuits.allConnected(numToConnect) {
			return connection.box1, connection.box2
		}
	}

	panic("Failed to connect all boxes in one circuit")
}

func (cm *CircuitMap) allConnected(numToConnect int) bool {
	if len(*cm) == 1 {
		for circuit := range maps.Values(*cm) {
			if len(circuit) == numToConnect {
				return true
			}
		}
	}

	return false
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
