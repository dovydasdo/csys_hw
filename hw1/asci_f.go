package hw1

import (
	"errors"
	"strconv"
)

func Ascii(input any) (string, error) {
	out := ""
	if in, ok := input.(string); ok {
		byteArray := []byte(in)
		for i, byte := range byteArray {
			out += strconv.Itoa(int(byte))
			if i != len(byteArray) {
				out += " "
			}
		}

	} else {
		return "", errors.New("expected string as input")
	}

	return out, nil
}
