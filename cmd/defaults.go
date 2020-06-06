package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/benfdking/jen/pkg/defaultjwt"
	"github.com/spf13/cobra"
)

var explain bool

var defaultsCmd = &cobra.Command{
	Use:   "defaults",
	Short: "Return default claims",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cs := defaultjwt.GetDefaultClaims()
		if explain {
			for _, c := range cs {
				fmt.Printf("key: %s\tvalue: %s\tdescription: %s\n", c.Key, c.Value, c.Description)
			}
		} else {
			m := make(map[string]string, len(cs))
			for _, c := range cs {
				m[c.Key] = c.Value
			}
			bytes, err := json.Marshal(m)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(bytes))
		}
	},
}

func init() {
	rootCmd.AddCommand(defaultsCmd)

	defaultsCmd.Flags().BoolVarP(&explain, "explain", "e", false, "outputs a table explaining the default values")
}
