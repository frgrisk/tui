package tui

import "github.com/charmbracelet/lipgloss"

const frgLime = "#93C30B"

const frgMagenta = "#BD368D"

var DefaultStyle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: frgMagenta, Dark: frgLime})
