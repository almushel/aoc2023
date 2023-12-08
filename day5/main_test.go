package main

import (
	"os"
	"testing"

	"github.com/almushel/aoc2023/utility"
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

soil-to-map:
0 15 37
37 52 2
39 0 15

to-water map:
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

type puzzleInput struct {
	seeds   []int
	almanac [][][3]int
}

var testInput puzzleInput
var input puzzleInput

func init() {
	buffer, err := os.ReadFile("../data/day5_input.txt")
	if err != nil {
		println(err.Error())
		return
	}
	str := string(buffer)
	input.seeds, input.almanac = ParseAlmanac(str)
	testInput.seeds, testInput.almanac = ParseAlmanac(testStr)
}

func TestPart1Example(t *testing.T) {
	utility.PuzzleTest(t, "", exampleSolution1, func(string) int { return SolvePart1(testInput.seeds, testInput.almanac) })
}

func TestPart2Example(t *testing.T) {
	utility.PuzzleTest(t, "", exampleSolution2, func(string) int { return PoorlySolvePart2(testInput.seeds, testInput.almanac) })
}

func TestPart1(t *testing.T) {
	utility.PuzzleTest(t, "", part1Solution, func(string) int { return SolvePart1(input.seeds, input.almanac) })
}

/*
func TestPart2(t *testing.T) {
	utility.PuzzleTest(t, "", part2Solution, func(string) int { return PoorlySolvePart2(input.seeds, input.almanac) })
}
*/
