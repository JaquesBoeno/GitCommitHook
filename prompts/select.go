package prompts

import (
	"fmt"
	"math"
	"strings"

	"github.com/JaquesBoeno/GitHook/config"
	tea "github.com/charmbracelet/bubbletea"
)

type SelectModel struct {
	Choice   config.Item
	cursor   int
	Question config.Question
}

func InitialModel(question config.Question) SelectModel {
	return SelectModel{
		cursor:   0,
		Question: question,
	}
}

func (m SelectModel) Init() tea.Cmd {
	return nil
}

func (m SelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "enter":
			m.Choice = m.Question.Options[m.cursor]
			return m, tea.Quit

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.Question.Options) - 1
			}

		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.Question.Options) {
				m.cursor = 0
			}
		}
	}

	return m, nil
}

func (m SelectModel) View() string {
	s := strings.Builder{}

	questionPerPage := 5 // odd number only
	questionInPager := make([]config.Item, questionPerPage)

	if len(m.Choice.Name) != 0 {
		// \033[1 is ANSI code for make text bold, check FormattingSheet.md in root for other styles
		// enter here when have a choice selected
		s.WriteString(fmt.Sprintf("%s: \033[1m%s\033[0m\n\n", m.Question.Label, m.Choice.Name))
	} else {
		s.WriteString(fmt.Sprintf("%s:\n\n", m.Question.Label))
		if len(m.Question.Options) > questionPerPage {
			half := int(math.Floor(float64(questionPerPage) / 2))
			start, end := 0, 0
			start = m.cursor - half
			end = m.cursor + half + 1

			if start < 0 {
				start = 0
				end = questionPerPage
			} else if end > len(m.Question.Options) {
				end = len(m.Question.Options)
				start = len(m.Question.Options) - questionPerPage
			}

			questionInPager = m.Question.Options[start:end]

			showChoices(&s, m.cursor, questionInPager, start)

		} else {
			showChoices(&s, m.cursor, m.Question.Options, 0)
		}

		s.WriteString(fmt.Sprintf("\n%s\n", m.Question.Options[m.cursor].Desc))

	}

	return s.String()
}

func showChoices(s *strings.Builder, cursor int, options []config.Item, start int) {
	for i := range options {
		if cursor == i+start {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}

		s.WriteString(options[i].Name)
		s.WriteString("\n")
	}
}
