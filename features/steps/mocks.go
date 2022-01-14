package steps

import (
	"fmt"

	"github.com/ess/fakefs"
	"github.com/ess/kennel"

	"github.com/ess/kit/fs"
	"github.com/ess/kit/os"
)

func stubExec() {
	os.Execute = func(command string, args ...string) error {
		printable := make([]interface{}, len(args)+2)
		printable[0] = "EXECUTE:"
		printable[1] = command
		for i := range args {
			printable[i+2] = args[i]
		}

		fmt.Println(printable...)
		return nil
	}
}

type mockUp struct{}

func (steps mockUp) StepUp(s kennel.Suite) {
	s.BeforeSuite(func() {
		fs.Root = fakefs.New()
		stubExec()
	})

	s.BeforeScenario(func(interface{}) {
		ResetMocks()
	})
}

func ResetMocks() {
	fs.Root = fakefs.New()
}

func init() {
	kennel.Register(
		&mockUp{},
	)
}
