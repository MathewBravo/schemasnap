package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Diff(cmd *cobra.Command, args []string) {
	fmt.Println("Doing a diff")
}
