package promptInputs

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

type PromptContent struct {
	Label    string
	ErrorMsg string
}

type PromptContentSelect struct {
	Label    string
	Items    []string
	ErrorMsg string
}

func PromptGetInput(pc PromptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.ErrorMsg)
		}

		return nil
	}
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.Label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}

func PromptGetSelect(pc PromptContentSelect) string {
	index := -1

	var result string
	var err error

	items := pc.Items

	templates := &promptui.SelectTemplates{
		Label:  "{{ . | bold }}",
		Active: "\u279c {{ . | bold }}",

		Selected: fmt.Sprintf("%s {{ . }}", pc.Label),
	}

	for index < 0 {
		prompt := promptui.Select{
			Label:     pc.Label,
			Items:     items,
			Pointer:   promptui.BlockCursor,
			Templates: templates,
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}
