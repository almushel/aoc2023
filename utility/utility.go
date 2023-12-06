package utility

import (
	"cmp"
	"errors"
	"strings"
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
	return RunesToInt([]rune(strings.TrimSpace(str))...)
}
