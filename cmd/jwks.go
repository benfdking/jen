package cmd

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"github.com/benfdking/jen/pkg/rsapem"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

var (
	outputJWKS    string
	outputPrivate string
)

// jwksCmd represents the jwks command
var jwksCmd = &cobra.Command{
	Use:   "jwks",
	Short: "Generates a jwk set with optional private and public key",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		privkey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			log.Fatalf("failed to generate private key: %s", err)
		}
		key, err := jwk.New(&privkey.PublicKey)
		if err != nil {
			log.Fatalf("failed to create JWK: %s", err)
		}
		jsonbuf, err := json.MarshalIndent(jwk.Set{Keys: []jwk.Key{key}}, "", "  ")
		if err != nil {
			log.Fatalf("failed to generate JSON: %s", err)
		}

		err = ioutil.WriteFile(outputJWKS, jsonbuf, os.ModePerm)
		if err != nil {
			log.Fatalf("failed to write jwk: %s", err)
		}
		// TODO MarshalPKCS1PrivateKey change this
		err = ioutil.WriteFile(outputPrivate, rsapem.RSAPrivateToPem(privkey), os.ModePerm)
		if err != nil {
			log.Fatalf("failed to write private key: %s", err)
		}
		//if outputPublic != "" {
		//	err = ioutil.WriteFile(outputPublic, x509.MarshalPKCS1PublicKey(&privkey.PublicKey), os.ModePerm)
		//	if err != nil {
		//		log.Fatalf("failed to write public key: %s", err)
		//	}
		//}
	},
}

func init() {
	rootCmd.AddCommand(jwksCmd)

	jwksCmd.Flags().StringVarP(&outputJWKS, "jwk", "j", "jwks.json", "Sets the jwks json output")
	jwksCmd.Flags().StringVarP(&outputPrivate, "private", "p", "private.txt", "Sets output path of private key")
}
