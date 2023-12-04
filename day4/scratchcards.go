package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/almushel/aoc2023/utility"
)

type scratchcard struct {
	winningNumbers, haveNumbers []string
}

func parseCards(input string) (cards map[int]scratchcard) {
	cards = make(map[int]scratchcard)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		numLists := strings.Split(line, " | ")
		if len(numLists) != 2 {
			fmt.Println("Not two lists of numbers???")
		}
		card, numbers, _ := strings.Cut(numLists[0], ":")
		winningNumbers := strings.Split(numbers, " ")
		haveNumbers := strings.Split(numLists[1], " ")

		var newCard scratchcard
		for _, num := range winningNumbers {
			num := strings.TrimSpace(num)
			if len(num) == 0 {
				continue
			}
			newCard.winningNumbers = append(newCard.winningNumbers, num)
		}
		for _, num := range haveNumbers {
			num := strings.TrimSpace(num)
			if len(num) == 0 {
				continue
			}
			newCard.haveNumbers = append(newCard.haveNumbers, num)
		}
		cardNumber, _ := utility.StringToInt(card[len("Card "):])
		cards[cardNumber] = newCard
	}

	return
}

func solvePart1(cards map[int]scratchcard) (result int) {
	for _, card := range cards {
		cardValue := 0
		for _, num := range card.winningNumbers {
			if slices.Contains(card.haveNumbers, num) {
				if cardValue == 0 {
					cardValue = 1
				} else {
					cardValue *= 2
				}
			}
		}
		result += cardValue
	}

	return
}

func part2ProcessR(cards map[int]scratchcard, currentCard int) (result []int) {
	result = []int{currentCard}
	card := cards[currentCard]
	matches := 0
	for _, num := range card.winningNumbers {
		if slices.Contains(card.haveNumbers, num) {
			matches++
		}
	}
	for i := currentCard + 1; i <= min(currentCard+matches, len(cards)); i++ {
		result = append(result, part2ProcessR(cards, i)...)
	}
	return
}

func solvePart2(cards map[int]scratchcard) int {
	var newCards []int
	for c, _ := range cards {
		newCards = append(newCards, part2ProcessR(cards, c)...)
	}

	return len(newCards)
}

func main() {
	buffer, _ := os.ReadFile("data/day4_input.txt")
	str := strings.TrimSpace(string(buffer))
	cards := parseCards(str)

	part1Result := solvePart1(cards)
	part2Result := solvePart2(cards)

	fmt.Println("Part 1:", part1Result)
	fmt.Println("Part 2:", part2Result)
}
