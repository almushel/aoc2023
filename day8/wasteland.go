package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func ParseInput(input string) (directions string, network map[string][2]string) {
	network = make(map[string][2]string)
	directions, input, found := strings.Cut(input, "\n\n")
	if !found {
		return
	}
	for _, line := range strings.Split(input, "\n") {
		node := line[:3]
		edges := strings.Split(line[len("XXX = ("):len(line)-1], ", ")
		network[node] = [2]string{edges[0], edges[1]}
	}
	return
}

func SolvePart1(input string) (result int) {
	directions, network := ParseInput(input)

	found := false
	currentNode := "AAA"
	for !found {
		for _, dir := range directions {
			result++
			if dir == 'L' {
				currentNode = network[currentNode][0]
			} else {
				currentNode = network[currentNode][1]
			}
			if currentNode == "ZZZ" {
				found = true
			}
		}
	}

	return
}

func getPrimeFactors(num int) []int {
	var result []int

	for num%2 == 0 {
		result = append(result, 2)
		num /= 2
	}

	for i := 3; i < num; i += 2 {
		for num%i == 0 {
			result = append(result, i)
			num /= i
		}
	}

	if num > 2 {
		result = append(result, num)
	}

	return result
}

func getUniqueFactors(nums []int) (uniqueFactors []int) {
	factors := make([][]int, len(nums))
	maxFactors := 0
	for i, _ := range factors {
		factors[i] = getPrimeFactors(nums[i])
		length := len(factors[i])
		if length > maxFactors {
			maxFactors = length
		}
	}

	for i := 0; i < maxFactors; i++ {
		var found []int
		for f, _ := range factors {
			if i >= len(factors[f]) {
				continue
			}
			checkFactor := factors[f][i]
			if !slices.Contains(found, checkFactor) {
				found = append(found, checkFactor)
			}
		}
		uniqueFactors = append(uniqueFactors, found...)
	}
	return
}

func SolvePart2(input string) (result int) {
	directions, network := ParseInput(input)

	var entrances []string
	for node, _ := range network {
		if node[2] == 'A' {
			entrances = append(entrances, node)
		}
	}
	var paths []int
	for _, entrance := range entrances {
		currentNode := entrance
		count := 0
		found := false
		for {
			for _, dir := range directions {
				count++
				if dir == 'L' {
					currentNode = network[currentNode][0]
				} else {
					currentNode = network[currentNode][1]
				}

				if currentNode[2] == 'Z' {
					paths = append(paths, count)
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}
	result = 1
	factors := getUniqueFactors(paths)
	for _, factor := range factors {
		result *= factor
	}
	return
}

func main() {
	buffer, err := os.ReadFile("data/day8_input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	str := strings.TrimSpace(string(buffer))

	part1Result := SolvePart1(str)
	part2Result := SolvePart2(str)

	fmt.Println(part1Result)
	fmt.Println(part2Result)
	return
}
