package cmd

import (
	"github.com/benfdking/jen/pkg/server"
	"github.com/spf13/cobra"
)

var port int

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Server which can be used to provide jwks and jwts",
	Args:  cobra.RangeArgs(0, 0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return server.Serve(port)
	},
}

func init() {
	serveCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to serve on")
	rootCmd.AddCommand(serveCmd)
}
