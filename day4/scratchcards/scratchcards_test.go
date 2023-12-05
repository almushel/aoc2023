package scratchcards

import (
	"log"
	"os"
	"strings"
	"testing"
)

const (
	part1Solution = 22674
	part2Solution = 5747443
)

var cards map[int]int

func init() {
	buffer, err := os.ReadFile("../../data/day4_input.txt")
	if err != nil {
		log.Fatalln(err.Error())
	}
	str := strings.TrimSpace(string(buffer))
	cards = ParseCards(str)
}

func TestSolvePart1(t *testing.T) {
	want := part1Solution
	result := SolvePart1(cards)
	if result != want {
		t.Fatalf("result = %d, want: %d),", result, want)
	}
}

func TestPart2RecursionList(t *testing.T) {
	want := part2Solution
	var newCards []int
	for c, _ := range cards {
		Part2ProcessR(cards, c, &newCards)
	}

	result := len(newCards)
	if result != part2Solution {
		t.Fatalf("result = %d, want: %d),", result, want)
	}
}

func TestPart2Recursion(t *testing.T) {
	want := part2Solution
	result := 0
	for c, _ := range cards {
		result += Part2ProcessR2(cards, c)
	}

	if result != part2Solution {
		t.Fatalf("result = %d, want: %d),", result, want)
	}
}

func TestPart2Queue(t *testing.T) {
	want := part2Solution
	var result int
	for c, _ := range cards {
		result += Part2ProcessQueue(cards, c)
	}

	if result != part2Solution {
		t.Fatalf("result = %d, want: %d),", result, want)
	}
}
