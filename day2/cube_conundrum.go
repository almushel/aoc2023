package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/almushel/aoc2023/utility"
)

type gameResult struct {
	id     int
	rounds []map[string]int
}

func parseGameResults(input string) (results []gameResult) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	results = make([]gameResult, len(lines))
	for i, line := range lines {
		colonIndex := strings.IndexRune(line, ':')
		gameID, _ := utility.StringToInt(line[len("Game "):colonIndex])
		results[i].id = gameID

		roundsSplit := strings.Split(line[colonIndex+1:], ";")
		for _, s := range roundsSplit {
			colorCounts := strings.Split(s, ",")

			round := make(map[string]int)
			for _, c := range colorCounts {
				c = strings.TrimSpace(c)
				space := strings.IndexRune(c, ' ')
				count, _ := utility.StringToInt(c[:space])
				color := strings.TrimSpace(c[space+1:])

				round[color] = count
			}
			results[i].rounds = append(results[i].rounds, round)
		}

	}
	return results
}

func solvePart1(input []gameResult) (result int) {
	const (
		RED_MAX   = 12
		GREEN_MAX = 13
		BLUE_MAX  = 14
	)

	for _, game := range input {
		valid := true
		for _, round := range game.rounds {
			if round["red"] > RED_MAX || round["green"] > GREEN_MAX || round["blue"] > BLUE_MAX {
				valid = false
				break
			}
		}
		if valid {
			result += game.id
		}

	}
	return
}

func solvePart2(input []gameResult) (result int) {
	for _, game := range input {
		gameMax := make(map[string]int)

		for _, round := range game.rounds {
			for key, val := range round {
				m, ok := gameMax[key]
				if !ok {
					gameMax[key] = val
				} else if val > m {
					gameMax[key] = val
				}
			}
		}

		power := 1
		for _, count := range gameMax {
			power *= count
		}

		result += power
	}

	return
}

func main() {
	buffer, _ := os.ReadFile("data/day2_input.txt")
	str := string(buffer)
	parsedInput := parseGameResults(str)
	part1Result := solvePart1(parsedInput)
	part2Result := solvePart2(parsedInput)

	fmt.Println("Part 1: ", part1Result)
	fmt.Println("Part 2: ", part2Result)
}
