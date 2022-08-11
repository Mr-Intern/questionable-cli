/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// oscpCmd represents the oscp command
var oscpCmd = &cobra.Command{
	Use:   "oscp",
	Short: "get notes on hacking oscp type machines",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("oscp called")
	},
}

func init() {
	hackCmd.AddCommand(oscpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// oscpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// oscpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
