package part1

import (
	"slices"
	"testing"
)

type test struct {
	input    []string
	expected int
}

var tests = []test{
	{[]string{
		"162,817,812",
		"57,618,57",
		"906,360,560",
		"592,479,940",
		"352,342,300",
		"466,668,158",
		"542,29,236",
		"431,825,988",
		"739,650,466",
		"52,470,668",
		"216,146,977",
		"819,987,18",
		"117,168,530",
		"805,96,715",
		"346,949,466",
		"970,615,88",
		"941,993,340",
		"862,61,35",
		"984,92,344",
		"425,690,689",
	}, 40},
}

func TestMultiplyCircuitLengths(t *testing.T) {
	for _, test := range tests {
		received := MultiplyCircuitLengths(test.input, 10)
		if received != test.expected {
			t.Errorf(`%v should have been %v, was %v`, test.input, test.expected, received)
		}
	}
}

type euclideanDistanceTest struct {
	input    []JunctionBox
	expected float64
}

var euclideanDistanceTests = []euclideanDistanceTest{
	{[]JunctionBox{{0, 0, 0}, {0, 0, 0}}, 0},
	{[]JunctionBox{{5, 5, 5}, {1, 2, 5}}, 5},
	{[]JunctionBox{{10, 10, 10}, {1, 2, 3}}, 13.92838827718412},
	{[]JunctionBox{{162, 817, 812}, {425, 690, 689}}, 316.90219311326956},
	{[]JunctionBox{{162, 817, 812}, {431, 825, 988}}, 321.560258738545},
	{[]JunctionBox{{906, 360, 560}, {805, 96, 715}}, 322.36935338211043},
	{[]JunctionBox{{431, 825, 988}, {425, 690, 689}}, 328.11888089532425},
}

func TestEuclideanDistance(t *testing.T) {
	for _, test := range euclideanDistanceTests {
		received := EuclideanDistance(test.input[0], test.input[1])
		if received != test.expected {
			t.Errorf(`%v should have been %v, was %v`, test.input, test.expected, received)
		}
	}
}

type closestTest struct {
	input    []JunctionBox
	expected []Connection
}

var closestTests = []closestTest{
	{[]JunctionBox{
		{162, 817, 812},
		{425, 690, 689},
		{431, 825, 988},
		{805, 96, 715},
		{906, 360, 560},
	}, []Connection{
		{box1: JunctionBox{162, 817, 812}, box2: JunctionBox{425, 690, 689}, distance: 316.90219311326956},
		{box1: JunctionBox{162, 817, 812}, box2: JunctionBox{431, 825, 988}, distance: 321.560258738545},
		{box1: JunctionBox{805, 96, 715}, box2: JunctionBox{906, 360, 560}, distance: 322.36935338211043},
		{box1: JunctionBox{425, 690, 689}, box2: JunctionBox{431, 825, 988}, distance: 328.11888089532425},
		{box1: JunctionBox{425, 690, 689}, box2: JunctionBox{906, 360, 560}, distance: 597.412755136681},
		{box1: JunctionBox{425, 690, 689}, box2: JunctionBox{805, 96, 715}, distance: 705.628797598284},
		{box1: JunctionBox{431, 825, 988}, box2: JunctionBox{906, 360, 560}, distance: 790.5909182377445},
		{box1: JunctionBox{431, 825, 988}, box2: JunctionBox{805, 96, 715}, distance: 863.6237606735933},
		{box1: JunctionBox{162, 817, 812}, box2: JunctionBox{906, 360, 560}, distance: 908.7843528582565},
		{box1: JunctionBox{162, 817, 812}, box2: JunctionBox{805, 96, 715}, distance: 970.9268767523124},
	}},
}

func TestClosestBoxes(t *testing.T) {
	for _, test := range closestTests {
		received := ClosestBoxes(test.input)
		if !slices.Equal(received, test.expected) {
			t.Errorf(`%v should have been %v, was %v`, test.input, test.expected, received)
		}
	}
}
