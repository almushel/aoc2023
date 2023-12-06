package fertilizer

import (
	"os"
	"testing"
)

const (
	exampleSolution1 = 35
	exampleSolution2 = 46
	part1Solution    = 379811651
	part2Solution    = 27992443
)

var testStr string = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

var seedNumbers []int
var almanac [][][3]int

func init() {
	buffer, err := os.ReadFile("../../data/day5_input.txt")
	if err != nil {
		println(err.Error())
		return
	}
	str := string(buffer)
	seedNumbers, almanac = ParseAlmanac(str)
}

func TestPart1Example(t *testing.T) {
	seedNumbers, almanac := ParseAlmanac(testStr)
	want := exampleSolution1
	result := SolvePart1(seedNumbers, almanac)
	if result != want {
		t.Fatalf("result = %d, want: %d),", result, want)
	}
}

func TestPart2Example(t *testing.T) {
	seedNumbers, almanac := ParseAlmanac(testStr)
	want := exampleSolution2
	result := PoorlySolvePart2(seedNumbers, almanac)
	if result != want {
		t.Fatalf("result = %d, want: %d),", result, want)
	}
}

func TestPart1(t *testing.T) {
	want := part1Solution
	result := SolvePart1(seedNumbers, almanac)
	if result != want {
		t.Fatalf("result = %d, want: %d),", result, want)
	}
}

/*
func TestPart2(t *testing.T) {
	want := part2Solution
	result := PoorlySolvePart2(seedNumbers, almanac)
	if result != want {
		t.Fatalf("result = %d, want: %d),", result, want)
	}
}
*/
