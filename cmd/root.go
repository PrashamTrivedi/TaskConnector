package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "taskConnector",
	Short: "A Connector that connect a URL to a task running local in your machine",
}


func init()  {
	
}