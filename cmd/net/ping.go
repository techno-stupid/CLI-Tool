/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"time"
)

var (
	urlPath string
	client  = http.Client{
		Timeout: time.Second * 2,
	}
)

func ping(domain string) (int, error) {
	url := "https://" + domain

	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil
	}
	resp.Body.Close()
	return resp.StatusCode, nil
}

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Send a ping request to a remote server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := ping(urlPath)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	},
}

func init() {

	pingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "The url to ping")

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
