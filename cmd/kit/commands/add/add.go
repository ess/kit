package add

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add a tool to the toolkit",
	Long:    `Add a tool to the toolkit`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Called as", cmd.CalledAs())

		return nil
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}
