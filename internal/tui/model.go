package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type screen int

const (
	screenMenu screen = iota
	screenURLInput
	screenResults
)

type Model struct {
	Message  string
	Options  []string
	Selected int
	Screen   screen
	URLInput textinput.Model
	ScanURL  string
}

func newURLInput() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "https://example.com"
	ti.Width = 60
	return ti
}

func NewModel() Model {
	return Model{
		Message:  "secdiag — security diagnostics",
		Options:  []string{"Scan URL", "TLS Check", "Headers Check", "Quit"},
		Selected: 0,
		URLInput: newURLInput(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
