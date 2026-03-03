package tui

import "strings"

func (m Model) View() string {
	switch m.Screen {
	case screenURLInput:
		return viewURLInput(m)
	case screenResults:
		return viewResults(m)
	default:
		return viewMenu(m)
	}
}

func viewMenu(m Model) string {
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

func viewURLInput(m Model) string {
	var b strings.Builder
	b.WriteString("Enter URL to scan\n\n")
	b.WriteString(m.URLInput.View())
	b.WriteString("\n\nEnter = scan · Esc = back\n")
	return b.String()
}

func viewResults(m Model) string {
	var b strings.Builder
	b.WriteString("Scan requested\n\n")
	b.WriteString("  URL: " + m.ScanURL + "\n\n")
	b.WriteString("(Placeholder — analyzers not implemented yet)\n\n")
	b.WriteString("Enter or B = back to menu · q = quit\n")
	return b.String()
}
