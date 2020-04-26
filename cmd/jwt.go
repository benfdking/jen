package cmd

import (
	"fmt"
	"github.com/benfking/jen/pkg/defaultjwt"
	"github.com/benfking/jen/pkg/url"
	"github.com/spf13/cobra"
	"log"
)

// jwtCmd represents the jwt command
var jwtCmd = &cobra.Command{
	Use:   "jwt",
	Short: "Generate a jwt",
	Run: func(cmd *cobra.Command, args []string) {
		_, pem, err := url.ReturnJWKSAndPrivateKey("a")
		if err != nil {
			log.Fatal(err)
		}
		jwtString, err :=  defaultjwt.Default(pem)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(jwtString)
	},
}

func init() {
	rootCmd.AddCommand(jwtCmd)
}
