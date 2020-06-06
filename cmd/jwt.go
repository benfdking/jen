package cmd

import (
	"fmt"
	"log"
	"os"

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
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return fmt.Errorf("incorrect number of arguments %d, must be 1 or less", len(args))
		}
		if len(args) == 1 {
			if _, err := os.Stat(args[0]); os.IsNotExist(err) {
				return fmt.Errorf("file %s does not exist", args[0])
			}
		}
		return nil
	},
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
