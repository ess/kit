package fs

import (
	"path/filepath"
	"testing"

	"github.com/ess/kit/core"
	"github.com/ess/testscope"
)

func TestLinkService_Link(t *testing.T) {
	testscope.SkipUnlessUnit(t)

	baseDir := "/links"
	service := NewLinkService(baseDir)

	target := "/my/sausages/turned/to/gold"

	t.Run("when there is already a link", func(t *testing.T) {
		setupFs()
		CreateDir(baseDir)

		tool := core.NewTool("existent")
		err := service.Link(tool, target)
		if err != nil {
			t.Errorf("could not create initial link for preconditions: %s", err)
		}

		err = service.Link(tool, target)

		t.Run("it returns no error", func(t *testing.T) {
			if err != nil {
				t.Errorf("expected no error, got %s", err)
			}
		})

	})

	t.Run("when there is not already a link", func(t *testing.T) {
		setupFs()
		CreateDir(baseDir)

		tool := core.NewTool("nonexistent")

		err := service.Link(tool, target)

		t.Run("it returns no error", func(t *testing.T) {
			if err != nil {
				t.Errorf("expected no error, got %s", err)
			}
		})

		t.Run("it creates the link", func(t *testing.T) {
			path := filepath.Join(baseDir, tool.Name)

			if !FileExists(path) {
				t.Errorf("expected %s to exist", path)
			}
		})
	})
}
