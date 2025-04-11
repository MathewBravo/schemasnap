package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Snapshot(cmd *cobra.Command, args []string) {
	fmt.Println("Taking a snapshot")
}
