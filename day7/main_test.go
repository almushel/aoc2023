package main

import (
	"os"
	"strings"
	"testing"

	"github.com/almushel/aoc2023/utility"
)

const (
	exampleSolution1 = 6440
	exampleSolution2 = 5905
	part1Solution    = 253205868
	part2Solution    = 253907829
)

var testInput string = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

var input string

func init() {
	buffer, err := os.ReadFile("../data/day7_input.txt")
	if err != nil {
		println(err.Error())
		return
	}
	input = strings.TrimSpace(string(buffer))
}

func TestPart1Example(t *testing.T) {
	utility.PuzzleTest(t, testInput, exampleSolution1, SolvePart1)
}

func TestPart2Example(t *testing.T) {
	utility.PuzzleTest(t, testInput, exampleSolution2, SolvePart2)
}

func TestPart1Solution(t *testing.T) {
	utility.PuzzleTest(t, input, part1Solution, SolvePart1)
}

func TestPart2Solution(t *testing.T) {
	utility.PuzzleTest(t, input, part2Solution, SolvePart2)
}
