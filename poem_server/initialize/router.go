package initialize

import (
	_ "poem_server/docs"
	"poem_server/global"
	"poem_server/router"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Router() *gin.Engine {
	var Router = gin.Default()
	// 首先添加一些静态资源（按需要进行添加）
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	poetryRouter := router.PoemRouterGroupApp.Poetry
	systemRouter := router.PoemRouterGroupApp.System
	PublicRouter := Router.Group("")
	{
		PublicRouter.GET("/health", func(context *gin.Context) {
			context.JSON(200, "ok")
		})
		systemRouter.InitBaseRouter(PublicRouter)
	}
	PrivateGroup := Router.Group("")

	{
		systemRouter.InitUserRouter(PrivateGroup)
		poetryRouter.InitAuthorRouter(PrivateGroup)   // 注册作者相关功能
		poetryRouter.InitPoemRouter(PrivateGroup)     // 注册诗句相关功能
		poetryRouter.InitSentenceRouter(PrivateGroup) // 注册名句相关功能
		poetryRouter.InitSearchRouter(PrivateGroup)   // 搜索相关功能
	}
	global.POEM_LOG.Info("router register success")

	return Router
}
