package config

import (
	"fmt"
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

		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "j", "down":
			if m.screen == languageSelect && m.cursor < len(m.languageChoices)-1 {
				m.cursor++
			} else if m.screen == formatSelect && m.cursor < len(m.formatChoices)-1 {
				m.cursor++
			}

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
				m.outputs = append(m.outputs, fmt.Sprintf("Error: %v", err))
			} else if processed != content {
				m.outputs = append(m.outputs, "Processed clipboard content")

				clipboard.Write(clipboard.FmtText, []byte(processed))
			}
		}

		return m, CheckClipboard()

	case ErrorMsg:
		m.outputs = append(m.outputs, fmt.Sprintf("Error: %v", msg))
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
