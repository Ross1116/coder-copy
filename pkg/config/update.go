package config

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"golang.design/x/clipboard"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "v":
			if m.screen == monitoring && m.lastProcessed != "" {
				m.screen = contentView
			}
			return m, nil

		case "esc":
			if m.screen == contentView {
				m.screen = monitoring
			}
			return m, nil

		case "backspace":
			if m.screen == formatSelect {
				m.screen = languageSelect
				m.cursor = 0
				for i, lang := range m.languageChoices {
					switch m.config.Language {
					case "go":
						if lang == "Go" {
							m.cursor = i
						}
					case "c":
						if lang == "C/C++" {
							m.cursor = i
						}
					case "java":
						if lang == "Java" {
							m.cursor = i
						}
					case "python":
						if lang == "Python" {
							m.cursor = i
						}
					case "javascript":
						if lang == "JavaScript" {
							m.cursor = i
						}
					case "jsx":
						if lang == "JSX" {
							m.cursor = i
						}
					}
				}
			}
			return m, nil

		case "s":
			if m.screen == monitoring {
				m.screen = languageSelect
				m.cursor = 0
				for i, lang := range m.languageChoices {
					switch m.config.Language {
					case "go":
						if lang == "Go" {
							m.cursor = i
						}
					case "c":
						if lang == "C/C++" {
							m.cursor = i
						}
					case "java":
						if lang == "Java" {
							m.cursor = i
						}
					case "python":
						if lang == "Python" {
							m.cursor = i
						}
					case "javascript":
						if lang == "JavaScript" {
							m.cursor = i
						}
					case "jsx":
						if lang == "JSX" {
							m.cursor = i
						}
					}
				}
			}
			return m, nil

		case "up", "k":
			if m.screen == contentView && m.scrollPosition > 0 {
				m.scrollPosition--
			} else if m.screen == languageSelect || m.screen == formatSelect {
				if m.cursor > 0 {
					m.cursor--
				}
			}
			return m, nil

		case "down", "j":
			if m.screen == contentView {
				contentLines := strings.Count(m.lastProcessed, "\n") + 1
				viewportHeight := 20
				maxScroll := max(0, contentLines-viewportHeight)

				if m.scrollPosition < maxScroll {
					m.scrollPosition++
				}
			} else if m.screen == languageSelect && m.cursor < len(m.languageChoices)-1 {
				m.cursor++
			} else if m.screen == formatSelect && m.cursor < len(m.formatChoices)-1 {
				m.cursor++
			}
			return m, nil

		case "enter", " ":
			if m.screen == languageSelect {
				switch m.cursor {
				case 0:
					m.config.Language = "go"
				case 1:
					m.config.Language = "c"
				case 2:
					m.config.Language = "java"
				case 3:
					m.config.Language = "python"
				case 4:
					m.config.Language = "javascript"
				case 5:
					m.config.Language = "jsx"
				}

				m.screen = formatSelect
				m.cursor = 0
				if m.config.Format {
					m.cursor = 0
				} else {
					m.cursor = 1
				}

			} else if m.screen == formatSelect {
				m.config.Format = m.cursor == 0

				m.screen = monitoring
				m.outputs = []string{}
				m.lastClipboard = ""

				return m, CheckClipboard()
			}
		}

	case ClipboardUpdateMsg:
		content := string(msg)
		if content != m.lastClipboard && content != "" {
			m.lastClipboard = content

			processed, err := m.processContent(content, m.config.Language, m.config.Format)
			if err != nil {
				if strings.Contains(err.Error(), "formatter not found") {
					m.config.Format = false
					m.addToOutputsQueue(fmt.Sprintf("Error: %v \nAutoformatting disabled", err))
				} else {
					m.addToOutputsQueue(fmt.Sprintf("Warning: %v", err))
				}
			}

			if processed != content {
				m.addToOutputsQueue("Processed clipboard content")
				m.lastProcessed = processed
				clipboard.Write(clipboard.FmtText, []byte(processed))
			}
		}

		return m, CheckClipboard()

	case ErrorMsg:
		m.addToOutputsQueue(fmt.Sprintf("Error: %v", msg))
		return m, CheckClipboard()
	}

	return m, nil
}

func CheckClipboard() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(500 * time.Millisecond)

		content := clipboard.Read(clipboard.FmtText)
		if content == nil {
			return ClipboardUpdateMsg("")
		}

		return ClipboardUpdateMsg(content)
	}
}

func (m *Model) addToOutputsQueue(message string) {
	const maxOutputs = 5

	m.outputs = append(m.outputs, message)

	if len(m.outputs) > maxOutputs {
		m.outputs = m.outputs[1:]
	}
}
