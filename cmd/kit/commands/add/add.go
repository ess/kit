package add

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/ess/kit/core"

	"github.com/ess/kit/cmd/kit/services"
)

var addImage string
var addTag string
var addNoStream bool

var Command = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add a tool to the toolkit",
	Long:    `Add a tool to the toolkit`,
	RunE: func(cmd *cobra.Command, args []string) error {
		image := addImage

		if !strings.HasPrefix(image, "docker.io/") {
			image = "docker.io/" + image
		}

		candidate := core.NewTool(args[0])
		candidate.Image = image
		candidate.Tag = addTag
		candidate.Stream = !addNoStream

		err := services.ToolService.Persist(candidate)
		if err != nil {
			return err
		}

		err = services.ContainerService.Pull(candidate)
		if err != nil {
			return err
		}

		me, err := filepath.Abs(os.Args[0])
		if err != nil {
			return fmt.Errorf("could not detect my absolute path")
		}

		return services.LinkService.Link(candidate, me)
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	Command.Flags().StringVarP(
		&addImage,
		"image",
		"i",
		"docker.io/wayneeseguin/c3tk",
		"The upstream container image for the tool",
	)

	Command.Flags().StringVarP(
		&addTag,
		"tag",
		"t",
		"latest",
		"The default upstream container tag for the tool",
	)

	Command.Flags().BoolVarP(
		&addNoStream,
		"no-stream",
		"n",
		false,
		"Disable IO streaming for the tool",
	)

}
