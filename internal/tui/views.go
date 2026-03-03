package tui

func (m Model) View() string {
	return m.Message + "\n\nPress Ctrl+C to quit.\n"
}
