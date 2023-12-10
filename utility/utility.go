package utility

import (
	"cmp"
	"errors"
	"strings"
	"testing"
)

const (
	UINT_MAX = ^uint(0)
	INT_MAX  = int(UINT_MAX >> 1)
)

func Clamp[T cmp.Ordered](val T, low T, high T) T {
	return min(max(val, low), high)
}

func IsNumeric(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

func RunesToInt(runes ...rune) (value int, err error) {
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

func StringToInt(str string) (int, error) {
	sign := 1
	if str[0] == '-' {
		sign = -1
		str = str[1:]
	}
	result, err := RunesToInt([]rune(strings.TrimSpace(str))...)
	return sign * result, err
}

func PuzzleTest(t *testing.T, input string, want int, solve func(string) int) {
	result := solve(input)
	if result != want {
		t.Fatalf("result = %d, want: %d),", result, want)
	}
}
