package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/benfking/jen/pkg/defaultjwt"
	"github.com/benfking/jen/pkg/url"
	"github.com/spf13/cobra"
)

var key string

// jwtCmd represents the jwt command
var jwtCmd = &cobra.Command{
	Use:   "jwt [json file path]",
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
		jwtString, err := defaultjwt.Default(pem)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(jwtString)
	},
}

func init() {
	rootCmd.AddCommand(jwtCmd)

	jwtCmd.Flags().StringVarP(&key, "key", "k", "a", "[abc] jwt key to use")
}
