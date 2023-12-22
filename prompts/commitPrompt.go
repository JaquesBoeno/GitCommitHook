package prompts

import (
	"fmt"
	"math"
	"strings"

	"github.com/JaquesBoeno/GitHook/config"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Questions []config.Question
	Responses []response

	currentTextinput    textinput.Model
	currentCursor       int
	currentChoice       config.Option
	currentQuestion     config.Question
	currentQuestionType string
	pastResponses       string
	index               int
}

type response struct {
	id    string
	value string
}

func InitialModel(questions []config.Question) Model {
	qType := ""
	if questions[0].Options[0].Name != "" {
		qType = "select"
	} else {
		qType = "text"
	}

	return Model{
		Questions:           questions,
		currentQuestionType: qType,
		currentQuestion:     questions[0],
		currentCursor:       0,
		index:               0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.currentQuestionType = "select"

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	}

	switch m.currentQuestionType {
	case "select":
		{
			switch msg := msg.(type) {
			case tea.KeyMsg:
				switch msg.String() {
				case "enter":
					m.Responses = append(m.Responses, response{
						id:    m.currentQuestion.Id,
						value: m.currentQuestion.Options[m.currentCursor].Name,
					})
					m.pastResponses = fmt.Sprint(m.pastResponses, fmt.Sprintf("%s: \033[1m%s\033[0m\n",
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
	}

	return m, nil
}

func (m Model) View() string {
	str := strings.Builder{}
	str.WriteString(m.pastResponses)
	// Question Type: selected
	switch m.currentQuestionType {
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

			for i, option := range questionInPage {
				if m.currentCursor == i+start {
					str.WriteString("(•) ")
				} else {
					str.WriteString("( ) ")
				}

				str.WriteString(option.Name)
				str.WriteString("\n")
			}

			str.WriteString(fmt.Sprintf("\n%s\n", m.currentQuestion.Options[m.currentCursor].Desc))
		}
	}

	return str.String()
}

func (m Model) NextQuestion() (tea.Model, tea.Cmd) {
	if m.index+1 >= len(m.Questions) {
		m.View()

		return m, nil
	} else {
		m.index++
		m.currentQuestion = m.Questions[m.index]
		qType := ""

		if m.currentQuestion.Options[0].Name != "" {
			qType = "select"
		} else {
			qType = "text"
		}

		m.currentQuestionType = qType
		return m, nil
	}

}
