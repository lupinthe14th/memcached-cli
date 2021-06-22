package cmd

import (
	"fmt"

	"github.com/lupinthe14th/memcached-cli/pkg/version"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of memcached-cli.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Client:\n")
		fmt.Printf(" Version:\t%s\n", version.Version)
		fmt.Printf(" Git commit:\t%s\n", version.Revision)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
