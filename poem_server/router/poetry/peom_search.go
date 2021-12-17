package poetry

import (
	"github.com/gin-gonic/gin"
	v1 "poem_server/api/v1"
)

type SearchRouter struct {
}

func (a *SearchRouter) InitSearchRouter(Router *gin.RouterGroup) {
	authorRouterWithoutRecord := Router.Group("all")
	apiRouterApi := v1.PoemApiGroupApp.PoetryApiGroup.SearchApi
	{
		authorRouterWithoutRecord.GET("searchAll", apiRouterApi.SearchAll)
	}
}
