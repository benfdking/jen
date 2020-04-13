package cmd

import (
	"github.com/spf13/cobra"
)

// jwtCmd represents the jwt command
var jwtCmd = &cobra.Command{
	Use:   "jwt",
	Short: "Generate a jwt",
	Run: func(cmd *cobra.Command, args []string) {
		//privkey, err := rsa.GenerateKey(rand.Reader, 2048)
		//if err != nil {
		//	log.Fatalf("failed to generate private key: %s", err)
		//}
		//key, err := jwk.New(&privkey.PublicKey)
		//if err != nil {
		//	log.Fatalf("failed to create JWK: %s", err)
		//}
		//jsonbuf, err := json.MarshalIndent(key, "", "  ")
		//if err != nil {
		//	log.Fatalf("failed to generate JSON: %s", err)
		//}
		//
		//err = ioutil.WriteFile(outputJWK, jsonbuf, os.ModePerm)
		//if err != nil {
		//	log.Fatalf("failed to write jwk: %s", err)
		//}
		//// TODO MarshalPKCS1PrivateKey change this
		//err = ioutil.WriteFile(outputPrivate, rsapem.RSAPrivateToPem(privkey), os.ModePerm)
		//if err != nil {
		//	log.Fatalf("failed to write private key: %s", err)
		//}
		//if outputPublic != "" {
		//	err = ioutil.WriteFile(outputPublic, x509.MarshalPKCS1PublicKey(&privkey.PublicKey), os.ModePerm)
		//	if err != nil {
		//		log.Fatalf("failed to write public key: %s", err)
		//	}
		//}
	},
}

func init() {
	rootCmd.AddCommand(jwtCmd)
}
