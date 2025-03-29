package config

type screenState int

const (
	languageSelect screenState = iota
	formatSelect
	monitoring
)

type Model struct {
	screen          screenState
	cursor          int
	languageChoices []string
	formatChoices   []string
	config          *Config
	outputs         []string
	lastClipboard   string
	processContent  func(string, string, bool) (string, error)
}

type ClipboardUpdateMsg string
type ErrorMsg error

func initialModel(processContentFn func(string, string, bool) (string, error)) Model {
	return Model{
		screen: languageSelect,
		languageChoices: []string{
			"Go",
			"C/C++",
			"Java",
			"Python",
			"JavaScript",
			"JSX",
		},
		formatChoices: []string{
			"Yes",
			"No",
		},
		config: &Config{
			Language: "go",
			Format:   false,
		},
		outputs:        []string{},
		processContent: processContentFn,
	}
}

func (m Model) GetCurrentConfig() *Config {
	return m.config
}
