package main

import (
	"testing"

	"github.com/almushel/aoc2023/utility"
)

const (
	exampleInput = `Time:      7  15   30
Distance:  9  40  200`

	input = `Time:        47     70     75     66
Distance:   282   1079   1147   1062`
	exampleSolution1 = 288
	exampleSolution2 = 71503
	part1Solution    = 281600
	part2Solution    = 33875953
)

func TestPart1Example(t *testing.T) {
	utility.PuzzleTest(t, exampleInput, exampleSolution1, SolvePart1)
}

func TestPart2Example(t *testing.T) {
	utility.PuzzleTest(t, exampleInput, exampleSolution2, SolvePart2)
}

func TestPart1(t *testing.T) {
	utility.PuzzleTest(t, input, part1Solution, SolvePart1)
}

func TestPart2(t *testing.T) {
	utility.PuzzleTest(t, input, part2Solution, SolvePart2)
}
