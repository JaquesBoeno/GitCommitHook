package prompts

import (
	"errors"
	"fmt"
	"strings"

	commitMessage "github.com/JaquesBoeno/GitCommitHook/commit"
	"github.com/JaquesBoeno/GitCommitHook/config"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Questions []config.Question
	Responses []commitMessage.Value

	currentTextinput textinput.Model
	currentCursor    int
	currentQuestion  config.Question
	pastResponses    string
	index            int
	Err              error
}

type (
	errMsg error
)

func InitialModel(questions []config.Question) Model {
	ti := textinput.New()

	if questions[0].Type == "text" {
		ti.Placeholder = "Message here"
		ti.Focus()
		ti.CharLimit = questions[0].Max
		ti.Width = 100
		ti.ShowSuggestions = true
		ti.Prompt = ""
	}

	return Model{
		currentTextinput: ti,
		Questions:        questions,
		currentQuestion:  questions[0],
		currentCursor:    0,
		index:            0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.Err = errors.New("exit")
			return m, tea.Quit
		}
	}

	switch m.currentQuestion.Type {
	case "select":
		return selectUpdate(msg, m)
	case "text":
		return textUpdate(msg, m)
	case "none":
		return m, nil
	}

	return m, nil
}

func (m Model) View() string {
	str := strings.Builder{}
	str.WriteString(m.pastResponses + "\n")

	switch m.currentQuestion.Type {
	case "select":
		str.WriteString(selectView(m))
	case "text":
		str.WriteString(textView(m))
	case "none":
		return str.String()
	}

	return str.String()
}

func (m Model) NextQuestion() (tea.Model, tea.Cmd) {
	if m.index >= len(m.Questions)-1 {
		m.currentQuestion = config.Question{
			Type: "none",
		}
		return m, tea.Quit
	} else {
		m.index++
		m.currentCursor = 0
		m.currentQuestion = m.Questions[m.index]

		if m.currentQuestion.Type == "text" {
			ti := textinput.New()
			ti.Placeholder = "Message here"
			ti.Focus()
			ti.CharLimit = m.Questions[m.index].Max
			ti.Width = 100
			ti.ShowSuggestions = true
			ti.Prompt = ""
			m.currentTextinput = ti
		}

		return m, nil
	}
}

func (m Model) checkLen() error {
	var err error

	c := m.currentTextinput
	if len(c.Value()) < m.currentQuestion.Min {
		err = fmt.Errorf(
			"minimal char: %d",
			m.currentQuestion.Min,
		)
	} else if len(c.Value()) > m.currentQuestion.Max {
		err = fmt.Errorf(
			"max char: %d",
			m.currentQuestion.Max,
		)
	}
	return err
}
