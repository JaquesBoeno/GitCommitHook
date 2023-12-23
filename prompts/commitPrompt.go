package prompts

import (
	"errors"
	"fmt"
	"math"
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
		{
			switch msg := msg.(type) {
			case tea.KeyMsg:
				switch msg.String() {
				case "enter":
					m.Responses = append(m.Responses, commitMessage.Value{
						Id:    m.currentQuestion.Id,
						Value: m.currentQuestion.Options[m.currentCursor].Name,
					})
					m.pastResponses = fmt.Sprint(m.pastResponses, fmt.Sprintf("\033[1m%s:\033[0m %s\n",
						m.currentQuestion.Label,
						m.currentQuestion.Options[m.currentCursor].Name))
					return m.NextQuestion()

				case "up", "k":
					m.currentCursor--
					if m.currentCursor < 0 {
						m.currentCursor = len(m.currentQuestion.Options) - 1
					}

				case "down", "j":
					m.currentCursor++
					if m.currentCursor >= len(m.currentQuestion.Options) {
						m.currentCursor = 0
					}
				}
			}
		}
	case "text":
		var cmd tea.Cmd

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyEnter:
				m.Err = m.checkLen()

				if m.Err == nil {
					m.pastResponses = fmt.Sprint(m.pastResponses, fmt.Sprintf("\033[1m%s:\033[0m\n> %s\n",
						m.currentQuestion.Label,
						m.currentTextinput.Value()))
					m.Responses = append(m.Responses, commitMessage.Value{
						Id:    m.currentQuestion.Id,
						Value: m.currentTextinput.Value(),
					})

					return m.NextQuestion()
				}
			}

		case errMsg:
			m.Err = msg
			return m, nil
		}

		m.currentTextinput, cmd = m.currentTextinput.Update(msg)
		return m, cmd
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
		{
			questionPerPage := 5
			questionInPage := make([]config.Option, questionPerPage)

			str.WriteString(fmt.Sprintf("%s:\n\n", m.currentQuestion.Label))
			half := int(math.Floor(float64(questionPerPage) / 2))
			start, end := 0, 0
			start = m.currentCursor - half
			end = m.currentCursor + half + 1

			if start < 0 {
				start = 0
				end = questionPerPage
			}
			if end > len(m.currentQuestion.Options) {
				end = len(m.currentQuestion.Options)
				start = 0
				if len(m.currentQuestion.Options)-questionPerPage >= 0 {
					start = len(m.currentQuestion.Options) - questionPerPage
				}
			}

			questionInPage = m.currentQuestion.Options[start:end]

			var maxNameLen int

			for _, v := range m.currentQuestion.Options {
				padding := 5
				if len(v.Name)+padding > maxNameLen {
					maxNameLen = len(v.Name) + padding
				}
			}

			for i, option := range questionInPage {
				var desc string
				var name string

				if strings.Contains(option.Desc, "\n") {
					split := strings.Split(option.Desc, "\n")
					desc = fmt.Sprintf("%s\n %s %s", split[0], strings.Repeat(" ", maxNameLen), split[1])
				} else {
					desc = option.Desc
				}
				name = fmt.Sprintf("%s%s", option.Name, strings.Repeat(" ", maxNameLen-len(option.Name)))

				if m.currentCursor == i+start {
					str.WriteString(fmt.Sprintf("\033[36m\033[1m❯ %s%s\033[0m", name, desc))
				} else {
					str.WriteString(fmt.Sprintf("  %s%s", name, desc))
				}

				str.WriteString("\n")
			}
		}
	case "text":
		{

			str.WriteString(fmt.Sprintf(
				"%s\n%d %s\n",
				m.currentQuestion.Label,
				len(m.currentTextinput.Value()),
				m.currentTextinput.View(),
			) + "\n")

		}
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

func chars(len int, min int, max int) string {
	if len < min || len > max {
		return fmt.Sprintf("\033[31m(%d)\033[0m", len)
	}
	return fmt.Sprintf("\033[32m(%d)\033[0m", len)
}
