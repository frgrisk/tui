package tui

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

// FormTheme returns a huh.Theme configured with the project colour palette.
func FormTheme() *huh.Theme {
	t := huh.ThemeBase()

	// Define adaptive colours based on the palette. Where possible a slightly
	// brighter variant is used in dark mode to improve contrast.
	var (
		lime        = lipgloss.AdaptiveColor{Light: frgLime, Dark: frgLimeDark}
		magenta     = lipgloss.AdaptiveColor{Light: frgMagenta, Dark: frgMagenta}
		forestGreen = lipgloss.AdaptiveColor{Light: frgForest, Dark: frgForest}
		white       = lipgloss.AdaptiveColor{Light: frgWhite, Dark: frgLightGray}
		gray        = lipgloss.AdaptiveColor{Light: frgGray, Dark: frgDarkGray}
		mint        = lipgloss.AdaptiveColor{Light: frgMint, Dark: frgMint}
		maroon      = lipgloss.AdaptiveColor{Light: frgMaroon, Dark: frgMaroon}
	)

	// Group styles.
	t.Group.Title = t.Group.Title.Foreground(magenta).Bold(true)

	// Focused field styles.
	t.Focused.Title = t.Focused.Title.Foreground(forestGreen).Bold(true)
	t.Focused.NoteTitle = t.Focused.NoteTitle.Background(mint).Foreground(magenta).Bold(true)
	t.Focused.Description = t.Focused.Description.Foreground(white)
	t.Focused.SelectSelector = t.Focused.SelectSelector.Foreground(magenta)
	t.Focused.SelectedOption = t.Focused.SelectedOption.Foreground(lime)
	t.Focused.MultiSelectSelector = t.Focused.MultiSelectSelector.Foreground(magenta)
	t.Focused.FocusedButton = t.Focused.FocusedButton.Foreground(lipgloss.Color(frgBlack)).Background(lime).Bold(true)
	t.Focused.TextInput.Prompt = t.Focused.TextInput.Prompt.Foreground(lime)
	t.Focused.TextInput.Cursor = t.Focused.TextInput.Cursor.Foreground(magenta)
	t.Focused.ErrorMessage = t.Focused.ErrorMessage.Foreground(maroon)
	t.Focused.ErrorIndicator = t.Focused.ErrorIndicator.Foreground(maroon)

	// Blurred field styles.
	t.Blurred.Title = t.Blurred.Title.Foreground(gray)
	t.Blurred.NoteTitle = t.Blurred.NoteTitle.Background(mint).Foreground(magenta).Bold(true)
	t.Blurred.Description = t.Blurred.Description.Foreground(gray)
	t.Blurred.TextInput.Prompt = t.Blurred.TextInput.Prompt.Foreground(gray)

	// Help styles.
	t.Help.ShortKey = t.Help.ShortKey.Foreground(lime)
	t.Help.FullKey = t.Help.FullKey.Foreground(lime)

	return t
}
