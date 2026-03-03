package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.Screen == screenURLInput {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter", "return":
				url := strings.TrimSpace(m.URLInput.Value())
				if url != "" {
					m.ScanURL = url
					m.Screen = screenResults
					m.URLInput.Reset()
				}
				return m, nil
			case "esc":
				m.Screen = screenMenu
				m.URLInput.Reset()
				return m, nil
			}
		}
		var cmd tea.Cmd
		m.URLInput, cmd = m.URLInput.Update(msg)
		return m, cmd
	}

	// Results screen: Enter or B goes back to menu
	if m.Screen == screenResults {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "enter" || msg.String() == "return" || msg.String() == "b" {
				m.Screen = screenMenu
				return m, nil
			}
			if msg.String() == "ctrl+c" || msg.String() == "q" {
				return m, tea.Quit
			}
		}
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.Selected > 0 {
				m.Selected--
			}
			return m, nil
		case "down", "j":
			if m.Selected < len(m.Options)-1 {
				m.Selected++
			}
			return m, nil
		case "enter", "return":
			switch m.Options[m.Selected] {
			case "Quit":
				return m, tea.Quit
			case "Scan URL":
				m.Screen = screenURLInput
				m.URLInput.Focus()
				return m, textinput.Blink
			}
			return m, nil
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}
