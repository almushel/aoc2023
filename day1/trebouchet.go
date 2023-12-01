// Day 1, Puzzle 1: Trebouchet?!
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func runesToInt(runes ...rune) (value int, err error) {
	for i, r := range runes {
		if r < '0' || r > '9' {
			return 0, errors.New("Non-numeric rune " + string(r))
		}
		digit := int(r - '0')
		for t := 1; t < (len(runes) - i); t++ {
			digit *= 10
		}
		value += digit
	}
	return value, nil
}

func stringToInt(str string) (int, error) {
	return runesToInt([]rune(str)...)
}

func main() {
	var digitWords []string = []string{
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}

	buffer, err := os.ReadFile("data/day1_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileStr := string(buffer)

	var first, last rune
	var sum int = 0

	for i, c := range fileStr {
		found := false
		for num, word := range digitWords {
			if strings.IndexRune(word, c) != 0 {
				continue
			}

			cword := fileStr[i:min(i+len(word), len(fileStr))]
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
			val, _ := runesToInt(first, last)
			sum += int(val)
			first, last = '\000', '\000'
		}
	}
	val, _ := runesToInt(first, last)
	sum += val

	fmt.Println(sum)
}
