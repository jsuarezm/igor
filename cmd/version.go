package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show igor version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(RootCmd.Use + " v" + VERSION)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
