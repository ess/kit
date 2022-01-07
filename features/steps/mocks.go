package steps

import (
	"github.com/ess/kennel"

	"github.com/ess/kit/cmd/kit/services"
	"github.com/ess/kit/mock"
)

// Mock out the services

var toolService = mock.NewToolService()
var containerService = mock.NewContainerService()
var linkService = mock.NewLinkService()

type mockUp struct{}

func (steps mockUp) StepUp(s kennel.Suite) {
	s.BeforeSuite(func() {
		services.ToolService = toolService
		services.ContainerService = containerService
		services.LinkService = linkService
	})

	s.BeforeScenario(func(interface{}) {
		ResetMocks()
	})
}

func ResetMocks() {
	toolService.Reset()
	containerService.Reset()
	linkService.Reset()
}

func init() {
	kennel.Register(
		&mockUp{},
	)
}
