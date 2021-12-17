package poetry

import (
	"github.com/gin-gonic/gin"
	v1 "poem_server/api/v1"
)

type PoemRouter struct {
}

func (r PoemRouter) InitPoemRouter(Router *gin.RouterGroup) {
	authorRouterWithoutRecord := Router.Group("poem")
	apiRouterApi := v1.PoemApiGroupApp.PoetryApiGroup.PoemApi
	{
		authorRouterWithoutRecord.POST("getPoemList", apiRouterApi.GetPoemList)
		authorRouterWithoutRecord.GET("getPoemByID/:poemType/:oid", apiRouterApi.GetPoemById)
	}
}
