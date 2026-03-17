package tui

import (
	"os"

	"charm.land/lipgloss/v2"
)

// FRG brand colors for consistent theming across applications
const (
	FrgLime      = "#93C30B"
	FrgMagenta   = "#BD368D"
	FrgForest    = "#004610"
	FrgMint      = "#DBEAE5"
	FrgMaroon    = "#4B0325"
	FrgBlue      = "#244A66"
	FrgLightGray = "#F5F5F5"
	FrgGray      = "#A6A6A6"
	FrgDarkGray  = "#4B4B4B"
	FrgWhite     = "#FFFFFF"
	FrgBlack     = "#000000"
)

// IsDark reports whether the terminal has a dark background.
var IsDark = lipgloss.HasDarkBackground(os.Stdin, os.Stdout)

var DefaultStyle = lipgloss.NewStyle().
	Foreground(lipgloss.LightDark(IsDark)(lipgloss.Color(FrgMagenta), lipgloss.Color(FrgLime)))
