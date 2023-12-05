package part1

import (
	"reflect"
	"testing"
)

func TestLowestLocationNumber(t *testing.T) {
	tests := []struct {
		lines []string
		want  int
	}{
		{
			[]string{
				"seeds: 79 14 55 13",
				"",
				"seed-to-soil map:",
				"50 98 2",
				"52 50 48",
				"",
				"soil-to-fertilizer map:",
				"0 15 37",
				"37 52 2",
				"39 0 15",
				"",
				"fertilizer-to-water map:",
				"49 53 8",
				"0 11 42",
				"42 0 7",
				"57 7 4",
				"",
				"water-to-light map:",
				"88 18 7",
				"18 25 70",
				"",
				"light-to-temperature map:",
				"45 77 23",
				"81 45 19",
				"68 64 13",
				"",
				"temperature-to-humidity map:",
				"0 69 1",
				"1 0 69",
				"",
				"humidity-to-location map:",
				"60 56 37",
				"56 93 4",
				"",
			},
			35,
		},
	}
	for _, tt := range tests {
		t.Run("AOC Example input", func(t *testing.T) {
			if got := LowestLocationNumber(tt.lines); got != tt.want {
				t.Errorf("LowestLocationNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseMap(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  AlmanacMap
	}{
		{
			"seed-to-soil map",
			[]string{
				"seed-to-soil map:",
				"50 98 2",
				"52 50 48",
			},
			AlmanacMap{
				source:      "seed",
				destination: "soil",
				conversions: []Conversion{
					{sourceStart: 50, sourceEnd: 97, difference: 2},
					{sourceStart: 98, sourceEnd: 99, difference: -48},
				},
			},
		},
		{
			"light-to-temperature map",
			[]string{
				"light-to-temperature map:",
				"45 77 23",
				"81 45 19",
				"68 64 13",
			},
			AlmanacMap{
				source:      "light",
				destination: "temperature",
				conversions: []Conversion{
					{sourceStart: 64, sourceEnd: 76, difference: 4},
					{sourceStart: 45, sourceEnd: 63, difference: 36},
					{sourceStart: 77, sourceEnd: 99, difference: -32},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseMap(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMap(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestParseConversion(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Conversion
	}{
		{"short range with big difference", "50 98 2", Conversion{sourceStart: 98, sourceEnd: 99, difference: -48}},
		{"long range with small difference", "52 50 48", Conversion{sourceStart: 50, sourceEnd: 97, difference: 2}},
		{"long range with big difference", "45 77 23", Conversion{sourceStart: 77, sourceEnd: 99, difference: -32}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseConversion(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMap(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestParseAlamanac(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  Almanac
	}{
		{
			"two simple",
			[]string{
				"seed-to-soil map:",
				"50 98 2",
				"52 50 48",
				"",
				"light-to-temperature map:",
				"45 77 23",
				"81 45 19",
				"68 64 13",
				"",
			},
			Almanac{
				"seed": AlmanacMap{
					source:      "seed",
					destination: "soil",
					conversions: []Conversion{
						{sourceStart: 50, sourceEnd: 97, difference: 2},
						{sourceStart: 98, sourceEnd: 99, difference: -48},
					},
				},
				"light": AlmanacMap{
					source:      "light",
					destination: "temperature",
					conversions: []Conversion{
						{sourceStart: 64, sourceEnd: 76, difference: 4},
						{sourceStart: 45, sourceEnd: 63, difference: 36},
						{sourceStart: 77, sourceEnd: 99, difference: -32},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseAlmanac(tt.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseAlmanac() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocationForSeed(t *testing.T) {
	tests := []struct {
		name    string
		seed    int
		almanac Almanac
		want    int
	}{
		{
			"All have mappings",
			79,
			Almanac{
				"seed": AlmanacMap{
					source:      "seed",
					destination: "soil",
					conversions: []Conversion{{sourceStart: 70, sourceEnd: 80, difference: 1}},
				},
				"soil": AlmanacMap{
					source:      "soil",
					destination: "location",
					conversions: []Conversion{{sourceStart: 80, sourceEnd: 90, difference: -20}},
				},
			},
			60,
		},
		{
			"Some have no mapping",
			79,
			Almanac{
				"seed": AlmanacMap{
					source:      "seed",
					destination: "soil",
					conversions: []Conversion{{70, 71, 1}}},
				"soil": AlmanacMap{
					source:      "soil",
					destination: "location",
					conversions: []Conversion{{75, 90, -20}},
				}},
			59,
		},
		{
			"Multiple conversions",
			79,
			Almanac{
				"seed": AlmanacMap{
					source:      "seed",
					destination: "soil",
					conversions: []Conversion{{0, 5, 1}, {70, 80, 1}}},
				"soil": AlmanacMap{
					source:      "soil",
					destination: "location",
					conversions: []Conversion{{5, 15, -20}, {80, 90, -20}},
				}},
			60,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LocationForSeed(tt.seed, tt.almanac); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseAlmanac(%d, %v) = %v, want %v", tt.seed, tt.almanac, got, tt.want)
			}
		})
	}
}
