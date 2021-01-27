package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"word-list-cli/cmd/bip39"
)

// bip39Cmd represents the encode to BIP39 command
var (
	bip39Cmd = &cobra.Command{
		Use:   "bip39",
		Short: "Encode binary data to BIP39 world list",
		Long: `Encode binary data to mnemonic BIP39 world list. Output is
100% compatible with BIP39 deterministic keys. Unlike original
BIP39 algorithm, used for deterministic crypto currency keys,
this tool supports encoding any binary data of any length.

For original BIP39 spec see:
https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki
`,
		Run: func(cmd *cobra.Command, args []string) {

			reader := bufio.NewReader(os.Stdin)
			encoder := bip39.NewEncoder()
			if flagNoArmor == false {
				fmt.Println("-----BEGIN BIP39-----")
			}
			for {
				mybyte, err := reader.ReadByte()
				if err != nil {
					break // EOF
				}
				word, ok := encoder.GetWord(mybyte)
				if ok {
					fmt.Print(word)
					fmt.Print(" ")
				}
			}
			word, ok := encoder.GetLastWord()
			if ok {
				fmt.Println(word)
			}
			if flagNoArmor == false {
				fmt.Println("-----END BIP39-----")
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(bip39Cmd)
}
