package steps

import (
	"github.com/ess/fakefs"
	"github.com/ess/kennel"

	"github.com/ess/kit/cmd/kit/services"
	"github.com/ess/kit/fs"
	"github.com/ess/kit/mock"
)

// Mock out the services

//var toolService = mock.NewToolService()
var containerService = mock.NewContainerService()

//var linkService = mock.NewLinkService()

type mockUp struct{}

func (steps mockUp) StepUp(s kennel.Suite) {
	s.BeforeSuite(func() {
		fs.Root = fakefs.New()
		//services.ToolService = toolService
		services.ContainerService = containerService
		//services.LinkService = linkService
	})

	s.BeforeScenario(func(interface{}) {
		ResetMocks()
	})
}

func ResetMocks() {
	fs.Root = fakefs.New()
	//toolService.Reset()
	containerService.Reset()
	//linkService.Reset()
}

func init() {
	kennel.Register(
		&mockUp{},
	)
}
