package system

import (
	"github.com/gin-gonic/gin"
	v1 "poem_server/api/v1"
)

type UserRouter struct {
}

func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	userRouter := Router.Group("user")
	var baseApi = v1.PoemApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.POST("createUserPoem", baseApi.PostUserPoem)
		userRouter.POST("findUserPoemByPage", baseApi.FindUserPoemByPage)
	}
	return userRouter
}
