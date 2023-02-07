package core

import (
	"fmt"
	"go.uber.org/zap"
	"jjchatgpt/global"
	"jjchatgpt/initialize"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.Chatgptcn_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.Chatgptcn_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf("欢迎使用 JJChatGPT. 访问:%s开始体验\n", address)
	fmt.Printf("在外部打开: http://localhost%s\n", address)
	global.Chatgptcn_LOG.Error(s.ListenAndServe().Error())
}
