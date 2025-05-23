package tui

import "github.com/charmbracelet/lipgloss"

const (
	frgLime      = "#93C30B"
	frgLimeDark  = "#2EC272"
	frgMagenta   = "#BD368D"
	frgForest    = "#004610"
	frgMint      = "#DBEAE5"
	frgMaroon    = "#4B0325"
	frgBlue      = "#244A66"
	frgLightGray = "#F5F5F5"
	frgGray      = "#A6A6A6"
	frgDarkGray  = "#4B4B4B"
	frgWhite     = "#FFFFFF"
	frgBlack     = "#000000"
)

var DefaultStyle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: frgMagenta, Dark: frgLime})
