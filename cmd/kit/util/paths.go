package util

import (
	"fmt"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"

	"github.com/ess/kit/fs"
)

var paths []string
var BasePath string
var ToolsPath string
var LinksPath string

func createPath(path string) error {
	if !fs.DirectoryExists(path) {
		if err := fs.CreateDir(path); err != nil {
			return fmt.Errorf("could not create %s: %s", path, err)
		}
	}

	return nil
}

func CreatePaths() error {
	for _, path := range paths {
		if err := createPath(path); err != nil {
			return err
		}
	}

	return nil
}

func init() {
	BasePath, err := homedir.Expand("~/.kit")
	if err != nil {
		BasePath = "."
	}

	ToolsPath = filepath.Join(BasePath, "tools")
	LinksPath = filepath.Join(BasePath, "bin")

	paths = make([]string, 0)
	paths = append(paths, BasePath)
	paths = append(paths, ToolsPath)
	paths = append(paths, LinksPath)
}
