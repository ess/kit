package fs

import (
	"fmt"

	"github.com/ess/kit/core"
	"gopkg.in/yaml.v2"
)

type ToolService struct {
	baseDir string
}

func NewToolService(baseDir string) *ToolService {
	t := &ToolService{baseDir}

	return t
}

func (service *ToolService) All() []*core.Tool {

	output := make([]*core.Tool, 0)

	candidates, err := ReadDir(service.baseDir)
	if err != nil {
		return output
	}

	for _, f := range candidates {
		path := service.baseDir + "/" + f.Name()
		if data, err := ReadFile(path); err == nil {
			tool := &core.Tool{}
			if yErr := yaml.Unmarshal(data, tool); yErr == nil {
				output = append(output, tool)
			}
		}
	}

	return output
}

func (service *ToolService) Find(name string) (*core.Tool, error) {
	path := service.baseDir + "/" + name

	data, err := ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read %s", path)
	}

	tool := &core.Tool{}
	err = yaml.Unmarshal(data, tool)
	if err != nil {
		return nil, fmt.Errorf("could not decode %s", path)
	}

	return tool, nil
}

func (service *ToolService) Persist(tool *core.Tool) (*core.Tool, error) {

	return nil, fmt.Errorf("fs.ToolService.Persist unimplemented")
}

func (service *ToolService) Delete(tool *core.Tool) error {

	return fmt.Errorf("fs.ToolService.Delete unimplemented")
}
