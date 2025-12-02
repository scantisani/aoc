package part1

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
		1227775554,
	},
	{
		[]string{
			"10-30",     // 11, 22
			"1010-1111", // 1010, 1111
		},
		2154,
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
	{idRange{95, 115}, []int{99}},
	{idRange{998, 1012}, []int{1010}},
	{idRange{1188511880, 1188511890}, []int{1188511885}},
	{idRange{222220, 222224}, []int{222222}},
	{idRange{1698522, 1698528}, []int{}},
	{idRange{446443, 446449}, []int{446446}},
	{idRange{38593856, 38593862}, []int{38593859}},
}

func TestInvalidIdsInRange(t *testing.T) {
	for _, test := range invalidsTests {
		received := InvalidIdsInRange(test.input)
		if !slices.Equal(test.expected, received) {
			t.Errorf(`%v should have been %v, was %v`, test.input, test.expected, received)
		}
	}
}
