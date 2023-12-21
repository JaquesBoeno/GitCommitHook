package prompts

import (
	"fmt"
	"strings"

	"github.com/JaquesBoeno/GitHook/config"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choice config.Item // items on the to-do list
	cursor int         // which to-do list item our cursor is pointing at
}

var choices = []config.Item{
	{Name: "front-end", Desc: "Change in front-end scope"},
	{Name: "back-end", Desc: "Change in back-end scope"},
}

func InitialModel() model {
	return model{
		cursor: 0,
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
			m.choice = choices[m.cursor]
			return m, tea.Quit

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}

		case "down", "j":
			m.cursor++
			if m.cursor >= len(choices) {
				m.cursor = 0
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := strings.Builder{}
	s.WriteString("What scope of this change? (e.g. backend or frontend)\n\n")

	for i, _ := range choices {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}

		s.WriteString(choices[i].Name)
		s.WriteString("\n")
	}

	s.WriteString(fmt.Sprintf("\n%s\n", choices[m.cursor].Desc))

	s.WriteString("\n(press q to quit)\n")

	return s.String()
}
