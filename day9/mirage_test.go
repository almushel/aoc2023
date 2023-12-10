package main

import (
	"os"
	"strings"
	"testing"

	"github.com/almushel/aoc2023/utility"
)

const (
	testInput = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`
	exampleSolution1 = 114
	exampleSolution2 = 2
	part1Solution    = 2174807968
	part2Solution    = 1208
)

var input string

func init() {
	buffer, err := os.ReadFile("../data/day9_input.txt")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	input = strings.TrimSpace(string(buffer))
}

func TestExample1(t *testing.T) {
	utility.PuzzleTest(t, testInput, exampleSolution1, solvePart1)
}

func TestExample2(t *testing.T) {
	utility.PuzzleTest(t, testInput, exampleSolution2, solvePart2)
}

func TestPart1(t *testing.T) {
	utility.PuzzleTest(t, input, part1Solution, solvePart1)
}

func TestPart2(t *testing.T) {
	utility.PuzzleTest(t, input, part2Solution, solvePart2)
}
