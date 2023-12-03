package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/almushel/aoc2023/utility"
)

func validPart(lines []string, lineIndex, start, end int) bool {
	for _, line := range lines[max(lineIndex-1, 0):min(lineIndex+2, len(lines))] {
		for _, checkChar := range line[max(start-1, 0):min(end+2, len(line))] {
			if checkChar != '.' && !utility.IsNumeric(checkChar) {
				return true
			}
		}
	}
	return false
}

func solvePart1(input string) (result int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for i, line := range lines {
		start, end := -1, -1
		for c, char := range line {
			endNum := false
			if utility.IsNumeric(char) {
				if start < 0 {
					start = c
				}
				end = c
			} else {
				endNum = true
			}

			if c == len(line)-1 {
				endNum = true
			}

			if endNum && start >= 0 {
				if validPart(lines, i, start, end) {
					num, _ := utility.StringToInt(line[start : end+1])
					result += num
				}

				start, end = -1, -1
			}
		}
	}
	return
}

func adjacentGears(lines []string, lineIndex, start, end int) (gears []int) {
	for l := max(lineIndex-1, 0); l < min(lineIndex+2, len(lines)); l++ {
		line := lines[l]
		for c := max(start-1, 0); c < min(end+2, len(line)); c++ {
			checkChar := line[c]
			if checkChar == '*' {
				gears = append(gears, l*len(line)+c)
			}
		}
	}

	return
}

func solvePart2(input string) (result int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	potentialGears := make(map[int][]int)
	for i, line := range lines {
		start, end := -1, -1
		for c, char := range line {
			endNum := false
			if utility.IsNumeric(char) {
				if start < 0 {
					start = c
				}
				end = c
			} else {
				endNum = true
			}

			if c == len(line)-1 {
				endNum = true
			}

			if endNum && start >= 0 {
				newGears := adjacentGears(lines, i, start, end)
				partNumber, _ := utility.StringToInt(line[start : end+1])
				for _, gear := range newGears {
					potentialGears[gear] = append(potentialGears[gear], partNumber)
				}
				start, end = -1, -1
			}
		}
	}

	for _, nums := range potentialGears {
		if len(nums) == 2 {
			result += nums[0] * nums[1]
		}
	}

	return
}

func main() {
	buffer, err := os.ReadFile("data/day3_input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	str := string(buffer)

	part1Result := solvePart1(str)
	fmt.Println("Part 1:", part1Result)
	part2Result := solvePart2(str)
	fmt.Println("Part 2:", part2Result)
	return
}
