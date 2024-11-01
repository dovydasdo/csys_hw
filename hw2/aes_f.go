package hw2

import (
	"crypto/aes"
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
		out += fmt.Sprintf("len of plaintext in bytes: %v, bytes: %v \n", len(plainTextBlock), plainTextBlock)

		keyBlock := ToBlockOfSize(in.KeySource, in.PlainTextBytesCount)
		out += fmt.Sprintf("len of key in bytes: %v, bytes: %v \n", len(keyBlock), keyBlock)

		cypher, err := aes.NewCipher(keyBlock)
		if err != nil {
			fmt.Println(err)
			return "", err
		}

		outBytes := make([]byte, in.PlainTextBytesCount)

		cypher.Encrypt(outBytes, plainTextBlock)
		out += fmt.Sprintf("encrypted bytes length: %v, bytes: %v \n", len(outBytes), outBytes)
		out += fmt.Sprintf("encrypted bytes as binary: %v \n", getBytesAsBinary(outBytes))

		decBytes := make([]byte, in.PlainTextBytesCount)
		cypher.Decrypt(decBytes, outBytes)
		out += fmt.Sprintf("decrypted bytes length: %v, bytes: %v string: %v \n", len(decBytes), decBytes, string(decBytes))

		// Flip last bit map
		flipMask := 0b00000001

		fPlainTextByte := plainTextBlock[in.PlainTextBytesCount-1]
		fPlainTextByte ^= byte(flipMask)
		fPlainTextBytes := append(plainTextBlock[:in.PlainTextBytesCount-1], fPlainTextByte)
		outBytesFB := make([]byte, in.PlainTextBytesCount)
		cypher.Encrypt(outBytesFB, fPlainTextBytes)
		diffPT := countBitDifferences(outBytes, outBytesFB)
		out += fmt.Sprintf("encrypted bytes as binary (flipped bit in plaintext): %v \n", getBytesAsBinary(outBytesFB))
		out += fmt.Sprintf("diff from flipping plain text: %v \n", diffPT)

		fkPlainTextByte := keyBlock[in.PlainTextBytesCount-1]
		fkPlainTextByte ^= byte(flipMask)
		fKeyBytes := append(keyBlock[:in.PlainTextBytesCount-1], fkPlainTextByte)
		outBytesFBK := make([]byte, in.PlainTextBytesCount)
		cypher.Encrypt(outBytesFBK, fKeyBytes)
		diffK := countBitDifferences(outBytes, outBytesFBK)
		out += fmt.Sprintf("encrypted bytes as binary (flipped bit in key): %v \n", getBytesAsBinary(outBytesFB))
		out += fmt.Sprintf("diff from flipping key: %v \n", diffK)
	} else {
		return "", errors.New("expected DESArgs as input")
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

func getBytesAsBinary(b []byte) string {
	result := ""
	for _, n := range b {
		result += fmt.Sprintf("%08b ", n)
	}

	return result
}
