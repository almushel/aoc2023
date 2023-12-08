package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/almushel/aoc2023/utility"
)

func ParseCards(input string) (cards map[int]int) {
	cards = make(map[int]int)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		numLists := strings.Split(line, " | ")
		if len(numLists) != 2 {
			fmt.Println("Not two lists of numbers???")
		}
		card, numbers, _ := strings.Cut(numLists[0], ":")

		var haveNumbers []string
		for _, num := range strings.Split(numLists[1], " ") {
			num := strings.TrimSpace(num)
			if len(num) > 0 {
				haveNumbers = append(haveNumbers, num)
			}
		}

		matches := 0
		for _, num := range strings.Split(numbers, " ") {
			num = strings.TrimSpace(num)
			if len(num) > 0 && slices.Contains(haveNumbers, num) {
				matches++
			}
		}
		cardNumber, _ := utility.StringToInt(card[len("Card "):])
		cards[cardNumber] = matches
	}

	return
}

func SolvePart1(cards map[int]int) (result int) {
	for _, matches := range cards {
		value := 0
		if matches > 0 {
			value++
			if matches > 1 {
				for i := 1; i < matches; i++ {
					value *= 2
				}
			}
		}
		result += value
	}

	return
}

// Recursively modifies a list of card numbers in place
// The solution is the length of the resulting list
func Part2ProcessR(cards map[int]int, currentCard int, result *[]int) {
	*result = append(*result, currentCard)
	for i := currentCard + 1; i <= min(currentCard+cards[currentCard], len(cards)); i++ {
		Part2ProcessR(cards, i, result)
	}
}

// Recursively counts the total number of cards
// (because the solution doesn't actually care what cards they are)
func Part2ProcessR2(cards map[int]int, currentCard int) (result int) {
	result++
	for i := currentCard + 1; i <= min(currentCard+cards[currentCard], len(cards)); i++ {
		result += Part2ProcessR2(cards, i)
	}
	return
}

// Inserts and counts cards inserted into a queue, until the queue is empty
func Part2ProcessQueue(cards map[int]int, currentCard int) (count int) {
	queue := []int{currentCard}
	for len(queue) > 0 {
		currentCard := queue[len(queue)-1]
		count++
		if len(queue) > 0 {
			queue = queue[:len(queue)-1]
		}
		for i := currentCard + 1; i <= min(currentCard+cards[currentCard], len(cards)); i++ {
			queue = append(queue, i)
		}
	}

	return
}

func SolvePart2(cards map[int]int) (result int) {
	for c, _ := range cards {
		result += Part2ProcessR2(cards, c)
	}
	return
}

func main() {
	buffer, _ := os.ReadFile("data/day4_input.txt")
	str := strings.TrimSpace(string(buffer))
	cards := ParseCards(str)

	part1Result := SolvePart1(cards)
	part2Result := SolvePart2(cards)

	fmt.Println("Part 1:", part1Result)
	fmt.Println("Part 2:", part2Result)
}
