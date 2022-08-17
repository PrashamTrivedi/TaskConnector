package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getConfig = &cobra.Command{
	Use:   "getConfig",
	Short: "Lists configurations",
	Run: func(cmd *cobra.Command, args []string) {
		commandMapping := GetCommandMapping()
		for key, val := range commandMapping {
			fmt.Printf("%s: %s\n", key, val)
		}
	},
}

func init() {
	RootCmd.AddCommand(getConfig)
}
