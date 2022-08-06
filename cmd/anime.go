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
	Long: `prints a list of links to various anime and manga I enjoy`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("--------------------------------------------------------------------------------")
		fmt.Println("ONE PIECE")
		fmt.Println("https://9anime.lu/watch/one-piece.ov8/ep-1")
		fmt.Println("--------------------------------------------------------------------------------")
		fmt.Println("BESERK manga")
		fmt.Println("https://readberserk.com/chapter/berserk-chapter-a0/")
		fmt.Println("--------------------------------------------------------------------------------")
	},
	/*	
	// do the same thing but with a loop, and less pretty
	fmt.Println("testing map and loop implementation...")

	var animeMap map[string]string
	// create a map
	animeMap = make(map[string]string)
	
	// insert key-value pairs in the map
	animeMap["One Piece"] = "https://9anime.lu/watch/one-piece.ov8/ep-1"
	animeMap["Beserk"] = "https://readberserk.com/chapter/berserk-chapter-a0/"

	// print map using keys
	for show := range animeMap {
	fmt.Println(show, "," , animeMap[show])
	}
	*/
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
