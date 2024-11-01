package hw1

import (
	"crypto/des"
	"errors"
	"fmt"
)

type AESArgs struct {
	PlainTextSource     string
	KeySource           string
	PlainTextBytesCount int
}

func AES(input any) (string, error) {
	out := ""
	if in, ok := input.(AESArgs); ok {
		plainTextBlock := ToBlockOfSize(in.PlainTextSource, in.PlainTextBytesCount)
		fmt.Printf("len of plaintext in bytes: %v, bytes: %v \n", len(plainTextBlock), plainTextBlock)

		keyBlock := ToBlockOfSize(in.KeySource, in.PlainTextBytesCount)
		fmt.Printf("len of key in bytes: %v, bytes: %v \n", len(keyBlock), keyBlock)

		cypher, err := des.NewCipher(keyBlock)
		if err != nil {
			return "", err
		}

		outBytes := make([]byte, in.PlainTextBytesCount)

		cypher.Encrypt(outBytes, plainTextBlock)
		fmt.Printf("encrypted bytes length: %v, bytes: %v \n", len(outBytes), outBytes)
		printBytesAsBinary(outBytes)

		decBytes := make([]byte, in.PlainTextBytesCount)
		cypher.Decrypt(decBytes, outBytes)
		fmt.Printf("decrypted bytes length: %v, bytes: %v string: %v \n", len(decBytes), decBytes, string(decBytes))

		// Flip last bit map
		flipMask := 0b00000001

		fPlainTextByte := plainTextBlock[in.PlainTextBytesCount-1]
		fPlainTextByte ^= byte(flipMask)
		fPlainTextBytes := append(plainTextBlock[:in.PlainTextBytesCount-1], fPlainTextByte)
		outBytesFB := make([]byte, in.PlainTextBytesCount)
		cypher.Encrypt(outBytesFB, fPlainTextBytes)
		diffPT := countBitDifferences(outBytes, outBytesFB)
		printBytesAsBinary(outBytesFB)
		fmt.Printf("diff from flipping plain text: %v \n", diffPT)

		fkPlainTextByte := keyBlock[in.PlainTextBytesCount-1]
		fkPlainTextByte ^= byte(flipMask)
		fKeyBytes := append(keyBlock[:in.PlainTextBytesCount-1], fkPlainTextByte)
		outBytesFBK := make([]byte, in.PlainTextBytesCount)
		cypher.Encrypt(outBytesFBK, fKeyBytes)
		diffK := countBitDifferences(outBytes, outBytesFBK)
		printBytesAsBinary(outBytesFBK)
		fmt.Printf("diff from flipping key: %v \n", diffK)
	} else {
		return "", errors.New("expected AESArgs as input")
	}

	return out, nil
}

func ToBlockOfSize(source string, size int) []byte {
	block := []byte(source)
	if len(block) > size {
		block = block[:size]
	}

	if len(block) < size {
		diff := size - len(block)

		for i := 0; i < diff; i++ {
			block = append(block, byte(0))
		}
	}

	return block
}

func countBitDifferences(b1, b2 []byte) int {
	totalDifferences := 0
	for i := 0; i < len(b1); i++ {
		diff := b1[i] ^ b2[i]
		for diff != 0 {
			totalDifferences += int(diff & 1)
			diff >>= 1
		}
	}

	return totalDifferences
}

func printBytesAsBinary(b []byte) {
	for _, n := range b {
		fmt.Printf("%08b ", n)
	}
}
