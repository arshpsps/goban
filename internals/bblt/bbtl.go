package bblt

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct{
    choices []string
    cursor int
    selected map[int]struct{}
}

func initialModel() model {
    return model {
        choices: []string{"item1", "item2"},
        selected: make(map[int]struct{}),
    }
}

func (m model) init() tea.Cmd {
    return nil
}
