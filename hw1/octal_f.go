package hw1

import (
	"errors"
	"strconv"
)

func Octal(input any) (string, error) {
	out := ""
	if in, ok := input.(int); ok {
		out = strconv.FormatInt(int64(in), 8)
	} else {
		return "", errors.New("expected int as input")
	}

	return out, nil
}
