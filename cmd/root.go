package cmd

import (
	"crypto"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/tabwriter"

	"github.com/benfdking/jen/pkg/defaultjwt"
	"github.com/benfdking/jen/pkg/oidc"
	"github.com/benfdking/jen/pkg/rsapem"
	"github.com/benfdking/jen/pkg/url"
	"github.com/spf13/cobra"
)

var (
	key                string
	keyFilePath        string
	addDefaults        bool
	filePath           string
	oidcStandardClaims bool
	claims             map[string]string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jen [claims json]",
	Short: "Jenerate a JWT token",
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
		if oidcStandardClaims {
			var err error
			token, err = defaultjwt.AddMapClaimsToToken(token, oidc.StandardClaims())
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
	rootCmd.AddCommand(versionCmd)

	rootCmd.Flags().BoolVarP(&oidcStandardClaims, "oidc", "o", false, "Flag to insert OIDC standard claims")
	rootCmd.Flags().StringToStringVarP(&claims, "claims", "c", map[string]string{}, "Claims for JWT")
	rootCmd.Flags().StringVarP(&key, "key", "k", "a", "[abc] jwt key to use")
	rootCmd.Flags().BoolVarP(&addDefaults, "defaults", "d", true, "Add default claims")
	rootCmd.Flags().StringVarP(&filePath, "file", "f", "", "Add claims from JSON file")
	rootCmd.Flags().StringVarP(&keyFilePath, "private", "p", "", "Use private key to sign jwt")
}

func newTabWriter() *tabwriter.Writer {
	return tabwriter.NewWriter(os.Stdout, 2, 0, 3, ' ', tabwriter.TabIndent)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version, commit, date string) {
	buildVersion, buildCommit, buildDate = version, commit, date
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
