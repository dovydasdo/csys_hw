package cmd

import (
	"fmt"
	"hwone/pkg"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hwone",
	Short: "Homework One For Cryptographic Systems",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		encOpts := pkg.GetEncryptorOptions(
			pkg.EncWithCapability(
				pkg.Capability{
					Name:   "a1z26",
					CapF:   pkg.A1z26,
					Inputs: []any{"DovydasDomarkas"},
				},
			),
			pkg.EncWithCapability(
				pkg.Capability{
					Name:   "ascii",
					CapF:   pkg.Ascii,
					Inputs: []any{"DovydasDomarkas"},
				},
			),
			pkg.EncWithCapability(
				pkg.Capability{
					Name:   "binary",
					CapF:   pkg.Binary,
					Inputs: []any{20172561},
				},
			),
			pkg.EncWithCapability(
				pkg.Capability{
					Name:   "octal",
					CapF:   pkg.Octal,
					Inputs: []any{20172561},
				},
			),
			pkg.EncWithCapability(
				pkg.Capability{
					Name:   "hexadecimal",
					CapF:   pkg.Hexadecimal,
					Inputs: []any{20172561},
				},
			),
			pkg.EncWithCapability(
				pkg.Capability{
					Name:   "base64",
					CapF:   pkg.Base64,
					Inputs: []any{20172561},
				},
			),
			pkg.EncWithCapability(
				pkg.Capability{
					Name:   "base58",
					CapF:   pkg.Base58,
					Inputs: []any{20172561},
				},
			),
			pkg.EncWithCapability(
				pkg.Capability{
					Name:   "caesar",
					CapF:   pkg.Caesar,
					Inputs: []any{[]any{20172561 % 26, "DOVYDASDOMARKAS"}},
				},
			),
			pkg.EncWithCapability(
				pkg.Capability{
					Name:   "vigenere",
					CapF:   pkg.Vigenere,
					Inputs: []any{[]any{"DOVYDASDOMARKAS", "VILNIUSGEDIMINASTECHNIKALUNIVERSITY"}},
				},
			),
		)
		enc := pkg.GetEncryptor(encOpts)

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
