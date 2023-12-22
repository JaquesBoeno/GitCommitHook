package prompts

import (
	"fmt"

	"github.com/JaquesBoeno/GitHook/config"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type InputModel struct {
	TextInput textinput.Model
	Question  config.Question
	err       error
}

func (m InputModel) checkMinLen() error {
	var err error

	c := m.TextInput
	if len(c.Value()) < m.Question.Min {
		err = fmt.Errorf(
			"minimal char: %d",
			m.Question.Min,
		)
	}
	return err
}

func validate(s string) error {
	if len(s) < 6 {
		return fmt.Errorf("too short")
	}

	return nil
}

func InitialModelInput(question config.Question) InputModel {
	ti := textinput.New()
	ti.Placeholder = "Message here"
	ti.Focus()
	ti.CharLimit = question.Max
	ti.Width = 100
	ti.ShowSuggestions = true

	return InputModel{
		TextInput: ti,
		err:       nil,
		Question:  question,
	}
}

func (m InputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m InputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			m.err = m.checkMinLen()

			if m.err == nil {
				return m, tea.Quit
			}
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.TextInput, cmd = m.TextInput.Update(msg)
	return m, cmd
}

func chars(len int, min int, max int) string {
	if len < min || len > max {
		return fmt.Sprintf("\033[31m(%d)\033[0m", len)
	}
	return fmt.Sprintf("\033[32m(%d)\033[0m", len)
}

func (m InputModel) View() string {

	return fmt.Sprintf(
		"%s\n%s %s\n",
		m.Question.Label,
		chars(len(m.TextInput.Value()), 1, 66),
		m.TextInput.View(),
	) + "\n"
}
