package service

import (
	"poem_server/service/poetry"
	"poem_server/service/system"
)

type PoemServiceGroup struct {
	PoetryServiceGroup poetry.ServiceGroup
	SystemServiceGroup system.ServiceGroup
}

var PoemServiceGroupApp = new(PoemServiceGroup)
