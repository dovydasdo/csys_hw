package hw2

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"math/rand"

	"golang.org/x/crypto/sha3"
)

type SHAArgs struct {
	PlainTextSource string
	MaxSearch       int
}

func SHA(input any) (string, error) {
	out := ""
	if in, ok := input.(SHAArgs); ok {
		h := sha256.New()
		h.Write([]byte(in.PlainTextSource))
		sha256Hash := h.Sum(nil)
		out += fmt.Sprintf("%v SHA256 hash: %x \n", in.PlainTextSource, sha256Hash)

		h3 := sha3.New256()
		h3.Write([]byte(in.PlainTextSource))
		// SHA256 hash
		sha3Hash := h3.Sum(nil)
		out += fmt.Sprintf("%v SHA3 hash: %x \n", in.PlainTextSource, sha3Hash)

		// Binary representation of SHA256 hash
		binSHA256 := getBytesAsBinary(sha256Hash)
		out += fmt.Sprintf("binary representation of (SHA256 hash) %x: %v \n", sha256Hash, binSHA256)

		// Last octal value of SHA256 hash
		octSHA256 := binSHA256[len(binSHA256)-4:]
		out += fmt.Sprintf("last octal value (last three bits) of binary representation of (SHA256 hash) %x: %v \n", sha256Hash, octSHA256)

		// Value who's hash has the same last octal value of SHA256 hash
		sameOctSHA256 := GetSHA256HashWithOctal(octSHA256, in.MaxSearch)
		out += fmt.Sprintf("Value with same hash octal value (SHA256): %v\n", sameOctSHA256)

		// Value who's hash has the same last byte value of SHA256 hash
		fmt.Println("test")
		fmt.Println(sha256Hash[len(sha256Hash)-1])
		sameByteSHA256 := GetSHA256HashWithByte(sha256Hash[len(sha256Hash)-1], in.MaxSearch)
		out += fmt.Sprintf("Value with same hash byte value (SHA256): %v\n", sameByteSHA256)

		// Value who's hash has the same last byte value of SHA256 hash
		sameByteSHA256 = GetSHA256HashWithByte(sha256Hash[len(sha256Hash)-1], in.MaxSearch)
		out += fmt.Sprintf("Value with same hash byte value (SHA256): %v\n", sameByteSHA256)

		// SHA3 hash
		binSHA3 := getBytesAsBinary(sha3Hash)
		out += fmt.Sprintf("binary representation of (SHA3 hash) %x: %v \n", sha3Hash, binSHA3)

		// Last octal value of SHA3 hash
		octSHA3 := binSHA3[len(binSHA3)-4:]
		out += fmt.Sprintf("last octal value (last three bits) of binary representation of (SHA3 hash) %x: %v \n", sha3Hash, octSHA3)

		// Value who's hash has the same last octal value as SHA3 hash
		sameOctSHA3 := GetSHA3HashWithOctal(octSHA3, in.MaxSearch)
		out += fmt.Sprintf("Value with same hash last octal value (SHA3): %v\n", sameOctSHA3)

		// Value who's hash has the same last byte value as SHA3 hash
		sameByteSHA3 := GetSHA3HashWithByte(sha3Hash[len(sha256Hash)-1], in.MaxSearch)
		out += fmt.Sprintf("Value with same hash byte value (SHA3): %v\n", sameByteSHA3)

		// Value who's hash has the same last byte value as SHA3 hash
		sameByteSHA3 = GetSHA3HashWithByte(sha3Hash[len(sha256Hash)-1], in.MaxSearch)
		out += fmt.Sprintf("Value with same hash byte value (SHA3): %v\n", sameByteSHA3)

	} else {
		return "", errors.New("expected SHAArgs as input")
	}

	return out, nil
}

func GetSHA256HashWithOctal(octal string, max int) string {
	out := ""
	i := 0

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	for {
		b := make([]rune, rand.Intn(100))
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}

		input := string(b)
		h := sha256.New()
		h.Write([]byte(input))
		sha256Hash := h.Sum(nil)
		binSHA256 := getBytesAsBinary(sha256Hash)
		if octal == binSHA256[len(binSHA256)-4:] {
			out = fmt.Sprintf("text val: %v, hash: %v \nhash binary: %v \n", input, sha256Hash, getBytesAsBinary(sha256Hash))
			break
		}

		i++
		if i > max {
			break
		}
	}

	return out

}

func GetSHA256HashWithByte(bt byte, max int) string {
	out := ""
	i := 0

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	for {
		b := make([]rune, rand.Intn(100))
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}

		input := string(b)
		h := sha256.New()
		h.Write([]byte(input))
		sha256Hash := h.Sum(nil)
		if bt == sha256Hash[len(sha256Hash)-1] {
			out = fmt.Sprintf("text val: %v, hash: %v \nhash binary: %v \n", input, sha256Hash, getBytesAsBinary(sha256Hash))
			break
		}

		i++
		if i > max {
			break
		}
	}

	return out

}

func GetSHA3HashWithOctal(octal string, max int) string {
	out := ""
	i := 0

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	for {
		b := make([]rune, rand.Intn(100))
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}

		input := string(b)
		h := sha3.New256()
		h.Write([]byte(input))
		sha256Hash := h.Sum(nil)
		binSHA256 := getBytesAsBinary(sha256Hash)
		if octal == binSHA256[len(binSHA256)-4:] {
			out = fmt.Sprintf("text val: %v, hash: %v \nhash binary: %v \n", input, sha256Hash, getBytesAsBinary(sha256Hash))
			break
		}

		i++
		if i > max {
			break
		}
	}

	return out

}

func GetSHA3HashWithByte(bt byte, max int) string {
	out := ""
	i := 0

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	for {
		b := make([]rune, rand.Intn(100))
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}

		input := string(b)
		h := sha3.New256()
		h.Write([]byte(input))
		sha256Hash := h.Sum(nil)
		if bt == sha256Hash[len(sha256Hash)-1] {
			out = fmt.Sprintf("text val: %v, hash: %v \nhash binary: %v \n", input, sha256Hash, getBytesAsBinary(sha256Hash))
			break
		}

		i++
		if i > max {
			break
		}
	}

	return out

}
