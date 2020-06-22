package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	buildVersion = ""
	buildCommit  = ""
	buildDate    = ""
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Return build information",
	Run: func(cmd *cobra.Command, args []string) {
		w := newTabWriter()
		fmt.Fprintf(w, "%s\t%s\n", "Version", buildVersion)
		fmt.Fprintf(w, "%s\t%s\n", "Commit", buildCommit)
		fmt.Fprintf(w, "%s\t%s\n", "Time", buildDate)
		w.Flush()
	},
}
