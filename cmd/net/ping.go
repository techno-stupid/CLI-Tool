/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	urlpath string
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Send a ping request to a remote server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	pingCmd.Flags().StringVarP(&urlpath, "url", "u", "", "The url to ping")

	if err := pingCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println(err)
	}

	NetCmd.AddCommand(pingCmd)
	//cmd.rootCmd.AddCommand(pingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
