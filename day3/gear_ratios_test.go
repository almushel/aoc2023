package main

import (
	"os"
	"testing"

	"github.com/almushel/aoc2023/utility"
)

const (
	testStr = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	exampleSolution1 = 4361
	exampleSolution2 = 467835
	part1Solution    = 546563
	part2Solution    = 91031374
)

var input string

func init() {
	buffer, err := os.ReadFile("../data/day3_input.txt")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	input = string(buffer)
}

func TestExample1(t *testing.T) {
	utility.PuzzleTest(t, testStr, exampleSolution1, solvePart1)
}

func TestExample2(t *testing.T) {
	utility.PuzzleTest(t, testStr, exampleSolution2, solvePart2)
}

func TestPart1(t *testing.T) {
	utility.PuzzleTest(t, input, part1Solution, solvePart1)
}

func TestPart2(t *testing.T) {
	utility.PuzzleTest(t, input, part2Solution, solvePart2)
}
