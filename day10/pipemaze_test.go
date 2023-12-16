package main

import (
	"os"
	"strings"
	"testing"

	"github.com/almushel/aoc2023/utility"
)

const (
	testInput1 = `.....
.S-7.
.|.|.
.L-J.
.....`
	testInput2 = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`
	testInput3 = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........
`
	testInput4 = `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`
	testInput5 = `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`
	exampleSolution1 = 4
	exampleSolution2 = 8
	exampleSolution3 = 4
	exampleSolution4 = 8
	exampleSolution5 = 10
	part1Solution    = 6613
	part2Solution    = 511
)

var input string

func init() {
	buffer, err := os.ReadFile("../data/day10_input.txt")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	input = strings.TrimSpace(string(buffer))
}

func TestPart1Example1(t *testing.T) {
	utility.PuzzleTest(t, testInput1, exampleSolution1, solvePart1)
}

func TestPart1Example2(t *testing.T) {
	utility.PuzzleTest(t, testInput2, exampleSolution2, solvePart1)
}

func TestPart2Example1(t *testing.T) {
	utility.PuzzleTest(t, testInput3, exampleSolution3, solvePart2)
}

func TestPart2Example2(t *testing.T) {
	utility.PuzzleTest(t, testInput4, exampleSolution4, solvePart2)
}

func TestPart2Example3(t *testing.T) {
	utility.PuzzleTest(t, testInput5, exampleSolution5, solvePart2)
}

func TestPart1(t *testing.T) {
	utility.PuzzleTest(t, input, part1Solution, solvePart1)
}

func TestPart2(t *testing.T) {
	utility.PuzzleTest(t, input, part2Solution, solvePart2)
}
