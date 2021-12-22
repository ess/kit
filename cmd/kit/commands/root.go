package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ess/kit/cmd/kit/commands/add"
	"github.com/ess/kit/cmd/kit/commands/configure"
	"github.com/ess/kit/cmd/kit/commands/fetch"
	"github.com/ess/kit/cmd/kit/commands/info"
	"github.com/ess/kit/cmd/kit/commands/list"
	"github.com/ess/kit/cmd/kit/commands/paths"
	"github.com/ess/kit/cmd/kit/commands/rm"
	"github.com/ess/kit/cmd/kit/commands/shell"
	"github.com/ess/kit/cmd/kit/commands/uninstall"
	"github.com/ess/kit/cmd/kit/commands/update"
	"github.com/ess/kit/cmd/kit/commands/workspace"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "kit",
	Short: "Containerized Command Center Toolkit",
	Long: `Containerized Command Center Toolkit

This top-level command is just a wrapper for other commands. Please see the
Available Commands section below.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		calledAs := filepath.Base(os.Args[0])

		if calledAs == "kit" {
			cmd.Help()

			return fmt.Errorf("ERROR: I'm not sure what you want to do.")
		}

		fmt.Println("I was called as", calledAs)

		return nil
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Execute() error {
	err := RootCmd.Execute()

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func init() {
	RootCmd.AddCommand(add.Command)
	RootCmd.AddCommand(configure.Command)
	RootCmd.AddCommand(fetch.Command)
	RootCmd.AddCommand(info.Command)
	RootCmd.AddCommand(list.Command)
	RootCmd.AddCommand(paths.Command)
	RootCmd.AddCommand(rm.Command)
	RootCmd.AddCommand(shell.Command)
	RootCmd.AddCommand(uninstall.Command)
	RootCmd.AddCommand(update.Command)
	RootCmd.AddCommand(workspace.Command)

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigName(".kit/config")
	viper.AddConfigPath("$HOME")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
	}
}
