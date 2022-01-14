package steps

import (
	"github.com/ess/kennel"
)

type OutputSteps struct{}

func (steps *OutputSteps) StepUp(s kennel.Suite) {
}

func init() {
	kennel.Register(new(OutputSteps))
}
