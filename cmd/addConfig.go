/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var key string
var value string

// addConfigCmd represents the addConfig command
var addConfigCmd = &cobra.Command{
	Use:   "addConfig",
	Short: "Add A configuration",

	Run: func(cmd *cobra.Command, args []string) {
		AddConfig(key, value)
		fmt.Printf("Key Value pair %s:%s added \n", key, value)
	},
}

func init() {
	addConfigCmd.Flags().StringVarP(&key, "key", "k", "", "Key of property")
	addConfigCmd.Flags().StringVarP(&value, "value", "v", "", "Value of property")

	addConfigCmd.MarkFlagRequired("key")
	addConfigCmd.MarkFlagRequired("value")

	RootCmd.AddCommand(addConfigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addConfigCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addConfigCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
