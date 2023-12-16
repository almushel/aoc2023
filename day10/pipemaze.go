package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/almushel/aoc2023/utility"
)

var pipes = map[rune][2][2]int{
	//| is a vertical pipe connecting north and south.
	'|': {{0, -1}, {0, 1}},
	//- is a horizontal pipe connecting east and west.
	'-': {{-1, 0}, {1, 0}},
	//L is a 90-degree bend connecting north and east.
	'L': {{0, -1}, {1, 0}},
	//J is a 90-degree bend connecting north and west.
	'J': {{0, -1}, {-1, 0}},
	//7 is a 90-degree bend connecting south and west.
	'7': {{0, 1}, {-1, 0}},
	//F is a 90-degree bend connecting south and east.
	'F': {{0, 1}, {1, 0}},
}

func getExit(node rune, entrance [2]int) [2]int {
	exits, ok := pipes[node]
	if ok {
		if exits[0] == entrance {
			return exits[1]
		} else {
			return exits[0]
		}
	}

	return [2]int{0, 0}
}

func parseInput(input string) (result []string, start [2]int) {
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		lb := []rune(line)
		for x, char := range line {
			if char == 'S' {
				start = [2]int{x, y}
				var exits [][2]int
				for i := -1; i < 2; i += 2 {
					neighbor := lines[utility.Clamp(y+i, 0, len(lines))][x]
					_, ok := pipes[rune(neighbor)]
					if neighbor != 'S' && ok {
						exits = append(exits, [2]int{0, i})
					}
					neighbor = lines[y][utility.Clamp(x+i, 0, len(lines[y]))]
					_, ok = pipes[rune(neighbor)]
					if neighbor != 'S' && ok {
						exits = append(exits, [2]int{i, 0})
					}
				}
				for key, val := range pipes {
					if exits[0] != val[0] && exits[0] != val[1] {
						continue
					}
					if exits[1] != val[0] && exits[1] != val[1] {
						continue
					}
					lb[x] = key
				}
			}
		}
		result = append(result, string(lb))
	}
	return
}

func solvePart1(input string) (result int) {
	lines, start := parseInput(input)
	char := rune(lines[start[1]][start[0]])

	steps := 1
	prev := start
	// NOTE: Secretly chose the right direction to start in
	// TO-DO: Maybe actually properly solve this
	exit := pipes[char][0]
	next := [2]int{start[0] + exit[0], start[1] + exit[1]}

	for next != start {
		char = rune(lines[next[1]][next[0]])
		exits := pipes[char]
		if exits[0] == [2]int{prev[0] - next[0], prev[1] - next[1]} {
			prev = next
			next = [2]int{next[0] + exits[1][0], next[1] + exits[1][1]}
		} else {
			prev = next
			next = [2]int{next[0] + exits[0][0], next[1] + exits[0][1]}
		}
		steps++
	}
	result = steps / 2

	return
}

func markPath(input string) [][]int {
	lines, start := parseInput(input)
	result := make([][]int, len(lines))
	for y, line := range lines {
		result[y] = make([]int, len(line))
	}
	char := rune(lines[start[1]][start[0]])

	for _, e := range pipes[char] {
		prev := start
		next := [2]int{start[0] + e[0], start[1] + e[1]}

		result[start[1]][start[0]] = 1
		result[next[1]][next[0]] = 1

		for next != start {
			char = rune(lines[next[1]][next[0]])
			exits := pipes[char]
			if exits[0] == [2]int{prev[0] - next[0], prev[1] - next[1]} {
				prev = next
				next = [2]int{next[0] + exits[1][0], next[1] + exits[1][1]}
			} else {
				prev = next
				next = [2]int{next[0] + exits[0][0], next[1] + exits[0][1]}
			}
			result[next[1]][next[0]] = 1
		}
	}

	return result
}

func solvePart2(input string) (result int) {
	lines, _ := parseInput(input)
	path := markPath(input)

	for y, line := range lines {
		for x := range line {
			if path[y][x] > 0 {
				continue
			}
			even := true
			for i := x + 1; i < len(line); i++ {
				if path[y][i] > 0 {
					check := rune(line[i])
					if strings.Contains("FL", string(check)) {
						i++
						for i < len(line) {
							if strings.Contains("J7", string(line[i])) {
								if check == 'F' && line[i] == 'J' {
									even = !even
								} else if check == 'L' && line[i] == '7' {
									even = !even
								}
								break
							}
							i++
						}
					} else {
						even = !even
					}
				}
			}

			if !even {
				result++
			}
		}
	}

	return
}

func main() {
	buffer, err := os.ReadFile("data/day10_input.txt")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	input := strings.TrimSpace(string(buffer))

	part1Result := solvePart1(input)
	part2Result := solvePart2(input)

	fmt.Println("Part 1:", part1Result)
	fmt.Println("Part 2:", part2Result)
	return
}
