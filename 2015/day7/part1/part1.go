package part1

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve() uint16 {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	circuit := AssembleCircuit(lines)
	return circuit.FindSignal("a")
}

type operation int

const (
	bare operation = iota
	and
	or
	lshift
	rshift
	not
)

type Circuit map[string]Wire
type Wire struct {
	op       operation
	deps     []string
	signal   uint16
	computed bool
}

func AssembleCircuit(lines []string) Circuit {
	wires := Circuit{}

	for _, line := range lines {
		name, wire := buildWire(line)
		wires[name] = wire
	}

	return wires
}

func buildWire(line string) (string, Wire) {
	tokens := strings.Split(line, " ")

	switch tokens[1] {
	case "->":
		return tokens[2], Wire{
			op:   bare,
			deps: []string{tokens[0]},
		}
	case "AND":
		return tokens[4], Wire{
			op:   and,
			deps: []string{tokens[0], tokens[2]},
		}
	case "OR":
		return tokens[4], Wire{
			op:   or,
			deps: []string{tokens[0], tokens[2]},
		}
	case "LSHIFT":
		return tokens[4], Wire{
			op:   lshift,
			deps: []string{tokens[0], tokens[2]},
		}
	case "RSHIFT":
		return tokens[4], Wire{
			op:   rshift,
			deps: []string{tokens[0], tokens[2]},
		}
	default:
		return tokens[3], Wire{
			op:   not,
			deps: []string{tokens[1]},
		}
	}
}

func (circuit *Circuit) FindSignal(wireName string) uint16 {
	parsed, err := strconv.ParseUint(wireName, 10, 16)
	if err == nil {
		return uint16(parsed)
	}

	wire := (*circuit)[wireName]
	if wire.computed {
		return wire.signal
	}

	signal := circuit.computeWire(wire)

	wire.signal = signal
	wire.computed = true

	(*circuit)[wireName] = wire

	return signal
}

func (circuit *Circuit) computeWire(wire Wire) uint16 {
	var depSignals []uint16
	for _, dep := range wire.deps {
		signal := circuit.FindSignal(dep)
		depSignals = append(depSignals, signal)
	}

	switch wire.op {
	case and:
		return depSignals[0] & depSignals[1]
	case or:
		return depSignals[0] | depSignals[1]
	case lshift:
		return depSignals[0] << depSignals[1]
	case rshift:
		return depSignals[0] >> depSignals[1]
	case not:
		return ^depSignals[0]
	case bare:
		return depSignals[0]
	default:
		panic("invalid wire op")
	}
}
