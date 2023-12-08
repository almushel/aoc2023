package main

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/almushel/aoc2023/day4/scratchcards"
	"github.com/almushel/aoc2023/utility"
)

const (
	part1Solution = 22674
	part2Solution = 5747443
)

var cards map[int]int

func init() {
	buffer, err := os.ReadFile("../data/day4_input.txt")
	if err != nil {
		log.Fatalln(err.Error())
	}
	str := strings.TrimSpace(string(buffer))
	cards = scratchcards.ParseCards(str)
}

func TestSolvePart1(t *testing.T) {
	utility.PuzzleTest(t, "", part1Solution, func(string) int { return scratchcards.SolvePart1(cards) })
}

func TestPart2RecursionList(t *testing.T) {
	solve := func(string) int {
		var newCards []int
		for c, _ := range cards {
			scratchcards.Part2ProcessR(cards, c, &newCards)
		}
		return len(newCards)
	}
	utility.PuzzleTest(t, "", part2Solution, solve)
}

func TestPart2Recursion(t *testing.T) {
	solve := func(string) (result int) {
		for c, _ := range cards {
			result += scratchcards.Part2ProcessR2(cards, c)
		}
		return
	}
	utility.PuzzleTest(t, "", part2Solution, solve)
}

func TestPart2Queue(t *testing.T) {
	solve := func(string) (result int) {
		for c, _ := range cards {
			result += scratchcards.Part2ProcessQueue(cards, c)
		}
		return
	}
	utility.PuzzleTest(t, "", part2Solution, solve)
}
