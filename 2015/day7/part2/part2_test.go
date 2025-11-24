package part2

import (
	"testing"
)

type test struct {
	input    []string
	expected map[string]uint16
}

var tests = []test{
	{
		[]string{
			"123 -> x",
			"456 -> y",
			"x AND y -> d",
			"x OR y -> e",
			"x LSHIFT 2 -> f",
			"y RSHIFT 2 -> g",
			"NOT x -> h",
			"NOT y -> i",
		},
		map[string]uint16{
			"d": 72,
			"e": 507,
			"f": 492,
			"g": 114,
			"h": 65412,
			"i": 65079,
			"x": 123,
			"y": 456,
		},
	},
	{
		[]string{
			"x AND y -> d",
			"x OR y -> e",
			"x LSHIFT 2 -> f",
			"y RSHIFT 2 -> g",
			"NOT x -> h",
			"NOT y -> i",
			"123 -> x",
			"456 -> y",
		},
		map[string]uint16{
			"d": 72,
			"e": 507,
			"f": 492,
			"g": 114,
			"h": 65412,
			"i": 65079,
			"x": 123,
			"y": 456,
		},
	},
	{
		[]string{
			"123 -> x",
			"456 -> y",
			"x AND 1 -> d",
			"1 OR y -> e",
		},
		map[string]uint16{
			"d": 1,
			"e": 457,
			"x": 123,
			"y": 456,
		},
	},
}

func TestAssembleCircuit(t *testing.T) {
	for _, test := range tests {
		circuit := AssembleCircuit(test.input)

		for name, expectedSignal := range test.expected {
			receivedSignal := circuit.FindSignal(name)
			if receivedSignal != expectedSignal {
				t.Errorf(`%v should have been %v, was %v`, name, expectedSignal, receivedSignal)
			}
		}
	}
}
