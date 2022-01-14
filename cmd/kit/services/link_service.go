package services

import (
	"github.com/ess/kit/core"
	"github.com/ess/kit/fs"

	"github.com/ess/kit/cmd/kit/util"
)

var LinkService core.LinkService

func init() {
	LinkService = fs.NewLinkService(util.LinksPath)
}
