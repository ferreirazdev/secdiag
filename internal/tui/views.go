package tui

import "strings"

func (m Model) View() string {
	var b strings.Builder
	b.WriteString(m.Message + "\n\n")
	for i, opt := range m.Options {
		if i == m.Selected {
			b.WriteString("▸ " + opt + "\n")
		} else {
			b.WriteString("  " + opt + "\n")
		}
	}

	b.WriteString("\n↑/↓ or j/k move, Enter select, q quit\n")
	return b.String()
}
