package main

import (
	"os"
	"strings"
	"testing"

	"github.com/almushel/aoc2023/utility"
)

const (
	exampleInput = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`
	example1Solution = 374
	example2Solution = 1030
	example3Solution = 8410
	part1Solution    = 10292708
	part2Solution    = 790194712336
)

var input string

func init() {
	buffer, err := os.ReadFile("../data/day11_input.txt")
	if err != nil {
		print(err.Error())
		os.Exit(1)
	}
	input = strings.TrimSpace(string(buffer))
}

func TestExample1(t *testing.T) {
	utility.PuzzleTest(t, exampleInput, example1Solution, solvePart1)
}

func Test10Expansion(t *testing.T) {
	utility.PuzzleTest(t, exampleInput, example2Solution, func(str string) int { return solve(str, 10) })
}

func Test100Expansion(t *testing.T) {
	utility.PuzzleTest(t, exampleInput, example3Solution, func(str string) int { return solve(str, 100) })
}

func TestPart1(t *testing.T) {
	utility.PuzzleTest(t, input, part1Solution, solvePart1)
}

func TestPart2(t *testing.T) {
	utility.PuzzleTest(t, input, part2Solution, solvePart2)
}
