package bblt

import (
	// "github.com/charmbracelet/bubbles/list"
	"fmt"
	"goban/internals/dataHandle"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var db dataHandle.DataHandler
var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type model struct {
	projectList []dataHandle.Project
	project     dataHandle.Project
	cursor      int
}

type projModel struct {
	boardList []dataHandle.Board
	project   dataHandle.Project
	model     model
	cursor    int
}

type boardModel struct {
	cardList []dataHandle.Card
	board    dataHandle.Board
	project  projModel
	cursor   int
}

type cardModel struct {
	inputs     []textinput.Model
	card       dataHandle.Card
	board      boardModel
	focusIndex int
	cursorMode cursor.Mode
}

type createProjModel struct {
	newModel   interface{}
	rootModel  tea.Model
	inputs     []textinput.Model
	focusIndex int
	cursorMode cursor.Mode
}

// TODO: generic bblt methods (update, view, etc) for createBoardModel and createProjModel
type createBoardModel struct {
	inputs     []textinput.Model
	newModel   dataHandle.Board
	rootModel  projModel
	focusIndex int
	cursorMode cursor.Mode
}

func initialModel() model {
	return model{
		projectList: db.GetAllProjects(),
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

func (m cardModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m createProjModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m createBoardModel) Init() tea.Cmd {
	return textinput.Blink
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
			if m.cursor < len(m.boardList) {
				m.cursor++
			}
		case "enter", " ":
			if m.cursor < len(m.boardList) {
				for _, b := range m.boardList {
					if m.boardList[m.cursor].Name == b.Name {

						n := boardModel{
							board:    b,
							project:  m,
							cardList: db.GetCardsInBoard(b.ID),
						}
						return n, nil
					}
				}
			} else if m.cursor == len(m.boardList) {
				n := createProjModel{
					inputs:    make([]textinput.Model, 1),
					rootModel: m,
				}

				var t textinput.Model
				for i := range n.inputs {
					t = textinput.New()
					t.Cursor.Style = cursorStyle
					t.CharLimit = 32
					switch i {
					case 0:
						t.Placeholder = "Board Title"
						t.Focus()
						t.PromptStyle = focusedStyle
						t.TextStyle = focusedStyle

						n.inputs[i] = t
					}
				}

				return n, nil

			}
		case "esc":
			return m.model, nil
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
			if m.cursor < len(m.projectList) { // -1 on the length if removing the input wala TextBox
				m.cursor++
			}
		case "enter", " ":
			if m.cursor < len(m.projectList) {
				for _, p := range m.projectList {
					if m.projectList[m.cursor].Name == p.Name {

						n := projModel{
							project:   p,
							model:     m,
							boardList: db.GetBoardsInProject(p.ID),
						}
						return n, nil
					}
				}
			} else if m.cursor == len(m.projectList) {
				n := createProjModel{
					inputs:    make([]textinput.Model, 1),
					rootModel: m,
				}

				var t textinput.Model
				for i := range n.inputs {
					t = textinput.New()
					t.Cursor.Style = cursorStyle
					t.CharLimit = 32
					switch i {
					case 0:
						t.Placeholder = "Project Title"
						t.Focus()
						t.PromptStyle = focusedStyle
						t.TextStyle = focusedStyle

						n.inputs[i] = t
					}
				}

				return n, nil

			}

		}
	}
	return m, nil
}

func (m boardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.cursor < len(m.cardList)-1 {
				m.cursor++
			}
		case "enter", " ":
			for _, p := range m.cardList {
				if m.cardList[m.cursor].Title == p.Title {

					n := cardModel{
						inputs: make([]textinput.Model, 2),
						board:  m,
						card:   p,
					}

					var t textinput.Model
					for i := range n.inputs {
						t = textinput.New()
						t.Cursor.Style = cursorStyle
						t.CharLimit = 32
						switch i {
						case 0:
							t.SetValue(n.card.Title)
							t.Focus()
							t.PromptStyle = focusedStyle
							t.TextStyle = focusedStyle
						case 1:
							t.SetValue(n.card.Description)
							t.CharLimit = 64
						}

						n.inputs[i] = t
					}

					return n, nil
				}
			}
		case "esc":
			return m.project, nil
		}
	}
	return m, nil
}

