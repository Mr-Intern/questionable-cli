/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// hackCmd represents the hack command
var hackCmd = &cobra.Command{
	Use:   "hack",
	Short: "get hacking notes",
	Long: `
	list - print list of hacking topics
	ex: qcli hack list
	ex: qcli hack kubernetes
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hack called")
	},
}

func init() {
	rootCmd.AddCommand(hackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
