package config

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
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
	title := titleStyle.Render(logo)
	subtitle := subtitleStyle.Render("Select language for comment removal:")

	var listItems strings.Builder
	for i, choice := range m.languageChoices {
		if m.cursor == i {
			listItems.WriteString(selectedItemStyle.Render(choice) + "\n")
		} else {
			listItems.WriteString(listItemStyle.Render(choice) + "\n")
		}
	}

	mutedInstructionStyle := buttonStyle
	mutedInstructionStyle = mutedInstructionStyle.
		Foreground(subtle).
		Background(lipgloss.NoColor{}).
		Bold(false)

	instruction := mutedInstructionStyle.Render("[ Enter ] to select")
	quitInstruction := mutedInstructionStyle.Render("[ q ] to quit")

	instructions := lipgloss.JoinHorizontal(
		lipgloss.Center,
		instruction,
		"    ",
		quitInstruction,
	)

	return appStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			subtitle,
			listItems.String(),
			"",
			instructions,
		),
	)
}

func (m Model) renderFormatSelect() string {
	title := titleStyle.Render(logo)

	langInfo := fmt.Sprintf("Selected language: %s\n",
		highlightedInfoStyle.Render(m.config.Language))

	subtitle := subtitleStyle.Render(
		"Do you want to autoformat the code?")

	var listItems strings.Builder
	for i, choice := range m.formatChoices {
		if m.cursor == i {
			listItems.WriteString(selectedItemStyle.Render(choice) + "\n")
		} else {
			listItems.WriteString(listItemStyle.Render(choice) + "\n")
		}
	}

	mutedInstructionStyle := buttonStyle
	mutedInstructionStyle = mutedInstructionStyle.
		Foreground(subtle).
		Background(lipgloss.NoColor{}).
		Bold(false)

	enterInstruction := mutedInstructionStyle.Render("[ Enter ] to select")
	backInstruction := mutedInstructionStyle.Render("[ Backspace ] to go back")
	quitInstruction := mutedInstructionStyle.Render("[ q ] to quit")

	instructions := lipgloss.JoinHorizontal(
		lipgloss.Center,
		enterInstruction,
		"    ",
		backInstruction,
		"    ",
		quitInstruction,
	)

	return appStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			langInfo,
			subtitle,
			listItems.String(),
			"",
			instructions,
		),
	)
}

func (m Model) renderMonitoring() string {
	title := titleStyle.Render(logo)
	subtitle := subtitleStyle.Render("Monitoring clipboard with configuration:")

	langInfo := infoStyle.Render(fmt.Sprintf("Language: %s",
		highlightedInfoStyle.Render(m.config.Language)))

	formatInfo := infoStyle.Render(fmt.Sprintf("Autoformat: %s",
		highlightedInfoStyle.Render(fmt.Sprintf("%v", m.config.Format))))

	var logSection string
	if len(m.outputs) == 0 {
		logSection = infoStyle.Render("Waiting for clipboard content...")
	} else {
		logTitle := infoStyle.Render("Activity log:")

		var logEntries strings.Builder
		startIdx := 0
		if len(m.outputs) > 10 {
			startIdx = len(m.outputs) - 10
		}

		for _, output := range m.outputs[startIdx:] {
			var style lipgloss.Style
			if strings.Contains(output, "Error") {
				style = errorLogStyle
			} else {
				style = successLogStyle
			}
			logEntries.WriteString(style.Render("• "+output) + "\n")
		}

		logSection = lipgloss.JoinVertical(
			lipgloss.Left,
			logTitle,
			logEntries.String(),
		)
	}

	mutedInstructionStyle := buttonStyle
	mutedInstructionStyle = mutedInstructionStyle.
		Foreground(subtle).
		Background(lipgloss.NoColor{}).
		Bold(false)

	settingsInstruction := mutedInstructionStyle.Render("[ s ] to change settings")
	quitInstruction := mutedInstructionStyle.Render("[ q ] to quit")

	instructions := lipgloss.JoinHorizontal(
		lipgloss.Center,
		settingsInstruction,
		"    ",
		quitInstruction,
	)

	return appStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			subtitle,
			lipgloss.JoinVertical(
				lipgloss.Left,
				langInfo,
				formatInfo,
			),
			"",
			logSection,
			"",
			instructions,
		),
	)
}
