/*
Copyright © 2023 JAQUES BOENO jaquesboeno@proton.me
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	commitMessage "github.com/JaquesBoeno/GitCommitHook/commit"
	"github.com/JaquesBoeno/GitCommitHook/config"
	"github.com/JaquesBoeno/GitCommitHook/git"
	"github.com/JaquesBoeno/GitCommitHook/prompts"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Create a commit",
	Long:  `creating commit with the pattern you defined in the settings file`,
	Run: func(cmd *cobra.Command, args []string) {
		hook, err := cmd.Flags().GetString("hook")
		if err != nil {
			panic(err)
		}

		Run(hook)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
	commitCmd.PersistentFlags().String("hook", "", "A search term for a dad joke.")
}

func Run(hook string) {
	config := config.ReadConfigs()

	p := tea.NewProgram(prompts.InitialModel(config.Questions))
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	if m, ok := m.(prompts.Model); ok && m.Err == nil && len(m.Responses) > 0 {
		message := commitMessage.CommitMessageBuilder(config.TemplateCommit, m.Responses)

		if len(hook) > 0 {
			git.Hook(message, hook)
		} else {
			_, err := git.Commit(message)
			if err != nil {
				log.Printf("run git commit failed, err=%v\n", err)
				log.Printf("commit message is: \n\n%s\n\n", string(message))
			}
		}
	}
}
