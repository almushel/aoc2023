package boatrace

import (
	"fmt"
	"strings"

	"github.com/almushel/aoc2023/utility"
)

func parseInput1(input string) (times []int, distances []int) {
	t, d, found := strings.Cut(input, "\n")
	if !found {
		return
	}

	for _, time := range strings.Split(t[len("Time:"):], " ") {
		time = strings.TrimSpace(time)
		if len(time) == 0 {
			continue
		}
		it, err := utility.StringToInt(time)
		if err != nil {
			fmt.Println(err.Error())
		}
		times = append(times, it)
	}

	for _, dist := range strings.Split(d[len("Distance:"):], " ") {
		dist = strings.TrimSpace(dist)
		if len(dist) == 0 {
			continue
		}
		id, err := utility.StringToInt(dist)
		if err != nil {
			fmt.Println(err.Error())
		}
		distances = append(distances, id)
	}

	return
}

func parseInput2(input string) (time int, distance int) {
	t, d, found := strings.Cut(input, "\n")
	if !found {
		return
	}

	ts := ""
	for _, tChar := range t[len("Time:"):] {
		if utility.IsNumeric(tChar) {
			ts += string(tChar)
		}
	}
	time, err := utility.StringToInt(ts)
	if err != nil {
		fmt.Println(err.Error())
	}

	ds := ""
	for _, dChar := range d[len("Distance:"):] {
		if utility.IsNumeric(dChar) {
			ds += string(dChar)
		}
	}
	distance, err = utility.StringToInt(ds)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

func victoryConditions(time, target int) int {
	minimum, maximum := 0, time
	for (time-maximum)*maximum < target {
		maximum--
	}
	for (time-minimum)*minimum < target {
		minimum++
	}
	return maximum - minimum + 1
}

func SolvePart1(input string) (result int) {
	result = 1
	times, distances := parseInput1(input)
	for i := 0; i < len(times); i++ {
		result *= victoryConditions(times[i], distances[i]+1)
	}
	return
}

func SolvePart2(input string) (result int) {
	time, distance := parseInput2(input)
	result = victoryConditions(time, distance+1)

	return
}
