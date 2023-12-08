package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/almushel/aoc2023/utility"
)

const (
	fiveOfAKind  = 7
	fourOfAKind  = 6
	fullHouse    = 5
	threeOfAKind = 4
	twoPair      = 3
	onePair      = 2
	highCard     = 1
)

type HB struct {
	hand string
	bid  int
}

func parseInput(input string) (hands []HB) {
	for _, line := range strings.Split(input, "\n") {
		hand, bidStr, _ := strings.Cut(line, " ")

		var newHB HB
		newHB.hand = hand
		newHB.bid, _ = utility.StringToInt(bidStr)
		hands = append(hands, newHB)
	}
	return
}

func getCardValue(card rune) int {
	if card >= '2' && card <= '9' {
		return int(card - '1')
	}
	nine := int('9' - '1')
	switch card {
	case 'T':
		return nine + 1
	case 'J':
		return nine + 2
	case 'Q':
		return nine + 3
	case 'K':
		return nine + 4
	case 'A':
		return nine + 5
	}

	return 0
}

func getCardValue2(card rune) int {
	if card == 'J' {
		return 0
	}
	return getCardValue(card)
}

func getHandValue(hand string, jokers bool) (handValue int) {
	cards := make(map[rune]int)
	highestCount := 0
	for _, card := range hand {
		_, ok := cards[card]
		if !ok {
			cards[card] = 0
		}
		cards[card]++
		if cards[card] > highestCount {
			if jokers && card == 'J' {
				continue
			}
			highestCount = cards[card]
		}
	}

	uniqueCards := len(cards)

	if jokers {
		jCount, ok := cards['J']
		if ok && uniqueCards > 1 {
			uniqueCards--
			highestCount += jCount
		}
	}

	switch uniqueCards {
	case 1:
		handValue = fiveOfAKind
	case 2:
		if highestCount == 4 {
			handValue = fourOfAKind
		} else {
			handValue = fullHouse
		}
	case 3:
		if highestCount == 3 {
			handValue = threeOfAKind
		} else {
			handValue = twoPair
		}
	case 4:
		handValue = onePair
	default:
		handValue = highCard
	}

	return
}

func getHandValue1(hand string) int {
	return getHandValue(hand, false)
}

func getHandValue2(hand string) int {
	return getHandValue(hand, true)
}

func buildSortFunc(handVal func(string) int, cardVal func(rune) int) func(HB, HB) int {
	result := func(a, b HB) int {
		ahv := handVal(a.hand)
		bhv := handVal(b.hand)

		if ahv == bhv {
			for i, aChar := range a.hand {
				bChar := rune(b.hand[i])
				av := cardVal(aChar)
				bv := cardVal(bChar)
				if av > bv {
					return 1
				} else if bv > av {
					return -1
				}
			}
			return 0
		} else if ahv > bhv {
			return 1
		}
		return -1
	}

	return result
}

func SolvePart1(input string) (result int) {
	hands := parseInput(input)
	compare := buildSortFunc(getHandValue1, getCardValue)
	slices.SortFunc(hands, compare)
	for i, hand := range hands {
		result += (i + 1) * hand.bid
	}
	return
}

func SolvePart2(input string) (result int) {
	hands := parseInput(input)
	compare := buildSortFunc(getHandValue2, getCardValue2)
	slices.SortFunc(hands, compare)
	for i, hand := range hands {
		result += (i + 1) * hand.bid
	}
	return
}

func main() {
	buffer, err := os.ReadFile("data/day7_input.txt")
	if err != nil {
		println(err.Error())
		return
	}
	str := strings.TrimSpace(string(buffer))
	part1Result := SolvePart1(str)
	part2Result := SolvePart2(str)

	fmt.Println("Part 1:", part1Result)
	fmt.Println("Part 2:", part2Result)
}
