package fetch

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:     "fetch",
	Aliases: []string{"f"},
	Short:   "Fetch a group manifest from an external source",
	Long:    `Fetch a group manifest from an external source`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Called as", cmd.CalledAs())

		return nil
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}
