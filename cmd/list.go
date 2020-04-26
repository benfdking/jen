package cmd

import (
	"fmt"
	"log"

	"github.com/benfking/jen/pkg/url"
	"github.com/spf13/cobra"
)


// listCmd represents the list command
var listCmd = &cobra.Command{
	// TODO Make this dependent on versions passed in
	Use:   "list [abc]",
	Short: "lists the default keys that can be used, optionally specifying one only returns one",
	Args: func(cmd *cobra.Command, args []string) error {
		switch {
		case len(args) > 1:
			return fmt.Errorf("can only provide no or 1 argument, not %d arguments", len(args))
		case len(args) == 1:
			if !url.IsVersion(args[0]) {
				return fmt.Errorf("can only be one of the aviable keys: %s", url.ReturnVersions())
			}
			return nil
		default:
			return nil
		}
	},
	Run: func(cmd *cobra.Command, args []string) {


		const alphabet = "abc"
		for _, c := range alphabet {
			jwksURL, _, err := url.ReturnJWKSAndPrivatePEMURL(string(c))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("key %c, jwks url: "+jwksURL+"\n", c)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// TODO ADD Showing pem url
	//var showPemURL bool
	//listCmd.Flags().BoolVarP(&showPemURL, "pem", "p", false, "Show URL For Private key")
}
