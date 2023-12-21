/*
Copyright © 2023 JAQUES BOENO jaquesboeno@proton.me
*/
package cmd

import (
	"github.com/JaquesBoeno/GitHook/config"
	promptInputs "github.com/JaquesBoeno/GitHook/prompt"
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
			ask := promptInputs.PromptContentSelect{
				Label: question.Label,
				Items: question.Items,
			}

			response := promptInputs.PromptContentSelect(ask).
				fmt.Println(response)
		} else {

		}
	}
}
