package commands

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/spf13/viper"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
)

type model struct {
	inputs     []textinput.Model
	focusIndex int
}

type Config struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Database string `mapstructure:"database"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

func InitialModel() model {
	m := model{
		inputs: make([]textinput.Model, 5),
	}

	placeholders := []string{"Host", "Port", "Database", "User", "Password"}

	for i := range m.inputs {
		t := textinput.New()
		t.Placeholder = placeholders[i]
		t.CharLimit = 64
		t.Width = 20

		if i == 4 {

			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
		}

		if i == 0 {
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		}

		m.inputs[i] = t
	}

	return m

}

func (m model) Init() tea.Cmd {
	m.createConfigDir()
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "enter":
			if m.focusIndex == len(m.inputs)-1 {
				if err := m.saveConfig(); err != nil {
					panic(err)
				}

				return m, tea.Quit
			}
			m.inputs[m.focusIndex].Blur()
			m.focusIndex++
			m.inputs[m.focusIndex].Focus()
		}
	}

	// Update all inputs
	for i := range m.inputs {
		m.inputs[i], _ = m.inputs[i].Update(msg)
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var b strings.Builder

	b.WriteString("\nAt the moment only POSTGRES is supported\n\n")

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		b.WriteRune('\n')
	}

	b.WriteString("\nPress Enter to submit, Ctrl+C or Esc to quit.\n")

	return b.String()
}

func (m model) createConfigDir() tea.Msg {
	cdir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	appConfigDir := filepath.Join(cdir, "schemasnap")
	if _, err := os.Stat(appConfigDir); os.IsNotExist(err) {
		if err := os.MkdirAll(appConfigDir, 0755); err != nil {
			return err
		}
	}
	return true
}

func (m model) saveConfig() error {
	cdir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	appConfigDir := filepath.Join(cdir, "schemasnap")
	configFile := filepath.Join(appConfigDir, "config.yaml")

	// Capture user input into Config struct
	cfg := Config{
		Host:     m.inputs[0].Value(),
		Port:     m.inputs[1].Value(),
		Database: m.inputs[2].Value(),
		User:     m.inputs[3].Value(),
		Password: m.inputs[4].Value(),
	}

	// Tell viper where to write
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")

	// Pass values into viper
	viper.Set("host", cfg.Host)
	viper.Set("port", cfg.Port)
	viper.Set("database", cfg.Database)
	viper.Set("user", cfg.User)
	viper.Set("password", cfg.Password)

	// Save the config
	return viper.WriteConfigAs(configFile)
}
