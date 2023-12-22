package prompts

import (
	"errors"
	"fmt"

	"github.com/JaquesBoeno/GitHook/config"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type inputModel struct {
	textInput textinput.Model
	question  config.Question
	err       error
}

func validate(s string) error {
	if len(s) <= 0 {
		return errors.New("write some text, not possible null")
	} else {
		return nil
	}
}

func InitialModelInput(question config.Question) inputModel {
	ti := textinput.New()
	ti.Placeholder = "Message here"
	ti.Focus()
	ti.CharLimit = question.Max
	ti.Validate = validate
	ti.Width = 100
	ti.ShowSuggestions = true

	return inputModel{
		textInput: ti,
		err:       nil,
		question:  question,
	}
}

func (m inputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m inputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func chars(len int, min int, max int) string {
	if len < min || len > max {
		return fmt.Sprintf("\033[31m(%d)\033[0m", len)
	}
	return fmt.Sprintf("\033[32m(%d)\033[0m", len)
}

func (m inputModel) View() string {

	return fmt.Sprintf(
		"%s\n%s %s\n",
		m.question.Label,
		chars(len(m.textInput.Value()), 1, 66),
		m.textInput.View(),
	) + "\n"
}
