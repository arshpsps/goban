package bblt

import (
	// "github.com/charmbracelet/bubbles/list"
	"fmt"
	"goban/internals/dataHandle"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var data dataHandle.JsonData = dataHandle.GrabJsonObj()

type model struct {
	projectList []dataHandle.Project
	project     dataHandle.Project
	cursor      int
}

type projModel struct {
	boardList []dataHandle.Board
	project   dataHandle.Project
	cursor    int
}

type boardModel struct {
	cardList []dataHandle.Card
	board    dataHandle.Board
	cursot   int
}

func initialModel() model {
	return model{
		projectList: data.Projects,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m projModel) Init() tea.Cmd {
	return nil
}

func (m boardModel) Init() tea.Cmd {
	return nil
}

func (m projModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.cursor < len(m.boardList)-1 {
				m.cursor++
			}
		case "enter", " ":
			for _, p := range m.boardList {
				if m.boardList[m.cursor].Name == p.Name {

					n := boardModel{
						board:    p,
						cardList: p.Cards,
					}
					return n, nil
				}
			}
		}
	}
	return m, nil
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
			if m.cursor < len(m.projectList)-1 {
				m.cursor++
			}
		case "enter", " ":
			for _, p := range data.Projects {
				if m.projectList[m.cursor].Name == p.Name {

					n := projModel{
						project:   p,
						boardList: p.Boards,
					}
					return n, nil
				}
			}
		}
	}
	return m, nil
}

func (m boardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (n boardModel) View() string {
	s := ""
	return s
}

func (n projModel) View() string {
	s := "Which option to select, uwu?\n\n"
	for i, choice := range n.boardList {
		cursor := " "
		if n.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice.Name)
	}

	s += "\nPress q to quit.\n"

	return s
}

func (m model) View() string {
	s := "Which option to select, uwu?\n\n"
	for i, choice := range m.projectList {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice.Name)
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
