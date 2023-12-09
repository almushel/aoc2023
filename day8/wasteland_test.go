package main

import (
	"os"
	"strings"
	"testing"

	"github.com/almushel/aoc2023/utility"
)

const (
	testInput1 = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

	testInput2 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`
	exampleSolution1 = 2
	exampleSolution2 = 6
	part1Solution    = 19951
	part2Solution    = 16342438708751
)

var input string

func init() {
	buffer, err := os.ReadFile("../data/day8_input.txt")
	if err != nil {
		println(err.Error)
		os.Exit(1)
	}
	input = strings.TrimSpace(string(buffer))
}
func TestPart1Example(t *testing.T) {
	utility.PuzzleTest(t, testInput1, exampleSolution1, SolvePart1)
}

func TestPart2Example(t *testing.T) {
	utility.PuzzleTest(t, testInput2, exampleSolution2, SolvePart2)
}

func TestPart1(t *testing.T) {
	utility.PuzzleTest(t, input, part1Solution, SolvePart1)
}

func TestPart2(t *testing.T) {
	utility.PuzzleTest(t, input, part2Solution, SolvePart2)
}
