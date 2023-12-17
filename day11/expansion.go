package main

import (
	"fmt"
	"os"
	"strings"
)

func solve(input string, expansion int) (result int) {
	lines := strings.Split(input, "\n")
	rows := make([]int, len(lines))
	cols := make([]int, len(lines[0]))
	var galaxies [][2]int

	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				galaxies = append(galaxies, [2]int{x, y})
				rows[y]++
				cols[x]++
			}
		}
	}
	var pairs [][2][2]int
	for i := range galaxies {
		for e := i + 1; e < len(galaxies); e++ {
			pairs = append(pairs, [2][2]int{galaxies[i], galaxies[e]})
		}
	}

	var start, end, dir [2]int
	var steps int
	for _, pair := range pairs {
		steps = 0
		start = pair[0]
		end = pair[1]
		if end[0] > start[0] {
			dir[0] = 1
		} else {
			dir[0] = -1
		}
		if end[1] > start[1] {
			dir[1] = 1
		} else {
			dir[1] = -1
		}

		next := start
		for next != end {
			if next[0] != end[0] {
				next[0] += dir[0]
				if cols[next[0]] == 0 {
					steps += expansion
				} else {
					steps++
				}
			}

			if next[1] != end[1] {
				next[1] += dir[1]
				if rows[next[1]] == 0 {
					steps += expansion
				} else {
					steps++
				}
			}
		}
		result += steps
	}
	return result
}

func solvePart1(input string) int {
	return solve(input, 2)
}
func solvePart2(input string) int {
	return solve(input, 1000000)
}

func main() {
	buffer, err := os.ReadFile("data/day11_input.txt")
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
