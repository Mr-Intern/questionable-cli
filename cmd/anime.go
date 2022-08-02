/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// animeCmd represents the anime command
var animeCmd = &cobra.Command{
	Use:   "anime",
	Short: "watch anime",
	Long: `A command that opens up an anime to watch.
	At the time of writing, that anime is OnePiece`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("anime called")
	},
}

func init() {
	rootCmd.AddCommand(animeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// animeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// animeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
