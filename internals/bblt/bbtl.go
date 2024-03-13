package bblt

import (
	// "github.com/charmbracelet/bubbles/list"
	"fmt"
	"goban/internals/dataHandle"
	"goban/internals/kanban"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	data     dataHandle.JsonData = kanban.GrabJsonObj()
	selcProj dataHandle.Project
	selcView int = 0
)

type model struct {
	choices []dataHandle.Project
	cursor  int
}

func initialModel() model {
	return model{
		choices: data.Projects,
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

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			for _, p := range data.Projects {
				if m.choices[m.cursor].Name == p.Name {
					selcProj = p
					selcView = 0
					break
				}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Which option to select, uwu?\n\n"
	switch selcView {

	case 0:
		for i, choice := range m.choices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			s += fmt.Sprintf("%s %s\n", cursor, choice.Name)
		}
	case 1:

	}

	s += "\nPress q to quit.\n"

	return s
}

func Run() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("err, you fucked up: %v", err)
		os.Exit(1)
	}
}
