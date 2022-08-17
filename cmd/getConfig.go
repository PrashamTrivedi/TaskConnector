package cmd

import "github.com/spf13/cobra"

var getConfig = &cobra.Command{
	Use: "getConfig",
	Short:"Lists configurations",
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}