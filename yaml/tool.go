package yaml

import (
	"gopkg.in/yaml.v2"

	"github.com/ess/kit/core"
)

var DecodeTool = func(data []byte) (*core.Tool, error) {
	tool := &core.Tool{}
	err := yaml.Unmarshal(data, tool)

	return tool, err
}

var EncodeTool = func(tool *core.Tool) ([]byte, error) {
	return yaml.Marshal(tool)
}
