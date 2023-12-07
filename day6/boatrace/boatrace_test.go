package boatrace

import (
	"testing"
)

const (
	exampleSolution1 = 288
	exampleSolution2 = 71503
	part1Solution    = 281600
	part2Solution    = 33875953
)

var exampleInput string = `Time:      7  15   30
Distance:  9  40  200`

var input string = `Time:        47     70     75     66
Distance:   282   1079   1147   1062`

func TestPart1Example(t *testing.T) {
	want := exampleSolution1
	result := SolvePart1(exampleInput)
	if result != want {
		t.Fatalf("result = %d, want: %d),", result, want)
	}
}

func TestPart2Example(t *testing.T) {
	want := exampleSolution2
	result := SolvePart2(exampleInput)
	if result != want {
		t.Fatalf("result = %d, want: %d),", result, want)
	}
}

func TestPart1(t *testing.T) {
	want := part1Solution
	result := SolvePart1(input)
	if result != want {
		t.Fatalf("result = %d, want: %d),", result, want)
	}
}

func TestPart2(t *testing.T) {
	want := part2Solution
	result := SolvePart2(input)
	if result != want {
		t.Fatalf("result = %d, want: %d),", result, want)
	}
}
