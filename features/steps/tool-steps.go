package steps

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/ess/jamaica"
	"github.com/ess/kennel"
	"github.com/ess/kit/cmd/kit/services"
	"github.com/ess/kit/cmd/kit/util"
	"github.com/ess/kit/core"
	"github.com/ess/kit/fs"
)

type ToolSteps struct{}

func (steps *ToolSteps) StepUp(s kennel.Suite) {
	s.Step(`^there's not a jq tool configured$`, func() error {
		if t, err := services.ToolService.Find("jq"); err == nil {
			services.ToolService.Delete(t)
		}

		return nil
	})

	s.Step(`^the jq tool is configured$`, func() error {
		_, err := services.ToolService.Find("jq")

		return err
	})

	s.Step(`^a jq symlink to kit now exists$`, func() error {
		t, err := services.ToolService.Find("jq")
		if err != nil {
			return err
		}

		//if !linkService.Linked(t) {
		//fmt.Println("the tool's name is", t.Name)
		//return fmt.Errorf("Expected jq to have been linked")
		//}

		if !fs.FileExists(filepath.Join(util.LinksPath, t.Name)) {
			return fmt.Errorf("Expected jq to have been linked")
		}

		return nil
	})

	s.Step(`^jq's image is (.+)$`, func(image string) error {
		t, err := services.ToolService.Find("jq")
		if err != nil {
			return err
		}

		expected := image

		if t.Image != expected {
			return fmt.Errorf("Expected '%s', got '%s'", expected, t.Image)
		}

		return nil
	})

	s.Step(`^jq's image gets pulled from upstream$`, func() error {
		t, err := services.ToolService.Find("jq")
		if err != nil {
			return err
		}

		expected := fmt.Sprintf("EXECUTE: docker pull %s:%s", t.Image, t.Tag)

		if !strings.Contains(jamaica.LastCommandStdout(), expected) {
			return fmt.Errorf("Expected jq to have been pulled")
		}

		return nil
	})

	s.Step(`^jq's default tag is (.+)$`, func(tag string) error {
		t, err := services.ToolService.Find("jq")
		if err != nil {
			return err
		}

		expected := tag

		if t.Tag != expected {
			return fmt.Errorf("Expected '%s', got '%s'", expected, t.Image)
		}

		return nil

	})

	s.Step(`^jq is set up to stream IO$`, func() error {
		t, err := services.ToolService.Find("jq")
		if err != nil {
			return err
		}

		if !t.Stream {
			return fmt.Errorf("Expected jq to be set up for streaming.")
		}

		return nil
	})

	s.Step(`^the jq tool is successfully added with the (.+) image$`, func(image string) error {
		t, err := services.ToolService.Find("jq")
		if err != nil {
			return err
		}

		actual := t.Image
		expected := image

		if actual != expected {
			return fmt.Errorf("Expected '%s', got '%s'", expected, actual)
		}

		return nil
	})

	s.Step(`^the jq tool is successfully added with the (.+) tag$`, func(tag string) error {
		t, err := services.ToolService.Find("jq")
		if err != nil {
			return err
		}

		actual := t.Tag
		expected := tag

		if actual != expected {
			return fmt.Errorf("Expected '%s', got '%s'", expected, actual)
		}

		return nil
	})

	s.Step(`^the jq tool is successfully added with IO streaming disabled$`, func() error {
		t, err := services.ToolService.Find("jq")
		if err != nil {
			return err
		}

		if t.Stream {
			return fmt.Errorf("Expected jq's streaming to be disabled")
		}

		return nil
	})

	s.Step(`^there is already a jq tool configured with default settings$`, func() error {
		jq := core.NewTool("jq")

		err := services.ToolService.Persist(jq)
		if err != nil {
			return fmt.Errorf("Could not set up default jq.")
		}

		orig := &core.Tool{}
		orig.Name = jq.Name
		orig.Image = jq.Image
		orig.Tag = jq.Tag
		orig.Stream = jq.Stream

		err = facts.Memorize("orig", orig)
		if err != nil {
			return err
		}

		return nil
	})

	s.Step(`^jq's configuration is updated`, func() error {
		tool, err := services.ToolService.Find("jq")
		if err != nil {
			return err
		}

		candidate, err := facts.Recall("orig")
		if err != nil {
			return err
		}

		orig, ok := candidate.(*core.Tool)
		if !ok {
			return fmt.Errorf("orig is not a tool!")
		}

		if orig.Image != tool.Image {
			return nil
		}

		if orig.Tag != tool.Tag {
			return nil
		}

		if orig.Stream != tool.Stream {
			return nil
		}

		return fmt.Errorf("Expected changes, but found none")
	})

}

func init() {
	kennel.Register(new(ToolSteps))
}
