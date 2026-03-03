package tui

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	Message  string
	Options  []string
	Selected int
}

func NewModel() Model {
	return Model{
		Message:  "secdiag — security diagnostics",
		Options:  []string{"Scan URL", "TLS Check", "Headers Check", "Quit"},
		Selected: 0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
