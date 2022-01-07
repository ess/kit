package mock

import (
	"github.com/ess/kit/core"
)

type LinkService struct {
	links []*core.Tool
}

func NewLinkService() *LinkService {
	t := &LinkService{}
	t.Reset()

	return t
}

func (service *LinkService) Link(tool *core.Tool, target string) error {
	service.links = append(service.links, tool)

	return nil
}

func (service *LinkService) Linked(tool *core.Tool) bool {
	for _, candidate := range service.links {
		if candidate == tool {
			return true
		}
	}

	return false
}

func (service *LinkService) Reset() {
	service.links = make([]*core.Tool, 0)
}
