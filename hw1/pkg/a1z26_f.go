package pkg

import (
	"errors"
	"strconv"
)

func A1z26(input any) (string, error) {
	// lower case: 49-74
	// upper case: 17-42

	out := ""
	if in, ok := input.(string); ok {
		for i, char := range in {
			intChar := int(char - '0')
			if intChar > 48 && intChar < 75 {
				intChar -= 48
			} else if intChar > 16 && intChar < 43 {
				intChar -= 16
			} else {
				out += string(char)
				continue
			}

			out += strconv.Itoa(intChar)
			if i != len(in)-1 {
				out += "-"
			}
		}

	} else {
		return "", errors.New("expected string as input")
	}

	return out, nil
}
