/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// kubernetesCmd represents the kubernetes command
var kubernetesCmd = &cobra.Command{
	Use:   "kubernetes",
	Short: "print notes for hacking kubernetes",
	Long: `usage: qcli hack kubernetes`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO: put kubernetes notes here")
	},
}

func init() {
	hackCmd.AddCommand(kubernetesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kubernetesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kubernetesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
