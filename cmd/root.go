package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jen",
	Short: "jen is your go to cmd tool for all things jwt",
}

var (
	buildVersion = ""
	buildCommit  = ""
	buildDate    = ""
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Return build information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version:\t%s\nCommit:\t\t%s\nTime:\t\t%s\n", buildVersion, buildCommit, buildDate)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
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
