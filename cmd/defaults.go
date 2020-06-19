package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/benfdking/jen/pkg/defaultjwt"
	"github.com/spf13/cobra"
)

var explain bool

var defaultsCmd = &cobra.Command{
	Use:   "defaults",
	Short: "Return default claims",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		cs := defaultjwt.GetDefaultClaims()
		if explain {
			w := newTabWriter()
			fmt.Fprintf(w, "%s\t%s\t%s\n", "Key", "Value", "Description")
			for _, c := range cs {
				fmt.Fprintf(w, "%s\t%s\t%s\n", c.Key, c.Value, c.Description)
			}
			return w.Flush()
		}
		m := make(map[string]string, len(cs))
		for _, c := range cs {
			m[c.Key] = c.Value
		}
		bytes, err := json.Marshal(m)
		if err != nil {
			return err
		}
		fmt.Println(string(bytes))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(defaultsCmd)

	defaultsCmd.Flags().BoolVarP(&explain, "explain", "e", false, "Explain default values")
}
