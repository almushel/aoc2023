package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/almushel/aoc2023/utility"
)

func lastElement[T any](list []T) *T {
	return &list[len(list)-1]
}

func parseInput(input string) (histories [][]int) {
	for _, hStr := range strings.Split(input, "\n") {
		var history []int
		for _, vStr := range strings.Split(hStr, " ") {
			val, err := utility.StringToInt(vStr)
			if err != nil {
				print(err.Error())
			}
			history = append(history, val)
		}
		histories = append(histories, history)
	}
	return
}
func getDiffs(nums []int) (diffs []int, zeros bool) {
	zeros = true
	if len(nums) == 0 {
		return
	}
	diffs = make([]int, len(nums)-1)
	for i := 0; i < len(diffs); i++ {
		diffs[i] = nums[i+1] - nums[i]
		if diffs[i] != 0 {
			zeros = false
		}
	}
	return
}

func getNext(history []int) int {
	diffs := [][]int{history}
	diff := history
	done := false
	for !done {
		diff, done = getDiffs(diff)
		diffs = append(diffs, diff)
	}

	*lastElement(diffs) = append(*lastElement(diffs), 0)
	for i := len(diffs) - 2; i >= 0; i-- {
		diffs[i] = append(diffs[i], 0)
		*lastElement(diffs[i]) = diffs[i][len(diffs[i])-2]
		*lastElement(diffs[i]) += *lastElement(diffs[i+1])
	}

	return *lastElement(diffs[0])
}

func solvePart1(input string) (result int) {
	histories := parseInput(input)
	for _, history := range histories {
		result += getNext(history)
	}

	return
}

func solvePart2(input string) (result int) {
	histories := parseInput(input)
	for _, history := range histories {
		slices.Reverse(history)
		result += getNext(history)
	}

	return
}

func main() {
	buffer, err := os.ReadFile("data/day9_input.txt")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	str := strings.TrimSpace(string(buffer))

	part1Result := solvePart1(str)
	part2Result := solvePart2(str)

	fmt.Println("Part 1:", part1Result)
	fmt.Println("Part 2:", part2Result)

	return
}
