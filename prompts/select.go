package prompts

import (
	"fmt"
	"math"
	"strings"

	"github.com/JaquesBoeno/GitHook/config"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choice   config.Item // items on the to-do list
	cursor   int         // which to-do list item our cursor is pointing at
	question config.Question
}

func InitialModel(question config.Question) model {
	return model{
		cursor:   0,
		question: question,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter":
			m.choice = m.question.Items[m.cursor]
			return m, tea.Quit

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.question.Items) - 1
			}

		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.question.Items) {
				m.cursor = 0
			}
		}
	}

	return m, nil
}

func showChoices(s *strings.Builder, cursor int, items []config.Item, start int) {
	for i := range items {
		if cursor == i+start {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}

		s.WriteString(items[i].Name)
		s.WriteString("\n")
	}
}

func (m model) View() string {
	s := strings.Builder{}

	questionPerPage := 5 // odd number only
	questionInPager := make([]config.Item, questionPerPage)

	if len(m.choice.Name) != 0 {
		// \033[1 is ANSI code for make text bold, check FormattingSheet.md in root for other styles
		// enter here when have a choice selected
		s.WriteString(fmt.Sprintf("%s: \033[1m%s\033[0m\n", m.question.Label, m.choice.Name))
	} else {
		s.WriteString(fmt.Sprintf("%s:\n\n", m.question.Label))
		if len(m.question.Items) > questionPerPage {
			half := int(math.Floor(float64(questionPerPage) / 2))
			start, end := 0, 0
			start = m.cursor - half
			end = m.cursor + half + 1

			if start < 0 {
				start = 0
				end = questionPerPage
			} else if end > len(m.question.Items) {
				end = len(m.question.Items)
				start = len(m.question.Items) - questionPerPage
			}

			questionInPager = m.question.Items[start:end]

			showChoices(&s, m.cursor, questionInPager, start)

		} else {
			showChoices(&s, m.cursor, m.question.Items, 0)
		}

		s.WriteString(fmt.Sprintf("\n%s\n", m.question.Items[m.cursor].Desc))

		s.WriteString("\n(press q to quit)\n")
	}

	return s.String()
}