func inputsToStrs(m []textinput.Model) []string {
	var strs []string
	for _, i := range m {
		strs = append(strs, i.Value())
	}
	return strs
}

func (m cardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
			// Change cursor mode

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusIndex == len(m.inputs) {
				inps := inputsToStrs(m.inputs)
				if len(inps[0]) > 0 {
					m.card.Title = inps[0]
				}
				if len(inps[1]) > 0 {
					m.card.Description = inps[1]
				}

				db.UpdateCard(m.card)
				m.board.cardList = db.GetCardsInBoard(m.board.board.ID)

				return m.board, nil
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)

		case "esc":
			return m.board, nil
		}
	}
	// Handle character input and blinking
	cmd := m.updateInputs(msg)
	return m, cmd
}

func (m createProjModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
			// Change cursor mode

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusIndex == len(m.inputs) {
				inps := inputsToStrs(m.inputs)
				switch m.rootModel.(type) {
				case model:
					{
						proj := dataHandle.Project{}
						if len(inps[0]) > 0 {
							proj.Name = inps[0]
						}

						db.Insert(&proj)
						m.newModel = proj
						root := m.rootModel.(model)
						root.projectList = db.GetAllProjects()
						m.rootModel = root
					}

				case projModel:
					{
						board := dataHandle.Board{}
						if len(inps[0]) > 0 {
							board.Name = inps[0]
						}

						root := m.rootModel.(projModel)

						board.ProjectID = root.project.ID
						db.InsertBoard(&board)
						m.newModel = board

						root.boardList = db.GetBoardsInProject(root.project.ID)
						m.rootModel = root
					}
				case boardModel:
					{
						card := dataHandle.Card{}
						if len(inps[0]) > 0 {
							card.Title = inps[0]
							card.Description = inps[1]
						}

						root := m.rootModel.(boardModel)

						card.BoardID = root.board.ID
						db.InsertCard(&card)
						m.newModel = card

						root.cardList = db.GetCardsInBoard(root.board.ID)
						m.rootModel = root
					}

				default:
					fmt.Fprintln(os.Stderr, "sorry")
				}

				return m.rootModel, nil
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)

		case "esc":
			return m.rootModel, nil
		}
	}
	// Handle character input and blinking
	cmd := m.updateInputs(msg)
	return m, cmd
}

func (m *cardModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m *createProjModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (n boardModel) View() string {
	s := "Which option to select, uwu?\n\n"
	for i, choice := range n.cardList {
		cursor := " "
		if n.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice.Title)
	}

	s += "\nPress q to quit.\n"

	return s
}

func (n projModel) View() string {
	cursor := " "
	s := "Which option to select, uwu?\n\n"
	for i, choice := range n.boardList {
		cursor = " "
		if n.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice.Name)
	}

	cursor = " "
	if n.cursor == len(n.boardList) {
		cursor = ">"
	}
	s += fmt.Sprintf("%s %s\n", cursor, "[ Create New Board]")

	s += "\nPress q to quit.\n"

	return s
}

func (m model) View() string {
	cursor := " "
	s := "Which option to select, uwu?\n\n"
	for i, choice := range m.projectList {
		cursor = " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice.Name)
	}
	cursor = " "
	if m.cursor == len(m.projectList) {
		cursor = ">"
	}
	s += fmt.Sprintf("%s %s\n", cursor, "[ Create New Project ]")

	s += "\nPress q to quit.\n"

	return s
}

func (m cardModel) View() string { // works actually
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return b.String()
}

func (m createProjModel) View() string { // works actually
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return b.String()
}

func Run() {
	dbConn := dataHandle.NewSqliteConndb()
	db = &dbConn
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("err, you fucked up: %v", err)
		os.Exit(1)
	}
}
