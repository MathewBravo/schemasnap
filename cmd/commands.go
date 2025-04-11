package cmd

import (
	"fmt"

	"github.com/MathewBravo/schemasnap/internal/commands"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "schemasnap [SCHEMANAME]",
	Short: "Schemasnap is a database schema snapshot tool",
	Args:  cobra.ExactArgs(1),
	Run:   commands.Snapshot,
}

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows installed version of Schemasnap",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Schemasnap v0.01")
	},
}

var DiffCmd = &cobra.Command{
	Use:   "diff [OUTPUT1] [OUTPUT2]",
	Short: "Shows the differences between two schema outputs",
	Run:   commands.Diff,
}

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes shemasnap based on your provided config parameters",
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(commands.InitialModel())
		if _, err := p.Run(); err != nil {
			fmt.Println("Error running UI for Init: %v\n")
		}
	},
}
