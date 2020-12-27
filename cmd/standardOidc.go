package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/benfdking/jen/pkg/defaultjwt"
	"github.com/spf13/cobra"
)

var standardOidcCmd = &cobra.Command{
	Use:   "oidc",
	Short: "Return sample standard oidc claims",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		cs := defaultjwt.OIDCStandardClaims()
		bytes, err := json.Marshal(cs)
		if err != nil {
			return err
		}
		fmt.Println(string(bytes))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(standardOidcCmd)
}
