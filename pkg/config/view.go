package config

import (
	"fmt"
)

func (m Model) View() string {
	switch m.screen {
	case languageSelect:
		return m.renderLanguageSelect()
	case formatSelect:
		return m.renderFormatSelect()
	case monitoring:
		return m.renderMonitoring()
	default:
		return "Unknown state"
	}
}

func (m Model) renderLanguageSelect() string {
	s := "\nSelect language for comment removal:\n\n"

	for i, choice := range m.languageChoices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress enter to select"
	return s
}

func (m Model) renderFormatSelect() string {
	s := fmt.Sprintf("\nSelected language: %s\n\n", m.config.Language)
	s += "Do you want to autoformat the code? (only works with golang currently)\n\n"

	for i, choice := range m.formatChoices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress enter to select"
	s += "\nPress backspace to go back"
	return s
}

func (m Model) renderMonitoring() string {
	s := "\nMonitoring clipboard with configuration:\n"
	s += fmt.Sprintf("\nLanguage: %s\n", m.config.Language)
	s += fmt.Sprintf("Autoformat: %v\n\n", m.config.Format)

	if len(m.outputs) == 0 {
		s += "Waiting for clipboard content...\n"
	} else {
		s += "Activity log:\n"

		startIdx := 0
		if len(m.outputs) > 10 {
			startIdx = len(m.outputs) - 10
		}

		for _, output := range m.outputs[startIdx:] {
			s += fmt.Sprintf("- %s\n", output)
		}
	}

	s += "\nPress s to change settings"
	s += "\nPress q to quit"
	return s
}
