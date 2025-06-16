package tui

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

// FormTheme returns a huh.Theme configured with the project color palette.
func FormTheme() *huh.Theme {
	t := huh.ThemeBase()

	// Define adaptive colors based on the palette. Where possible a slightly
	// brighter variant is used in dark mode to improve contrast.
	var (
		accent = lipgloss.AdaptiveColor{Light: FrgMagenta, Dark: FrgLime}
		white  = lipgloss.AdaptiveColor{Light: FrgDarkGray, Dark: FrgLightGray}
		gray   = lipgloss.AdaptiveColor{Light: FrgDarkGray, Dark: FrgGray}
		mint   = lipgloss.AdaptiveColor{Light: FrgMint, Dark: FrgMint}
		red    = lipgloss.Color("#FF0000")
	)

	// Group styles.
	t.Group.Title = t.Group.Title.Foreground(accent).Bold(true)

	// Focused field styles.
	t.Focused.Title = t.Focused.Title.Foreground(accent).Bold(true)
	t.Focused.NoteTitle = t.Focused.NoteTitle.Background(mint).Foreground(accent).Bold(true)
	t.Focused.Description = t.Focused.Description.Foreground(white)
	t.Focused.SelectSelector = t.Focused.SelectSelector.Foreground(accent)
	t.Focused.SelectedOption = t.Focused.SelectedOption.Foreground(accent)
	t.Focused.MultiSelectSelector = t.Focused.MultiSelectSelector.Foreground(accent)
	t.Focused.FocusedButton = t.Focused.FocusedButton.Foreground(lipgloss.Color(FrgBlack)).Background(accent).Bold(true)
	t.Focused.TextInput.Prompt = t.Focused.TextInput.Prompt.Foreground(accent)
	t.Focused.TextInput.Cursor = t.Focused.TextInput.Cursor.Foreground(accent)
	t.Focused.ErrorMessage = t.Focused.ErrorMessage.Foreground(red)
	t.Focused.ErrorIndicator = t.Focused.ErrorIndicator.Foreground(red)

	// Blurred field styles.
	t.Blurred.Title = t.Blurred.Title.Foreground(gray)
	t.Blurred.NoteTitle = t.Blurred.NoteTitle.Background(mint).Foreground(accent).Bold(true)
	t.Blurred.Description = t.Blurred.Description.Foreground(gray)
	t.Blurred.TextInput.Prompt = t.Blurred.TextInput.Prompt.Foreground(gray)
	t.Blurred.Option = t.Blurred.Option.Foreground(gray)
	t.Blurred.UnselectedOption = t.Blurred.UnselectedOption.Foreground(gray)
	t.Blurred.UnselectedPrefix = t.Blurred.UnselectedPrefix.Foreground(gray)
	t.Blurred.SelectedOption = t.Blurred.SelectedOption.Foreground(gray)
	t.Blurred.SelectedPrefix = t.Blurred.SelectedPrefix.Foreground(gray)

	// Help styles.
	t.Help.ShortKey = t.Help.ShortKey.Foreground(accent)
	t.Help.FullKey = t.Help.FullKey.Foreground(accent)
	t.Help.ShortDesc = t.Help.ShortDesc.Foreground(gray)
	t.Help.FullDesc = t.Help.FullDesc.Foreground(gray)

	return t
}
