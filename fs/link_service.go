package fs

import (
	"path/filepath"

	"github.com/ess/kit/core"
)

type LinkService struct {
	baseDir string
}

func NewLinkService(baseDir string) *LinkService {
	return &LinkService{baseDir}
}

func (service *LinkService) Link(tool *core.Tool, target string) error {
	path := service.path(tool.Name)

	if !FileExists(path) {
		return SymlinkIfPossible(target, path)
	}

	return nil
}

func (service *LinkService) path(name string) string {
	return filepath.Join(service.baseDir, name)
}
