package kit

import (
	"os"
	"testing"

	"github.com/ess/kit/features/steps"

	"github.com/DATA-DOG/godog"
	"github.com/ess/jamaica"
	"github.com/ess/kennel"
	"github.com/ess/kit/cmd/kit/commands"
	"github.com/ess/mockable"
	"github.com/ess/testscope"
)

var commandOutput string
var lastCommandRanErr error

func TestMain(m *testing.M) {
	if testscope.Integration() {
		status := godog.RunWithOptions(
			"godog",
			func(s *godog.Suite) {
				mockable.Enable()
				jamaica.SetRootCmd(commands.RootCmd)
				jamaica.StepUp(s)
				steps.Register()
				kennel.StepUp(s)
			},

			godog.Options{
				Format: "pretty",
				Paths:  []string{"features"},

				// This isn't optimal, but it looks like we have to run tests in-order
				// for now until we figure out what's going on with jamaica.
				//Randomize: time.Now().UTC().UnixNano(),
			},
		)

		if st := m.Run(); st > status {
			status = st
		}

		os.Exit(status)
	}
}

func TestTrue(t *testing.T) {
}
