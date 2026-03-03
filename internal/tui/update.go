package tui

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				return m, nil
			}
			return m, nil
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}
