package services

import (
	"github.com/ess/kit/core"
	"github.com/ess/kit/fs"

	"github.com/ess/kit/cmd/kit/util"
)

var ToolService core.ToolService

func init() {
	ToolService = fs.NewToolService(util.ToolsPath)
}
