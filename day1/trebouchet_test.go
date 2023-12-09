package main

import (
	"os"
	"strings"
	"testing"

	"github.com/almushel/aoc2023/utility"
)

const (
	testInput1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	testInput2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
	exampleSolution1 = 142
	exampleSolution2 = 281
	part1Solution    = 55621
	part2Solution    = 53592
)

var input string

func init() {
	buffer, err := os.ReadFile("../data/day1_input.txt")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	input = strings.TrimSpace(string(buffer))
}

func TestExample1(t *testing.T) {
	utility.PuzzleTest(t, testInput1, exampleSolution1, solvePart1)
}

func TestExample2(t *testing.T) {
	utility.PuzzleTest(t, testInput2, exampleSolution2, solvePart2)
}

func TestPart1(t *testing.T) {
	utility.PuzzleTest(t, input, part1Solution, solvePart1)
}

func TestPart2(t *testing.T) {
	utility.PuzzleTest(t, input, part2Solution, solvePart2)
}
