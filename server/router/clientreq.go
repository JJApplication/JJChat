package router

import (
	"github.com/gin-gonic/gin"
	"jjchatgpt/api/v1"
	"jjchatgpt/model/response"
)

func InitBaseRouter(router *gin.RouterGroup) {
	{
		router.OPTIONS("quest", func(c *gin.Context) {
			response.Ok(c)
		})
		router.POST("quest", v1.Questansw)
		router.GET("quest", func(c *gin.Context) {
			response.Ok(c)
		})
	}
}
