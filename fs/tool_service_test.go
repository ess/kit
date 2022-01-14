package fs

import (
	"testing"

	"github.com/ess/testscope"
)

func TestToolService_All(t *testing.T) {
	testscope.SkipUnlessUnit(t)

	baseDir := "/tools"
	service := NewToolService(baseDir)

	t.Run("when there are no tools", func(t *testing.T) {
		setupFs()
		CreateDir(baseDir)

		t.Run("it is empty", func(t *testing.T) {
			it := service.All()

			actual := len(it)
			expected := 0

			if actual != expected {
				t.Errorf("expected %d Tools, got %d", expected, actual)
			}
		})
	})

	t.Run("when there are tools", func(t *testing.T) {
		setupFs()
		CreateDir(baseDir)
		known, err := stubTools(baseDir, 10)
		if err != nil {
			t.Errorf("could not stub tools (%s)", err)
		}

		t.Run("it contains each of the known tools", func(t *testing.T) {
			it := service.All()

			actual := len(it)
			expected := len(known)

			if actual != expected {
				t.Errorf("expected %d Tools, got %d", expected, actual)
			}
		})
	})
}

func TestToolService_Find(t *testing.T) {
	testscope.SkipUnlessUnit(t)

	baseDir := "/tools"
	service := NewToolService(baseDir)

	t.Run("when the requested tool exists", func(t *testing.T) {
		setupFs()
		CreateDir(baseDir)
		known, err := stubTools(baseDir, 1)
		if err != nil {
			t.Errorf("could not stub tools")
		}

		expectedName := known[0].Name

		tool, err := service.Find(expectedName)

		t.Run("it returns the requested tool", func(t *testing.T) {
			if tool == nil {
				t.Errorf("expected a tool, but got nil")
				return
			}

			if tool.Name != expectedName {
				t.Errorf(
					"expected a tool named '%s', got a tool named '%s'",
					tool.Name,
					expectedName,
				)
			}
		})

		t.Run("it returns no error", func(t *testing.T) {
			if err != nil {
				t.Errorf("expected no error, got %s", err)
			}
		})
	})

	t.Run("when the requested tool does not exist", func(t *testing.T) {
		setupFs()
		CreateDir(baseDir)

	})
}
