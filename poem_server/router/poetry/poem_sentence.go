package poetry

import (
	"github.com/gin-gonic/gin"
	v1 "poem_server/api/v1"
)

type SentenceRouter struct {
}

func (a *SentenceRouter) InitSentenceRouter(Router *gin.RouterGroup) {
	sentenceRouterWithoutRecord := Router.Group("sentence")
	apiRouterApi := v1.PoemApiGroupApp.PoetryApiGroup.SentenceApi
	{
		sentenceRouterWithoutRecord.POST("getSentenceList", apiRouterApi.GetSentenceList)
		//authorRouterWithoutRecord.GET("getAuthorByID/:oid", apiRouterApi.GetAuthorById)
	}
}
