/*
Copyright © 2023 JAQUES BOENO jaquesboeno@proton.me
*/
package cmd

import (
	"fmt"

	promptInputs "github.com/JaquesBoeno/GitHook/utils"

	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Create a commit",
	Long:  `creating commit with the pattern you defined in the settings file`,
	Run: func(cmd *cobra.Command, args []string) {
		whatsYourName()
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}

func whatsYourName() {
	namePromptCommand := promptInputs.PromptContent{
		Label:    "What's your name?",
		ErrorMsg: "write your name",
	}
	name := promptInputs.PromptGetInput(namePromptCommand)

	teamPromptCommand := promptInputs.PromptContentSelect{
		Label:    "What's your team?",
		ErrorMsg: "select a valid team",
		Items:    []string{"gremio", "internacional"},
	}

	team := promptInputs.PromptGetSelect(teamPromptCommand)

	fmt.Println(fmt.Sprintf(`Hi! %s
You team are %s`, name, team))
}
