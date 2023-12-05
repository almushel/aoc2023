package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/almushel/aoc2023/day4/scratchcards"
)

func main() {
	buffer, _ := os.ReadFile("data/day4_input.txt")
	str := strings.TrimSpace(string(buffer))
	cards := scratchcards.ParseCards(str)

	part1Result := scratchcards.SolvePart1(cards)
	part2Result := scratchcards.SolvePart2(cards)

	fmt.Println("Part 1:", part1Result)
	fmt.Println("Part 2:", part2Result)
}
