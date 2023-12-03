package utility

import "errors"

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
	return RunesToInt([]rune(str)...)
}
