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
var filePath string
var claims map[string]string

// jwtCmd represents the jwt command
var jwtCmd = &cobra.Command{
	Use:   "jwt [claims json]",
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
		if filePath != "" {
			token, err = defaultjwt.AddJSONFileClaimsToToken(token, filePath)
			if err != nil {
				log.Fatal(err)
			}
		}
		if len(args) == 1 {
			token, err = defaultjwt.AddJSONStringClaimsToToken(token, args[0])
			if err != nil {
				log.Fatal(err)
			}
		}
		if len(claims) > 0 {
			token, err = defaultjwt.AddMapClaimsToToken(token, claims)
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

	jwtCmd.Flags().StringToStringVarP(&claims, "claims", "c", map[string]string{}, "claims to be added to the jwt")
	jwtCmd.Flags().StringVarP(&key, "key", "k", "a", "[abc] jwt key to use")
	jwtCmd.Flags().BoolVarP(&addDefaults, "defaults", "d", true, "adds default oidc parameters, true by default")
	jwtCmd.Flags().StringVarP(&filePath, "file", "f", "", "json file to read claims from")
}
