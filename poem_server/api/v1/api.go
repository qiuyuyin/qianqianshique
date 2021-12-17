package v1

import (
	"poem_server/api/v1/poetry"
	"poem_server/api/v1/system"
)

type PoemApiGroup struct {
	PoetryApiGroup poetry.PoemApiGroup
	SystemApiGroup system.SysApiGroup
}

var PoemApiGroupApp = new(PoemApiGroup)
