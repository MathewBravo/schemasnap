package main

import (
	"fmt"
	"os"

	"github.com/MathewBravo/schemasnap/cmd"
)

func init() {
	cmd.RootCmd.AddCommand(cmd.VersionCmd, cmd.DiffCmd, cmd.InitCmd, cmd.TestCmd)
}

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
