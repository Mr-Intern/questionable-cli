/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "print notes for hacking web apps",
	Long: `usage: qcli hack web`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("web called")
	},
}

func init() {
	hackCmd.AddCommand(webCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// webCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
