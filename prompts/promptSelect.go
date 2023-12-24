package prompts

import (
	"fmt"
	"math"
	"strings"

	commitMessage "github.com/JaquesBoeno/GitCommitHook/commit"
	"github.com/JaquesBoeno/GitCommitHook/config"
	tea "github.com/charmbracelet/bubbletea"
)

func selectUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
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

	return m, nil
}

func selectView(m Model) string {
	str := strings.Builder{}

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

	return str.String()
}
