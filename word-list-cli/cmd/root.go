package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var flagNoArmor bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "word-list-cli",
	Version: "0.0.1",
	Short:   "word-list-cli is utility for encoding and decoding binary data to different mnemonic world lists",
	Long: `Utility for encoding and decoding binary data to different
mnemonic world lists. Currently only BIP39 is supported.

Utility supports encoding data of any length, it is rather
practical for storing a small amount of data. For example
to encode passwords, passphrases or cryptographic keys and
print then on paper.

Source code: TODO
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&flagNoArmor, "no-armor", false, "turns off ASCII Armored output")
}
