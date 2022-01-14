package core

type Tool struct {
	Name   string `yaml:"name"`
	Image  string `yaml:"image"`
	Tag    string `yaml:"tag"`
	Stream bool   `yaml:"stream"`
}

func NewTool(name string) *Tool {
	return &Tool{
		Name:   name,
		Image:  "docker.io/wayneeseguin/c3tk",
		Tag:    "latest",
		Stream: true,
	}
}

type ToolService interface {
	All() []*Tool
	Find(string) (*Tool, error)
	Persist(*Tool) error
	Delete(*Tool) error
}

type LinkService interface {
	Link(*Tool, string) error
}

type ContainerService interface {
	Pull(*Tool) error
}

type Group struct {
	Tools []*Tool `yaml:tools`
}

type GroupFetcher interface {
	Fetch(string) (*Group, error)
}
