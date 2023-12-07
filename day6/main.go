package main

import (
	"fmt"

	"github.com/almushel/aoc2023/day6/boatrace"
)

func main() {
	str :=
		`Time:        47     70     75     66
Distance:   282   1079   1147   1062`
	part1Result := boatrace.SolvePart1(str)
	part2Result := boatrace.SolvePart2(str)

	fmt.Println(part1Result)
	fmt.Println(part2Result)
	return
}
