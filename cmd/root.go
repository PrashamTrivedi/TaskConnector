package cmd

import "github.com/spf13/cobra"

var configFile string

var RootCmd = &cobra.Command{
	Use:   "taskConnector",
	Short: "A Connector that connect a URL to a task running local in your machine",
	
}


func init()  {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVarP(&configFile,"configFile","c","./taskconnectorConfig.json","Path of configuration file")
}

func initConfig(){
	if configFile != ""{
		SetcommandFile(configFile)
	}
}