package system

import (
	"github.com/gin-gonic/gin"
	v1 "poem_server/api/v1"
)

type BaseRouter struct {
}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	var baseApi = v1.PoemApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("register", baseApi.Register)
	}
	return baseRouter
}
