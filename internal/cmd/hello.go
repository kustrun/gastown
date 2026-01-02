package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:     "hello",
	GroupID: GroupDiag,
	Short:   "Print a friendly greeting",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello Gas Town")
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
