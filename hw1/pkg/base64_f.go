package pkg

import (
	"errors"
	"math/big"
)

const base64Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func Base64(input any) (string, error) {
	out := ""
	if in, ok := input.(int); ok {
		if in == 0 {
			return string(base58Alphabet[0]), nil
		}
		bigIntValue := big.NewInt(int64(in))
		base := big.NewInt(64)
		zero := big.NewInt(0)
		var outBytes []byte
		for bigIntValue.Cmp(zero) != 0 {
			mod := new(big.Int)
			bigIntValue.DivMod(bigIntValue, base, mod)
			outBytes = append([]byte{base58Alphabet[mod.Int64()]}, outBytes...)
		}
		out = string(outBytes[:])
	} else {
		return out, errors.New("expected int as input")
	}

	return out, nil
}