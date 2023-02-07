package initialize

import (
	gogpt "jjchatgpt/chatgpt"
	"jjchatgpt/global"
)

func InitGoGPTclient() (client *gogpt.Client) {
	api_key := global.Chatgptcn_CONFIG.Proxy.Chatgpttoken
	client = gogpt.NewClient(api_key)
	return
}
