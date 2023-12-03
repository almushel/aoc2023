// Day 1, Puzzle 1: Trebouchet?!
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/almushel/aoc2023/utility"
)

func solvePart1(input string) (result int) {
	var first, last rune

	for _, c := range input {
		if c >= '0' && c <= '9' {
			if first == '\000' {
				first = rune(c)
			}
			last = rune(c)
		} else if c == '\n' {
			val, _ := utility.RunesToInt(first, last)
			result += int(val)
			first, last = '\000', '\000'
		}
	}
	val, _ := utility.RunesToInt(first, last)

	result += val
	return
}

func solvePart2(input string) (result int) {
	var digitWords []string = []string{
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}

	var first, last rune

	for i, c := range input {
		found := false
		for num, word := range digitWords {
			if word[0] != byte(c) {
				continue
			}

			cword := input[i:min(i+len(word), len(input))]
			if cword == word {
				if first == '\000' {
					first = '0' + rune(num+1)
				}
				last = '0' + rune(num+1)
				found = true
				break
			}
		}
		if found {
			continue
		}

		if c >= '0' && c <= '9' {
			if first == '\000' {
				first = rune(c)
			}
			last = rune(c)
		} else if c == '\n' {
			val, _ := utility.RunesToInt(first, last)
			result += int(val)
			first, last = '\000', '\000'
		}
	}
	val, _ := utility.RunesToInt(first, last)

	result += val
	return
}

func main() {

	buffer, err := os.ReadFile("data/day1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileStr := string(buffer)
	part1Result := solvePart1(fileStr)
	part2Result := solvePart2(fileStr)

	fmt.Println("Part 1: ", part1Result)
	fmt.Println("Part 2: ", part2Result)
}
