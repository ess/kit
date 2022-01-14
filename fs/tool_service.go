package fs

import (
	"fmt"

	"github.com/ess/kit/core"
	"github.com/ess/kit/yaml"
)

type ToolService struct {
	baseDir string
}

func NewToolService(baseDir string) *ToolService {
	t := &ToolService{baseDir}

	return t
}

func (service *ToolService) readTool(name string) (*core.Tool, error) {
	path := service.path(name)

	data, err := ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read %s", path)
	}

	tool, err := yaml.DecodeTool(data)
	if err != nil {
		return nil, fmt.Errorf("could not decode %s", path)
	}

	return tool, nil

}

func (service *ToolService) All() []*core.Tool {

	output := make([]*core.Tool, 0)

	candidates, err := ReadDir(service.baseDir)
	if err != nil {
		return output
	}

	for _, f := range candidates {
		if tool, err := service.readTool(f.Name()); err == nil {
			output = append(output, tool)
		}
	}

	return output
}

func (service *ToolService) Find(name string) (*core.Tool, error) {
	return service.readTool(name)
}

func (service *ToolService) Persist(tool *core.Tool) error {
	data, err := yaml.EncodeTool(tool)
	if err != nil {
		return fmt.Errorf("could not encode tool %s", tool.Name)
	}

	err = WriteFile(service.path(tool.Name), data, 0644)
	if err != nil {
		return fmt.Errorf("could not write tool %s: %s", tool.Name, err)
	}

	return nil
}

func (service *ToolService) Delete(tool *core.Tool) error {
	path := service.path(tool.Name)

	if !FileExists(path) {
		return nil
	}

	return Delete(path)
}

func (service *ToolService) path(name string) string {
	return service.baseDir + "/" + name
}
