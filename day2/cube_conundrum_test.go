package main

import (
	"os"
	"strings"
	"testing"

	"github.com/almushel/aoc2023/utility"
)

const (
	testStr = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	exampleSolution1 = 8
	exampleSolution2 = 2286
	part1Solution    = 2545
	part2Solution    = 78111
)

var testInput []gameResult
var input []gameResult

func init() {
	testInput = parseGameResults(testStr)
	buffer, err := os.ReadFile("../data/day2_input.txt")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	str := strings.TrimSpace(string(buffer))
	input = parseGameResults(str)
}

func TestExample1(t *testing.T) {
	utility.PuzzleTest(t, "", exampleSolution1, func(string) int { return solvePart1(testInput) })
}

func TestExample2(t *testing.T) {
	utility.PuzzleTest(t, "", exampleSolution2, func(string) int { return solvePart2(testInput) })
}

func TestPart1(t *testing.T) {
	utility.PuzzleTest(t, "", part1Solution, func(string) int { return solvePart1(input) })
}

func TestPart2(t *testing.T) {
	utility.PuzzleTest(t, "", part2Solution, func(string) int { return solvePart2(input) })
}
