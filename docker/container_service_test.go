package docker

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ess/kit/core"
	"github.com/ess/kit/os"
	"github.com/ess/testscope"
)

func stubExecute(fail bool) {
	os.Execute = func(command string, args ...string) error {
		output := make([]string, 0)
		output = append(output, command)

		for _, arg := range args {
			output = append(output, arg)
		}

		if fail {
			return fmt.Errorf("%s", strings.Join(output, " "))
		}

		return nil
	}
}

func TestContainerService_Pull(t *testing.T) {
	testscope.SkipUnlessUnit(t)
	service := NewContainerService()

	tool := core.NewTool("sausages")

	t.Run("when docker pull succeeds", func(t *testing.T) {
		// stub execution to mimic a success
		stubExecute(false)

		err := service.Pull(tool)

		t.Run("it returns no error", func(t *testing.T) {
			if err != nil {
				t.Errorf("expected no error, got %s", err)
			}
		})
	})

	t.Run("when docker pull fails", func(t *testing.T) {
		// stub execution to mimic a failure
		stubExecute(true)

		err := service.Pull(tool)

		t.Run("it returns an error", func(t *testing.T) {
			if err == nil {
				fmt.Errorf("expected an error, got nil")
			}
		})
	})

}
