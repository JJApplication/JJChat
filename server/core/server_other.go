package core

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

func initServer(address string, router *gin.Engine) server {
	fmt.Println("server started by [endless]")
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 30 * time.Second
	s.WriteTimeout = 60 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
