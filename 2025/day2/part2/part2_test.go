package part2

import (
	"slices"
	"testing"
)

type sumTest struct {
	input    []string
	expected int
}

var sumTests = []sumTest{
	{
		[]string{
			"11-22",
			"95-115",
			"998-1012",
			"1188511880-1188511890",
			"222220-222224",
			"1698522-1698528",
			"446443-446449",
			"38593856-38593862",
			"565653-565659",
			"824824821-824824827",
			"2121212118-2121212124",
		},
		4174379265,
	},
}

func TestSumOfInvalidIds(t *testing.T) {
	for _, test := range sumTests {
		received := SumOfInvalidIds(test.input)
		if received != test.expected {
			t.Errorf(`%v should have been %v, was %v`, test.input, test.expected, received)
		}
	}
}

type invalidsTest struct {
	input    idRange
	expected []int
}

var invalidsTests = []invalidsTest{
	{idRange{11, 22}, []int{11, 22}},
	{idRange{95, 115}, []int{99, 111}},
	{idRange{998, 1012}, []int{999, 1010}},
	{idRange{222220, 222224}, []int{222222}},
	{idRange{446443, 446449}, []int{446446}},
	{idRange{565653, 565659}, []int{565656}},
	{idRange{1698522, 1698528}, []int{}},
	{idRange{38593856, 38593862}, []int{38593859}},
	{idRange{824824821, 824824827}, []int{824824824}},
	{idRange{1188511880, 1188511890}, []int{1188511885}},
	{idRange{2121212118, 2121212124}, []int{2121212121}},
}

func TestInvalidIdsInRange(t *testing.T) {
	for _, test := range invalidsTests {
		received := InvalidIdsInRange(test.input)
		if !slices.Equal(test.expected, received) {
			t.Errorf(`%v should have been %v, was %v`, test.input, test.expected, received)
		}
	}
}

type splitsTest struct {
	input    string
	numParts int
	expected []string
}

var splitsTests = []splitsTest{
	{"111", 3, []string{"1", "1", "1"}},
	{"1212", 2, []string{"12", "12"}},
	{"1234567890", 5, []string{"12", "34", "56", "78", "90"}},
	{"1234567890", 2, []string{"12345", "67890"}},
}

func TestSplitIntoParts(t *testing.T) {
	for _, test := range splitsTests {
		received := SplitIntoParts(test.input, test.numParts)
		if !slices.Equal(test.expected, received) {
			t.Errorf(`%v should have been %v, was %v`, test.input, test.expected, received)
		}
	}
}
