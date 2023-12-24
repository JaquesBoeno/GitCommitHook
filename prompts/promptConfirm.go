package prompts

import (
	"fmt"
	"strings"

	commitMessage "github.com/JaquesBoeno/GitCommitHook/commit"
	tea "github.com/charmbracelet/bubbletea"
)

func confirmUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "n", "N":
			m.confirmResponse = msg.String()
			m.pastResponses = fmt.Sprint(m.pastResponses, fmt.Sprintf("\033[1m%s\033[0m \033[30m(y/N)\033[0m %s\n", m.currentQuestion.Label, m.confirmResponse))

			for _, question := range m.currentQuestion.Questions {
				m.Responses = append(m.Responses, commitMessage.Value{
					Id:    question.Id,
					Value: "",
				})
			}

			return m.NextQuestion()
		case "y", "Y":
			m.confirmResponse = msg.String()
			m.pastResponses = fmt.Sprint(m.pastResponses, fmt.Sprintf("\033[1m%s\033[0m \033[30m(y/N)\033[0m %s\n", m.currentQuestion.Label, m.confirmResponse))
			m.Questions = append(m.Questions[:m.index+1], m.Questions[m.index:]...)
			m.Questions[m.index+1] = m.currentQuestion.Questions[0]
			return m.NextQuestion()
		}
	}

	return m, nil
}

func confirmView(m Model) string {
	str := strings.Builder{}
	str.WriteString(fmt.Sprintf("%s \033[30m(y/N)\033[0m %s", m.currentQuestion.Label, m.confirmResponse))

	return str.String()
}
