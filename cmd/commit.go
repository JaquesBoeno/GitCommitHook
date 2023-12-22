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
		commit()
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}

func commit() {
	questions := config.ReadConfigs().Questions

	for _, question := range questions {
		if question.Items != nil {
			p := tea.NewProgram(prompts.InitialModel(question))
			if _, err := p.Run(); err != nil {
				fmt.Printf("Alas, there's been an error: %v", err)
				os.Exit(1)
			}
		} else {
		}
	}
}
