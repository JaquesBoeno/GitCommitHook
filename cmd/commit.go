/*
Copyright © 2023 JAQUES BOENO jaquesboeno@proton.me
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/JaquesBoeno/GitHook/config"
	"github.com/JaquesBoeno/GitHook/prompts"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Create a commit",
	Long:  `creating commit with the pattern you defined in the settings file`,
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}

func Run() {
	config := config.ReadConfigs()
	p := tea.NewProgram(prompts.InitialModel(config.Questions))
	_, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}
}
