package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	gogpt "jjchatgpt/chatgpt"
	"jjchatgpt/config"
)

var (
	Chatgptcn_CONFIG config.Server
	Chatgptcn_LOG    *zap.Logger
	Chatgptcn_VP     *viper.Viper
	Gogpt            *gogpt.Client
)
