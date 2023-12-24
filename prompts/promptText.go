package prompts

import (
	"fmt"
	"strings"

	commitMessage "github.com/JaquesBoeno/GitCommitHook/commit"
	tea "github.com/charmbracelet/bubbletea"
)

func textUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.Err = m.checkLen()

			if m.Err == nil {
				m.pastResponses = fmt.Sprint(m.pastResponses, fmt.Sprintf("\033[1m%s:\033[0m\n\033[32m(%d)\033[0m %s\n",
					m.currentQuestion.Label,
					len(m.currentTextinput.Value()),
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
}

func textView(m Model) string {
	str := strings.Builder{}

	err := m.checkLen()

	if err == nil {
		str.WriteString(fmt.Sprintf(
			"%s\n\033[32m(%d) \033[0m%s\n",
			m.currentQuestion.Label,
			len(m.currentTextinput.Value()),
			m.currentTextinput.View(),
		) + "\n")
	} else {
		str.WriteString(fmt.Sprintf(
			"%s\n\033[31m(%d) \033[0m%s\n",
			m.currentQuestion.Label,
			len(m.currentTextinput.Value()),
			m.currentTextinput.View(),
		) + "\n")
	}

	return str.String()
}
