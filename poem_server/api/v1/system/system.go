package system

import "poem_server/service"

type SysApiGroup struct {
	BaseApi
}

var userService = service.PoemServiceGroupApp.SystemServiceGroup.UserService
