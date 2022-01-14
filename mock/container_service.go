package mock

import (
	"github.com/ess/kit/core"
)

type ContainerService struct {
	pulls map[string]bool
}

func NewContainerService() *ContainerService {
	t := &ContainerService{}
	t.Reset()

	return t
}

func (service *ContainerService) Pull(tool *core.Tool) error {
	service.pulls[tool.Name] = true

	return nil
}

func (service *ContainerService) Pulled(tool *core.Tool) bool {
	found, ok := service.pulls[tool.Name]
	if !ok {
		return false
	}

	return found
}

func (service *ContainerService) Reset() {
	service.pulls = make(map[string]bool)
}
