package tui

import (
	"os"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/compat"
	"github.com/charmbracelet/x/term"
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

var DefaultStyle = lipgloss.NewStyle().
	Foreground(compat.AdaptiveColor{Light: lipgloss.Color(FrgMagenta), Dark: lipgloss.Color(FrgLime)})

// IsAccessible reports whether accessible mode should be enabled.
// Returns true if the ACCESSIBLE env var is set or stdin is not a terminal.
func IsAccessible() bool {
	return os.Getenv("ACCESSIBLE") != "" || !term.IsTerminal(os.Stdin.Fd())
}
