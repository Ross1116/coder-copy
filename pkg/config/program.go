package config

import (
	tea "github.com/charmbracelet/bubbletea"
)

func NewProgram(processContentFn func(string, string, bool) (string, error)) *tea.Program {
	return tea.NewProgram(initialModel(processContentFn))
}

func (m Model) Init() tea.Cmd {
	return nil
}
