package config

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Config struct {
	Language string
	Format   bool
}

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	config   *Config
	done     bool
}

func GetConfig() *Config {
	if len(os.Args) > 1 {
		return parseFlags()
	}

	return launchUI()
}

func parseFlags() *Config {
	goPtr := flag.Bool("go", true, "Remove Go style comments")
	cLangPtr := flag.Bool("c", false, "Remove C/C++ style comments")
	javaPtr := flag.Bool("java", false, "Remove Java style comments")
	pythonPtr := flag.Bool("python", false, "Remove Python style comments")
	jsPtr := flag.Bool("js", false, "Remove JavaScript style comments")
	jsxPtr := flag.Bool("jsx", false, "Remove JSX style comments")
	formatPtr := flag.Bool("format", false, "Format the copied code automatically")
	flag.Parse()

	language := "go"
	if !*goPtr {
		if *cLangPtr {
			language = "c"
		} else if *javaPtr {
			language = "java"
		} else if *pythonPtr {
			language = "python"
		} else if *jsPtr {
			language = "javascript"
		} else if *jsxPtr {
			language = "jsx"
		}
	}

	return &Config{
		Language: language,
		Format:   *formatPtr,
	}
}

func launchUI() *Config {
	p := tea.NewProgram(initialModel())
	m, err := p.Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	finalModel := m.(model)
	if !finalModel.done {
		fmt.Println("Configuration cancelled")
		os.Exit(0)
	}

	return finalModel.config
}

func initialModel() model {
	return model{
		choices: []string{
			"Go",
			"C/C++",
			"Java",
			"Python",
			"JavaScript",
			"JSX",
			"Autoformat code? (only works for golang currently)",
		},
		selected: make(map[int]struct{}),
		config: &Config{
			Language: "go",
			Format:   false,
		},
		done: false,
	}
}

func (m model) Init() tea.Cmd {
	m.selected[0] = struct{}{}
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "j", "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			if m.cursor == len(m.choices)-1 {
				m.config.Format = !m.config.Format
				return m, nil
			}

			for i := range m.choices[:len(m.choices)-1] {
				delete(m.selected, i)
			}

			m.selected[m.cursor] = struct{}{}

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
		case " ":
			if m.cursor == len(m.choices)-1 {
				m.config.Format = !m.config.Format
				return m, nil
			}

			for i := range m.choices[:len(m.choices)-1] {
				delete(m.selected, i)
			}

			m.selected[m.cursor] = struct{}{}

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
		case "y", "Y", "s", "S":
			m.done = true
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "Select language for comment removal:\n\n"

	for i, choice := range m.choices[:len(m.choices)-1] {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\n"

	cursor := " "
	if m.cursor == len(m.choices)-1 {
		cursor = ">"
	}

	formatStatus := " "
	if m.config.Format {
		formatStatus = "x"
	}

	s += fmt.Sprintf("%s [%s] %s\n", cursor, formatStatus, m.choices[len(m.choices)-1])

	s += "\nPress space or enter to select an option\n"
	s += "Press y to confirm and start monitoring\n"
	s += "Press q to quit\n"

	s += "\nCurrent configuration:\n"
	s += fmt.Sprintf("Language: %s\n", m.config.Language)
	s += fmt.Sprintf("Autoformat: %v\n", m.config.Format)

	return s
}
