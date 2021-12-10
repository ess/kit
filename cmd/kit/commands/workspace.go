package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var workspaceCmd = &cobra.Command{
	Use:     "workspace",
	Aliases: []string{"w"},
	Short:   "workspace commands",
	Long:    `workspace commands`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Called as", cmd.CalledAs())

		return nil
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	RootCmd.AddCommand(workspaceCmd)
}
