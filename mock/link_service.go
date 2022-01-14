package mock

import (
	"github.com/ess/kit/core"
)

type LinkService struct {
	links map[string]bool
}

func NewLinkService() *LinkService {
	t := &LinkService{}
	t.Reset()

	return t
}

func (service *LinkService) Link(tool *core.Tool, target string) error {
	service.links[tool.Name] = true

	return nil
}

func (service *LinkService) Linked(tool *core.Tool) bool {
	found, ok := service.links[tool.Name]
	if !ok {
		return false
	}

	return found
}

func (service *LinkService) Reset() {
	service.links = make(map[string]bool)
}
