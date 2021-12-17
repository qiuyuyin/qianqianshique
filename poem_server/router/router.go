package router

import (
	"poem_server/router/poetry"
	"poem_server/router/system"
)

type PoemRouterGroup struct {
	Poetry poetry.RouterGroup
	System system.RouterGroup
}

var PoemRouterGroupApp = new(PoemRouterGroup)
