package cmd

import (
	"fmt"

	"github.com/benfdking/jen/pkg/defaultjwt"
	"github.com/spf13/cobra"
)

var defaultsCmd = &cobra.Command{
	Use:   "defaults",
	Short: "returns default claims",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cs := defaultjwt.GetDefaultClaims()
		for _, c := range cs {
			fmt.Printf("key: %s\tvalue: %s\tdescription: %s\n", c.Key, c.Value, c.Description)
		}
	},
}

func init() {
	rootCmd.AddCommand(defaultsCmd)
}
