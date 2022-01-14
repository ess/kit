package fs

import (
	"os"

	"github.com/spf13/afero"
)

var Root = afero.NewOsFs()

func ReadFile(path string) ([]byte, error) {
	return afero.ReadFile(Root, path)
}

func WriteFile(path string, data []byte, perms os.FileMode) error {
	return afero.WriteFile(Root, path, data, perms)
}

var FileExists = func(path string) bool {
	_, err := Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return true
}

var IsDir = func(path string) bool {
	info, err := Stat(path)
	if err != nil {
		return false
	}

	return info.IsDir()
}

var DirectoryExists = func(path string) bool {
	if !FileExists(path) {
		return false
	}

	if !IsDir(path) {
		return false
	}

	return true
}

var CreateDir = func(path string) error {
	if !FileExists(path) {
		err := Root.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

var DeleteDir = func(path string) error {
	return Root.RemoveAll(path)
}

func ReadDir(path string) ([]os.FileInfo, error) {
	return afero.ReadDir(Root, path)
}

func Stat(path string) (os.FileInfo, error) {
	return Root.Stat(path)
}
