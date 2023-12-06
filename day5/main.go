package main

import (
	"fmt"
	"os"

	"github.com/almushel/aoc2023/day5/fertilizer"
)

func main() {
	buffer, _ := os.ReadFile("data/day5_input.txt")
	str := string(buffer)
	seedNums, almanac := fertilizer.ParseAlmanac(str)
	part1Result := fertilizer.SolvePart1(seedNums, almanac)
	part2Result := fertilizer.PoorlySolvePart2(seedNums, almanac)

	fmt.Println("Part 1:", part1Result)
	fmt.Println("Part 2:", part2Result)
}
