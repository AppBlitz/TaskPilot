package handlers

import (
	"github.com/charmbracelet/lipgloss"
)

var Style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#3C3C3C")).
	PaddingTop(1).
	PaddingLeft(4).
	Width(60)
