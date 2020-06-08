package cmd

import (
	"fmt"
	"log"

	"github.com/benfdking/jen/pkg/url"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [abc]",
	Short: "List the default keys, optionally specifying one only returns one jwks url",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			jwksURL, _, err := url.ReturnJWKSAndPrivatePEMURL(args[0])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(jwksURL)
		} else {
			w := newTabWriter()
			fmt.Fprintf(w, "%s\t%s\t%s\n", "Key", "JWKS URL", "Private PEM URL")
			const alphabet = "abc"
			for _, c := range alphabet {
				jwks, pem, err := url.ReturnJWKSAndPrivatePEMURL(string(c))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Fprintf(w, "%s\t%s\t%s\n", string(c), jwks, pem)
			}
			w.Flush()
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
