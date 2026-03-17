package tui

import (
	"os"
	"sync"

	"charm.land/lipgloss/v2"
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

// DefaultStyle returns the FRG accent style adapted to the terminal background.
func DefaultStyle(isDark bool) lipgloss.Style {
	accent := lipgloss.LightDark(isDark)(lipgloss.Color(FrgMagenta), lipgloss.Color(FrgLime))
	return lipgloss.NewStyle().Foreground(accent)
}

// IsAccessible reports whether accessible mode should be enabled.
// Returns true if the ACCESSIBLE env var is set or stdin is not a terminal.
func IsAccessible() bool {
	return os.Getenv("ACCESSIBLE") != "" || !term.IsTerminal(os.Stdin.Fd())
}

// GlamourStyle returns the glamour style name appropriate for the terminal.
// It respects the GLAMOUR_STYLE env var, auto-detects the terminal background
// on TTYs, and defaults to "dark" for non-TTY environments (CI, pipes).
// The result is cached after the first call because HasDarkBackground queries
// the terminal, which hangs after a bubbletea program has run.
func GlamourStyle() string {
	glamourStyleOnce.Do(func() {
		if style := os.Getenv("GLAMOUR_STYLE"); style != "" {
			glamourStyleValue = style
			return
		}
		if term.IsTerminal(os.Stdout.Fd()) && !lipgloss.HasDarkBackground(os.Stdin, os.Stdout) {
			glamourStyleValue = "light"
			return
		}
		glamourStyleValue = "dark"
	})
	return glamourStyleValue
}

var (
	glamourStyleOnce  sync.Once
	glamourStyleValue string
)

