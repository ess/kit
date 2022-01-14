package services

import (
	"github.com/ess/kit/core"
	"github.com/ess/kit/docker"
)

var ContainerService core.ContainerService

func init() {
	ContainerService = docker.NewContainerService()
}
