package mock

import (
	"github.com/ess/kit/core"
)

type ContainerService struct {
	pulls []*core.Tool
}

func NewContainerService() *ContainerService {
	t := &ContainerService{}
	t.Reset()

	return t
}

func (service *ContainerService) Pull(tool *core.Tool) error {
	service.pulls = append(service.pulls, tool)

	return nil
}

func (service *ContainerService) Pulled(tool *core.Tool) bool {
	for _, candidate := range service.pulls {
		if candidate == tool {
			return true
		}
	}

	return false
}

func (service *ContainerService) Reset() {
	service.pulls = make([]*core.Tool, 0)
}
