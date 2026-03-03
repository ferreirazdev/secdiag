package tui

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	Message string
}

func NewModel() Model {
	return Model{
		Message: "secdiag — security diagnostics",
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
