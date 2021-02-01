package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"word-list-cli/cmd/pgp"
)

// pgpCmd represents the encode to PPG command
var (
	pgpCmd = &cobra.Command{
		Use:   "pgp",
		Short: "Encode binary data to PGP word list",
		Long: `Encode binary data to mnemonic PGP word list.

For PGP word list spec see:
https://en.wikipedia.org/wiki/PGP_word_list
`,
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)
			encoder := pgp.NewEncoder()

			if flagNoArmor == false {
				fmt.Println("-----BEGIN PGP WORD LIST-----")
			}

			separator := ""
			for {
				mybyte, err := reader.ReadByte()
				if err != nil {
					break // EOF
				}
				word := encoder.GetWord(mybyte)
				fmt.Print(separator)
				fmt.Print(word)
				separator = " "
			}
			if flagNoArmor == false {
				fmt.Println("")
				fmt.Println("-----END PGP WORD LIST-----")
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(pgpCmd)
}
