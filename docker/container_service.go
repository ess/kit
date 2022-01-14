package docker

import (
	"github.com/ess/kit/core"
	"github.com/ess/kit/os"
)

type ContainerService struct{}

func NewContainerService() *ContainerService {
	return &ContainerService{}
}

func (service *ContainerService) Pull(tool *core.Tool) error {
	return os.Execute("docker", "pull", tool.Image+":"+tool.Tag)
}
