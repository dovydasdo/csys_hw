package hw1

import (
	"errors"
	"fmt"
	"strings"
)

const vigenereAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Vigenere(input any) (string, error) {
	out := ""
	// 0 = key, 1 = input
	if in, ok := input.([]any); ok {
		if len(in) != 2 {
			return out, errors.New("expected two input parameters")
		}
		key, ok := in[0].(string)
		if !ok {
			return out, errors.New("key should be a string")
		}
		val, ok := in[1].(string)
		if !ok {
			return out, errors.New("value should be a string")
		}
		fmt.Sprintf("key: %v \n", key)
		key = makeKeyOfLen(key, len(val))
		fmt.Sprintf("keyL: %v \n val: %v \n", key, val)

		for i := 0; i < len(val); i++ {
			c := val[i]
			indexC := strings.Index(caesarAlphabet, string(c))
			if indexC == -1 {
				fmt.Printf("char %v not found in alphabet", c)
				return "", errors.New(fmt.Sprintf("char %v not found in alphabet", c))
			}

			k := key[i]
			indexK := strings.Index(caesarAlphabet, string(k))
			if indexK == -1 {
				fmt.Printf("char %v not found in alphabet", k)
				return "", errors.New(fmt.Sprintf("char %v not found in alphabet", k))
			}

			toReplaceIndex := (indexC + indexK) % len(caesarAlphabet)
			out += string(caesarAlphabet[toReplaceIndex])

		}
	} else {
		return out, errors.New("expected an array as input")
	}

	return out, nil
}

func makeKeyOfLen(key string, length int) string {
	if len(key) >= length {
		return key[:length]
	}
	keyL := key
	// add the diff of key vs val
	for i := 0; i < length-len(key); i++ {
		keyL += string(key[i%len(key)])
	}
	return keyL
}
