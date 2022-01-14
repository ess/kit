package util

import (
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

var BasePath string
var ToolsPath string
var LinksPath string

func init() {
	BasePath, err := homedir.Expand("~/.kit")
	if err != nil {
		BasePath = "."
	}

	ToolsPath = filepath.Join(BasePath, "tools")
	LinksPath = filepath.Join(BasePath, "bin")
}
