package cmd

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"word-list-cli/cmd/bip39"
)

// decodeBip39Cmd represents the decode BIP39 command
var (
	decodeBip39Cmd = &cobra.Command{
		Use:   "decodebip39",
		Short: "Decode BIP39 word list back to binary data",
		Long: `Decode mnemonic BIP39 word list back to binary data.

For original BIP39 spec see:
https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki
`,
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)
			scanner := bufio.NewScanner(reader)
			scanner.Split(bufio.ScanWords)
			decoder := bip39.NewDecoder()

			for scanner.Scan() {
				word := scanner.Text()

				if strings.Contains(word, "-") {
					// skip ASCII ARMOR header and footer
					continue
				}

				b, err := decoder.GetBytes(word)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Unknown word '%v' found. Cannot parse. Exiting\n", word)
					return
				}
				binary.Write(os.Stdout, binary.LittleEndian, b)
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "error reading standard input:", err)
				return
			}

			_, err := decoder.CheckHash()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Hash problem: '%v'. Exiting\n", err)
				return
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(decodeBip39Cmd)
}
