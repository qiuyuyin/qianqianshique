package poetry

import (
	"github.com/gin-gonic/gin"
	v1 "poem_server/api/v1"
)

type AuthorRouter struct {
}

func (a *AuthorRouter) InitAuthorRouter(Router *gin.RouterGroup) {
	authorRouterWithoutRecord := Router.Group("author")
	apiRouterApi := v1.PoemApiGroupApp.PoetryApiGroup.AuthorApi
	{
		authorRouterWithoutRecord.POST("getAuthorList", apiRouterApi.GetAuthorList)
		authorRouterWithoutRecord.GET("getAuthorByID/:oid", apiRouterApi.GetAuthorById)
		authorRouterWithoutRecord.GET("getAuthorPoemCount/:limit", apiRouterApi.GetAuthorPoemCountByLimit)
	}
}
