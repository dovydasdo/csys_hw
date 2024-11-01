package hw1

import (
	"errors"
	"fmt"
	"strings"
)

const caesarAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Caesar(input any) (string, error) {
	out := ""
	// 0 = key, 1 = input
	if in, ok := input.([]any); ok {
		if len(in) != 2 {
			return out, errors.New("expected two input parameters")
		}
		key, ok := in[0].(int)
		if !ok {
			return out, errors.New("key should be an intager")
		}
		val, ok := in[1].(string)
		if !ok {
			return out, errors.New("key should be an intager")
		}

		for i := 0; i < len(val); i++ {
			c := val[i]
			index := strings.Index(caesarAlphabet, string(c))
			if index == -1 {
				fmt.Printf("char %v not found in alphabet", c)
				return "", errors.New(fmt.Sprintf("char %v not found in alphabet", c))
			}
			toReplaceIndex := (index + key) % len(caesarAlphabet)
			out += string(caesarAlphabet[toReplaceIndex])
		}
	} else {
		return out, errors.New("expected an array as input")
	}

	return out, nil
}
