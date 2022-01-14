package mock

import (
	"fmt"

	"github.com/ess/kit/core"
)

type ToolService struct {
	tools map[string]*core.Tool
}

func NewToolService() *ToolService {
	t := &ToolService{}
	t.Reset()

	return t
}

func (service *ToolService) All() []*core.Tool {

	output := make([]*core.Tool, 0)

	for _, t := range service.tools {
		output = append(output, t)
	}

	return output
}

func (service *ToolService) Find(name string) (*core.Tool, error) {

	if t, ok := service.tools[name]; ok {
		return t, nil
	}

	return nil, fmt.Errorf("no such tool")
}

func (service *ToolService) Persist(tool *core.Tool) error {

	service.tools[tool.Name] = tool

	return nil
}

func (service *ToolService) Delete(tool *core.Tool) error {

	delete(service.tools, tool.Name)

	return nil
}

func (service *ToolService) Reset() {

	service.tools = make(map[string]*core.Tool)
	for k := range service.tools {
		delete(service.tools, k)
	}
}
