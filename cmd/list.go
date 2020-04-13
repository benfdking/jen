package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [abc]",
	Short: "lists the default keys that can be used, optionally specifying one only returns one",
	//Args: func(cmd *cobra.Command, args []string) error {
	//	if len(args) > 1 {
	//		return fmt.Errorf("can only provide no or 1 argument, not %d", len(args))
	//	}
	//	// TODO Add check that it's a b or c
	//	return nil
	//},
	Run: func(cmd *cobra.Command, args []string) {
		const alphabet = "abc"
		for _, c := range alphabet {
			fmt.Printf("%c, URL: " + url +"\n", c, c)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
