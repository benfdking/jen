package cmd

import (
	"fmt"
	"log"

	"github.com/benfdking/jen/pkg/defaultjwt"
	"github.com/benfdking/jen/pkg/url"
	"github.com/spf13/cobra"
)

var key string
var addDefaults bool

// jwtCmd represents the jwt command
var jwtCmd = &cobra.Command{
	Use:   "jwt [optional json file path]",
	Short: "Generate a jwt",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		_, pem, err := url.ReturnJWKSAndPrivateKey(key)
		if err != nil {
			log.Fatal(err)
		}

		token := defaultjwt.NewToken()
		if addDefaults {
			token, err = defaultjwt.AddDefaultClaims(token)
			if err != nil {
				log.Fatal(err)
			}
		}
		if len(args) == 1 {
			token, err = defaultjwt.AddJSONFileClaimsToToken(token, args[0])
			if err != nil {
				log.Fatal(err)
			}
		}

		s, err := defaultjwt.SignToken(token, pem)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(s)
	},
}

func init() {
	rootCmd.AddCommand(jwtCmd)

	jwtCmd.Flags().StringVarP(&key, "key", "k", "a", "[abc] jwt key to use")
	jwtCmd.Flags().BoolVarP(&addDefaults, "defaults", "d", true, "adds default oidc parameters, true by default")
}
