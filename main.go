package main

import "github/Mr-Intern/questionable-cli/cmd"
import "github/Mr-Intern/questionable-cli/configs"

func main() {
	cmd.Execute()
	configs.GetCurrentEpisode()
}
