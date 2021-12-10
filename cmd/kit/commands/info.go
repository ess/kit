package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:     "info",
	Aliases: []string{"info"},
	Short:   "Show information about tools",
	Long:    `Show information about tools`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Called as", cmd.CalledAs())

		return nil
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	RootCmd.AddCommand(infoCmd)
}
