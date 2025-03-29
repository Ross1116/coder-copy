package config

import (
	"github.com/charmbracelet/lipgloss"
)

var logo = ` 
 $$$$$$\                  $$\                            $$$$$$\                                
$$  __$$\                 $$ |                          $$  __$$\                               
$$ /  \__| $$$$$$\   $$$$$$$ | $$$$$$\   $$$$$$\        $$ /  \__| $$$$$$\   $$$$$$\  $$\   $$\ 
$$ |      $$  __$$\ $$  __$$ |$$  __$$\ $$  __$$\       $$ |      $$  __$$\ $$  __$$\ $$ |  $$ |
$$ |      $$ /  $$ |$$ /  $$ |$$$$$$$$ |$$ |  \__|      $$ |      $$ /  $$ |$$ /  $$ |$$ |  $$ |
$$ |  $$\ $$ |  $$ |$$ |  $$ |$$   ____|$$ |            $$ |  $$\ $$ |  $$ |$$ |  $$ |$$ |  $$ |
\$$$$$$  |\$$$$$$  |\$$$$$$$ |\$$$$$$$\ $$ |            \$$$$$$  |\$$$$$$  |$$$$$$$  |\$$$$$$$ |
 \______/  \______/  \_______| \_______|\__|             \______/  \______/ $$  ____/  \____$$ |
                                                                            $$ |      $$\   $$ |
                                                                            $$ |      \$$$$$$  |
                                                                            \__|       \______/`

var (
	subtle    = lipgloss.AdaptiveColor{Light: "#6C7A96", Dark: "#4A4F5E"}
	highlight = lipgloss.AdaptiveColor{Light: "#8A67B3", Dark: "#B58ED2"}
	special   = lipgloss.AdaptiveColor{Light: "#5A9CBF", Dark: "#5FA4C3"}
	text      = lipgloss.AdaptiveColor{Light: "#2A2E38", Dark: "#D8DEE9"}

	errorColor   = lipgloss.AdaptiveColor{Light: "#F25D94", Dark: "#F25D94"}
	successColor = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	warningColor = lipgloss.AdaptiveColor{Light: "#FFA500", Dark: "#FFD700"}
)

var (
	appStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(highlight)

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(highlight).
			MarginBottom(1)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(text).
			MarginBottom(1)

	listItemStyle = lipgloss.NewStyle().
			PaddingLeft(2)

	selectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(0).
				Foreground(special).
				SetString("▸ ")

	infoStyle = lipgloss.NewStyle().
			Foreground(text)

	highlightedInfoStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(highlight)

	buttonStyle = lipgloss.NewStyle().
			Padding(0, 1).
			MarginTop(1).
			Bold(true).
			Foreground(text).
			Background(subtle)

	successLogStyle = lipgloss.NewStyle().
			Foreground(successColor)

	errorLogStyle = lipgloss.NewStyle().
			Foreground(errorColor)

	warningLogStyle = lipgloss.NewStyle().
			Foreground(warningColor)
)
