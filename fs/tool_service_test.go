package fs

import (
	"testing"

	"github.com/ess/kit/core"
	"github.com/ess/testscope"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
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

			for _, term := range known {
				found := false

				for _, prospect := range it {
					if term.Name == prospect.Name {
						found = true
						break
					}
				}

				if !found {
					t.Errorf("expected %s to be in the returned collection", term.Name)
				}
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

		expectedName := "nonexistenttool"

		tool, err := service.Find(expectedName)

		t.Run("it returns no tool", func(t *testing.T) {
			if tool != nil {
				t.Errorf("expected no tool, got %s", tool.Name)
			}
		})

		t.Run("it returns an error", func(t *testing.T) {
			if err == nil {
				t.Errorf("expected an error, got nil")
			}
		})

	})
}

func TestToolService_Persist(t *testing.T) {
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

		tool := known[0]

		// let's set up some stuff so we can ensure that it actually writes the
		// data
		origTag := tool.Tag
		tool.Tag = origTag + "sausages"

		err = service.Persist(tool)

		t.Run("it returns no error", func(t *testing.T) {
			if err != nil {
				t.Errorf("expected no error, got %s", err)
			}
		})

		t.Run("it writes the tool to disk", func(t *testing.T) {
			path := baseDir + "/" + tool.Name

			data, err := ReadFile(path)
			if err != nil {
				t.Errorf("could not read %s", path)
			}

			prospect := &core.Tool{}
			err = yaml.Unmarshal(data, prospect)
			if err != nil {
				t.Errorf("could not decode %s", path)
			}

			if prospect.Tag == origTag {
				t.Errorf("expected new content to be written")
			}

			if !FileExists(path) {
				t.Errorf("expected %s to exist", path)
			}
		})

	})

	t.Run("when the requested tool doesn't exist", func(t *testing.T) {
		setupFs()
		CreateDir(baseDir)

		name := "totallynewtool"

		tool := core.NewTool(name)

		err := service.Persist(tool)

		t.Run("it returns no error", func(t *testing.T) {
			if err != nil {
				t.Errorf("expected no error, got %s", err)
			}
		})

		t.Run("it writes the tool to disk", func(t *testing.T) {
			path := baseDir + "/" + name

			if !FileExists(path) {
				t.Errorf("expected %s to exist", path)
			}
		})
	})

	t.Run("when there are write errors", func(t *testing.T) {
		setupFs()
		CreateDir(baseDir)
		known, err := stubTools(baseDir, 1)
		if err != nil {
			t.Errorf("could not stub tools")
		}

		// let's force some write errors
		Root = afero.NewReadOnlyFs(Root)

		tool := known[0]

		err = service.Persist(tool)

		t.Run("it returns an error", func(t *testing.T) {
			if err == nil {
				t.Errorf("expected an error, but got nil")
			}
		})

	})

}

func TestToolService_Delete(t *testing.T) {
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

		tool := known[0]

		err = service.Delete(tool)

		t.Run("it returns no error", func(t *testing.T) {
			if err != nil {
				t.Errorf("expected no error, got %s", err)
			}
		})

		t.Run("it removes the tool from the file system", func(t *testing.T) {
			path := baseDir + "/" + tool.Name
			if FileExists(path) {
				t.Errorf("expected %s to no longer exist", path)
			}
		})

	})

	t.Run("when the requested tool does not exist", func(t *testing.T) {
		setupFs()
		CreateDir(baseDir)

		tool := core.NewTool("nonexistenttool")

		err := service.Delete(tool)

		t.Run("it returns no error", func(t *testing.T) {
			if err != nil {
				t.Errorf("expected no error, got %s", err)
			}
		})
	})

	t.Run("when there are fs delete errors", func(t *testing.T) {
		setupFs()
		CreateDir(baseDir)
		known, err := stubTools(baseDir, 1)
		if err != nil {
			t.Errorf("could not stub tools")
		}

		// let's force some delete errors
		Root = afero.NewReadOnlyFs(Root)

		tool := known[0]

		err = service.Delete(tool)

		t.Run("it returns an error", func(t *testing.T) {
			if err == nil {
				t.Errorf("expected an error, but got nil")
			}
		})

	})

}
