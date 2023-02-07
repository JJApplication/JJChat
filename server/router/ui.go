package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitUI(router *gin.Engine) {
	router.StaticFile("/", "ui/index.html")
	router.StaticFS("/js", http.Dir("ui/js"))
	router.StaticFS("/css", http.Dir("ui/css"))
	router.StaticFS("/img", http.Dir("ui/img"))
	router.StaticFS("/fonts", http.Dir("ui/fonts"))
}
