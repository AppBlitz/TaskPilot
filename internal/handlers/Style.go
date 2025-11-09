package handlers

import (
	"github.com/charmbracelet/lipgloss"
)

func Styles(message string, argsColar []string) *lipgloss.Style {
	Style := lipgloss.NewStyle().
		SetString(message).
		BorderStyle(lipgloss.InnerHalfBlockBorder()).
		BorderForeground(lipgloss.Color(argsColar[0])).
		Bold(true).
		Foreground(lipgloss.Color(argsColar[1])).
		Background(lipgloss.Color(argsColar[2])).
		PaddingTop(1).
		PaddingLeft(4).
		Width(60)
	return &Style
}
