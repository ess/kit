package kit

type Tool struct {
	Name   string `yaml:"name"`
	Image  string `yaml:"image"`
	Tag    string `yaml:"tag"`
	TTY    bool   `yaml:"tty"`
	Stream bool   `yaml:"stream"`
}

type ToolService interface {
	All() []*Tool
	Find(string) (*Tool, error)
	Persist(*Tool) (*Tool, error)
}
