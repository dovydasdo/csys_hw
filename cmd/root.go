package cmd

import (
	hwone "csyshw/hw1"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hw",
	Short: "Homework One For Cryptographic Systems",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: parse somek YAML to configure
		encOpts := hwone.GetEncryptorOptions(
			hwone.EncWithCapability(
				hwone.Capability{
					Name:   "a1z26",
					CapF:   hwone.A1z26,
					Inputs: []any{"DovydasDomarkas"},
				},
			),
			hwone.EncWithCapability(
				hwone.Capability{
					Name:   "ascii",
					CapF:   hwone.Ascii,
					Inputs: []any{"DovydasDomarkas"},
				},
			),
			hwone.EncWithCapability(
				hwone.Capability{
					Name:   "binary",
					CapF:   hwone.Binary,
					Inputs: []any{20172561},
				},
			),
			hwone.EncWithCapability(
				hwone.Capability{
					Name:   "octal",
					CapF:   hwone.Octal,
					Inputs: []any{20172561},
				},
			),
			hwone.EncWithCapability(
				hwone.Capability{
					Name:   "hexadecimal",
					CapF:   hwone.Hexadecimal,
					Inputs: []any{20172561},
				},
			),
			hwone.EncWithCapability(
				hwone.Capability{
					Name:   "base64",
					CapF:   hwone.Base64,
					Inputs: []any{20172561},
				},
			),
			hwone.EncWithCapability(
				hwone.Capability{
					Name:   "base58",
					CapF:   hwone.Base58,
					Inputs: []any{20172561},
				},
			),
			hwone.EncWithCapability(
				hwone.Capability{
					Name:   "caesar",
					CapF:   hwone.Caesar,
					Inputs: []any{[]any{20172561 % 26, "DOVYDASDOMARKAS"}},
				},
			),
			hwone.EncWithCapability(
				hwone.Capability{
					Name:   "vigenere",
					CapF:   hwone.Vigenere,
					Inputs: []any{[]any{"DOVYDASDOMARKAS", "VILNIUSGEDIMINASTECHNIKALUNIVERSITY"}},
				},
			),
			hwone.EncWithCapability(
				hwone.Capability{
					Name: "aes",
					CapF: hwone.AES,
					Inputs: []any{hwone.AESArgs{
						PlainTextSource:     "DovydasDomarkas",
						PlainTextBytesCount: 8,
						KeySource:           "20172561",
					}},
				},
			),
		)
		enc := hwone.GetEncryptor(encOpts)

		result, err := enc.Execute()
		if err != nil {
			panic(err)
		}

		fmt.Print(result)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
