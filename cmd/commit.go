/*
Copyright © 2023 JAQUES BOENO jaquesboeno@proton.me
*/
package cmd

import (
	"fmt"

	"github.com/JaquesBoeno/GitHook/commit"
	"github.com/JaquesBoeno/GitHook/config"
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

	fmt.Println(commit.CommitMessageBuilder(config.TemplateCommit, []commit.Value{
		{Id: "type", Value: "feat"},
		{Id: "scope", Value: "front-end"},
		{Id: "subject", Value: "user card added"},
		{Id: "desc", Value: "some description"},
	}))

	// 	var scope, typ string
	// 	var subject, description string

	// 	for i, question := range questions {
	// 		if question.Items != nil {
	// 			p := tea.NewProgram(prompts.InitialModel(question))
	// 			// Run returns the model as a tea.Model.
	// 			m, err := p.Run()
	// 			if err != nil {
	// 				fmt.Println("Oh no:", err)
	// 				os.Exit(1)
	// 			}

	// 			if m, ok := m.(prompts.SelectModel); ok && m.Choice.Name != "" {
	// 				switch i {
	// 				case 0:
	// 					scope = m.Choice.Name
	// 				case 1:
	// 					typ = m.Choice.Name
	// 				}
	// 			}

	// 		} else {
	// 			p := tea.NewProgram(prompts.InitialModelInput(question))
	// 			m, err := p.Run()
	// 			if err != nil {
	// 				fmt.Println("Oh no:", err)
	// 				os.Exit(1)
	// 			}

	// 			if m, ok := m.(prompts.InputModel); ok && m.TextInput.Value() != "" {
	// 				switch i {
	// 				case 2:
	// 					subject = m.TextInput.Value()
	// 				case 3:
	// 					description = m.TextInput.Value()
	// 				}
	// 			}
	// 		}
	// 	}

	// 	fmt.Println(fmt.Sprintf(`%s (%s): %s

	// %s
	// `, typ, scope, subject, description))
}
