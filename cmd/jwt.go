package cmd

import (
	"crypto"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/benfdking/jen/pkg/defaultjwt"
	"github.com/benfdking/jen/pkg/rsapem"
	"github.com/benfdking/jen/pkg/url"
	"github.com/spf13/cobra"
)

var key string
var keyFilePath string
var addDefaults bool
var filePath string
var claims map[string]string

// jwtCmd represents the jwt command
var jwtCmd = &cobra.Command{
	Use:   "jwt [claims json]",
	Short: "Generate a jwt",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		var privateKey crypto.PrivateKey
		if keyFilePath != "" {
			keyBytes, err := ioutil.ReadFile(keyFilePath)
			if err != nil {
				log.Fatal(err)
			}
			privateKey, err = rsapem.RSAPrivateFromPen(keyBytes)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			var err error
			_, privateKey, err = url.ReturnJWKSAndPrivateKey(key)
			if err != nil {
				log.Fatal(err)
			}
		}

		token := defaultjwt.NewToken()
		if addDefaults {
			var err error
			token, err = defaultjwt.AddDefaultClaims(token)
			if err != nil {
				log.Fatal(err)
			}
		}
		if filePath != "" {
			var err error
			token, err = defaultjwt.AddJSONFileClaimsToToken(token, filePath)
			if err != nil {
				log.Fatal(err)
			}
		}
		if len(args) == 1 {
			var err error
			token, err = defaultjwt.AddJSONStringClaimsToToken(token, args[0])
			if err != nil {
				log.Fatal(err)
			}
		}
		if len(claims) > 0 {
			var err error
			token, err = defaultjwt.AddMapClaimsToToken(token, claims)
			if err != nil {
				log.Fatal(err)
			}
		}

		s, err := defaultjwt.SignToken(token, privateKey)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(s)
	},
}

func init() {
	rootCmd.AddCommand(jwtCmd)

	jwtCmd.Flags().StringToStringVarP(&claims, "claims", "c", map[string]string{}, "Claims for JWT")
	jwtCmd.Flags().StringVarP(&key, "key", "k", "a", "[abc] jwt key to use")
	jwtCmd.Flags().BoolVarP(&addDefaults, "defaults", "d", true, "Add default claims")
	jwtCmd.Flags().StringVarP(&filePath, "file", "f", "", "Add claims from JSON file")
	jwtCmd.Flags().StringVarP(&keyFilePath, "private", "p", "", "Use private key to sign jwt")
}
