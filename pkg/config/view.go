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
	case contentView:
		return m.renderContentView()
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

		const maxVisibleLogs = 8

		startIdx := max(0, len(m.outputs)-maxVisibleLogs)
		recentOutputs := m.outputs[startIdx:]

		var logEntries strings.Builder
		for _, output := range recentOutputs {
			var style lipgloss.Style
			if strings.Contains(output, "Error") {
				style = errorLogStyle
			} else if strings.Contains(output, "Warning") {
				style = warningLogStyle
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

	mutedInstructionStyle := buttonStyle.
		Foreground(subtle).
		Background(lipgloss.NoColor{}).
		Bold(false)

	settingsInstruction := mutedInstructionStyle.Render("[ s ] to change settings")
	viewInstruction := mutedInstructionStyle.Render("[ v ] to view last processed content")
	quitInstruction := mutedInstructionStyle.Render("[ q ] to quit")

	instructions := lipgloss.JoinHorizontal(
		lipgloss.Center,
		settingsInstruction,
		"    ",
		viewInstruction,
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

func (m Model) renderContentView() string {
	title := titleStyle.Render(logo)
	subtitle := subtitleStyle.Render("Last Processed Content:")

	contentLines := strings.Split(m.lastProcessed, "\n")

	viewportHeight := 20
	totalLines := len(contentLines)
	maxScroll := max(0, totalLines-viewportHeight)

	startLine := m.scrollPosition
	endLine := min(startLine+viewportHeight, totalLines)

	visibleContent := strings.Join(contentLines[startLine:endLine], "\n")

	contentBox := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(highlight).
		Padding(1).
		Width(80).
		Height(viewportHeight + 2).
		Render(visibleContent)

	var scrollInfo string
	if totalLines > viewportHeight {
		scrollPercentage := float64(m.scrollPosition) / float64(maxScroll) * 100
		scrollInfo = infoStyle.Render(fmt.Sprintf("Line %d of %d (%.0f%%)",
			m.scrollPosition+1, totalLines, scrollPercentage))

		if m.scrollPosition > 0 && m.scrollPosition < maxScroll {
			scrollInfo += "  ↑ ↓"
		} else if m.scrollPosition > 0 {
			scrollInfo += "  ↑"
		} else if m.scrollPosition < maxScroll {
			scrollInfo += "  ↓"
		}
	}

	mutedInstructionStyle := buttonStyle.
		Foreground(subtle).
		Background(lipgloss.NoColor{}).
		Bold(false)

	upDownInstruction := mutedInstructionStyle.Render("[ ↑/↓ ] to scroll")
	escInstruction := mutedInstructionStyle.Render("[ ESC ] to return to monitoring view")
	quitInstruction := mutedInstructionStyle.Render("[ q ] to quit")

	instructions := lipgloss.JoinHorizontal(
		lipgloss.Center,
		upDownInstruction,
		"    ",
		escInstruction,
		"    ",
		quitInstruction,
	)

	return appStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			subtitle,
			scrollInfo,
			"",
			contentBox,
			"",
			instructions,
		),
	)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
