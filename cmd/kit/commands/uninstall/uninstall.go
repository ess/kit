package uninstall

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall kit",
	Long:  `Uninstall kit`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Called as", cmd.CalledAs())

		return nil
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}
