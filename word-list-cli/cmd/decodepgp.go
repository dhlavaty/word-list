package cmd

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"word-list-cli/cmd/pgp"
)

// decodePgpCmd represents the decode PGP command
var (
	decodePgpCmd = &cobra.Command{
		Use:   "decodepgp",
		Short: "Decode PGP word list back to binary data",
		Long: `Decode mnemonic PGP word list back to binary data.

For PGP word list spec see:
https://en.wikipedia.org/wiki/PGP_word_list
`,
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)
			scanner := bufio.NewScanner(reader)
			scanner.Split(bufio.ScanWords)
			decoder := pgp.NewDecoder()

			for scanner.Scan() {
				word := scanner.Text()

				if strings.Contains(word, "-") || strings.Contains(word, "PGP") || strings.Contains(word, "WORD") {
					// skip ASCII ARMOR header and footer
					continue
				}

				b, err := decoder.GetByte(word)
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
		},
	}
)

func init() {
	rootCmd.AddCommand(decodePgpCmd)
}
