package initialize

import (
	"github.com/gin-gonic/gin"
	"jjchatgpt/global"
	"jjchatgpt/middleware"
	"jjchatgpt/router"
)

func Routers() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	Router := gin.Default()

	router.InitUI(Router)
	ApiGroup := Router.Group("base", middleware.Cors())
	router.InitBaseRouter(ApiGroup)

	global.Chatgptcn_LOG.Info("router register success")
	return Router
}
