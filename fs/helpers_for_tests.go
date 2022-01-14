package fs

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"

	"github.com/ess/kit/core"
)

func setupFs() {
	Root = afero.NewMemMapFs()
}

func stubTools(baseDir string, quantity int) ([]*core.Tool, error) {
	tools := make([]*core.Tool, 0)

	for len(tools) < quantity {
		name := gofakeit.Generate("??????????")
		tool := core.NewTool(name)

		data, err := yaml.Marshal(tool)
		if err != nil {
			return tools, fmt.Errorf("could not marshal tool %s", tool.Name)
		}

		err = WriteFile(baseDir+"/"+tool.Name, data, 0644)
		if err != nil {
			return tools, fmt.Errorf("could not store tool %s", tool.Name)
		}

		tools = append(tools, tool)

	}

	return tools, nil
}
