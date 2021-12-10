package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update kit to the latest version",
	Long:  `Update kit to the latest version`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Called as", cmd.CalledAs())

		return nil
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	RootCmd.AddCommand(updateCmd)
}
